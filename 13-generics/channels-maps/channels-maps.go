package main

import (
	"fmt"
)

// Drain drains any elements remaining on the channel.
func Drain[T any](c <-chan T) {
	for range c {
	}
}

// Merge merges two channels of some element type into a single channel.
func Merge[T any](c1, c2 <-chan T) <-chan T {
	r := make(chan T)
	go func(c1, c2 <-chan T, r chan<- T) {
		defer close(r)
		for c1 != nil || c2 != nil {
			select {
			case v1, ok := <-c1:
				if ok {
					r <- v1
				} else {
					c1 = nil
				}
			case v2, ok := <-c2:
				if ok {
					r <- v2
				} else {
					c2 = nil
				}
			}
		}
	}(c1, c2, r)
	return r
}

// Ranger provides a convenient way to exit a goroutine sending values
// when the receiver stops reading them.
//
// Ranger returns a Sender and a Receiver. The Receiver provides a
// Next method to retrieve values. The Sender provides a Send method
// to send values and a Close method to stop sending values. The Next
// method indicates when the Sender has been closed, and the Send
// method indicates when the Receiver has been freed.
func Ranger[T any]() (*Sender[T], *Receiver[T]) {
	c := make(chan T)
	d := make(chan bool)
	s := &Sender[T]{values: c, done: d}
	r := &Receiver[T]{values: c, done: d}
	// The finalizer on the receiver will tell the sender
	// if the receiver stops listening.
	//runtime.SetFinalizer(r, r.finalize)
	return s, r
}

// A Sender is used to send values to a Receiver.
type Sender[T any] struct {
	values chan<- T
	done   <-chan bool
}

// Send sends a value to the receiver. It reports whether any more
// values may be sent; if it returns false the value was not sent.
func (s *Sender[T]) Send(v T) bool {
	select {
	case s.values <- v:
		return true
	case <-s.done:
		// The receiver has stopped listening.
		return false
	}
}

// Close tells the receiver that no more values will arrive.
// After Close is called, the Sender may no longer be used.
func (s *Sender[T]) Close() {
	close(s.values)
}

// A Receiver receives values from a Sender.
type Receiver[T any] struct {
	values <-chan T
	done   chan<- bool
}

// Next returns the next value from the channel. The bool result
// reports whether the value is valid. If the value is not valid, the
// Sender has been closed and no more values will be received.
func (r *Receiver[T]) Next() (T, bool) {
	v, ok := <-r.values
	return v, ok
}

// finalize is a finalizer for the receiver.
// It tells the sender that the receiver has stopped listening.
func (r *Receiver[T]) finalize() {
	close(r.done)
}

// Generic OderedMap, implemented as a binary tree.
// Map is an ordered map.
type Map[K comparable, V any] struct {
	root    *node[K, V]
	compare func(K, K) int
}

// node is the type of a node in the binary tree.
type node[K, V any] struct {
	k           K
	v           V
	left, right *node[K, V]
}

// New returns a new map.
// Since the type parameter V is only used for the result,
// type inference does not work, and calls to New must always
// pass explicit type arguments.
func New[K comparable, V any](compare func(K, K) int) *Map[K, V] {
	return &Map[K, V]{compare: compare}
}

// find looks up k in the map, and returns either a pointer
// to the node holding k, or a pointer to the location where
// such a node would go.
func (m *Map[K, V]) find(k K) **node[K, V] {
	pn := &m.root
	for *pn != nil {
		switch cmp := m.compare(k, (*pn).k); {
		case cmp < 0:
			pn = &(*pn).left
		case cmp > 0:
			pn = &(*pn).right
		default:
			return pn
		}
	}
	return pn
}

// Insert inserts a new key/value into the map.
// If the key is already present, the value is replaced.
// Reports whether this is a new key.
func (m *Map[K, V]) Insert(k K, v V) bool {
	pn := m.find(k)
	if *pn != nil {
		(*pn).v = v
		return false
	}
	*pn = &node[K, V]{k: k, v: v}
	return true
}

// Find returns the value associated with a key, or zero if not present.
// The bool result reports whether the key was found.
func (m *Map[K, V]) Find(k K) (V, bool) {
	pn := m.find(k)
	if *pn == nil {
		var zero V // see the discussion of zero values, above
		return zero, false
	}
	return (*pn).v, true
}

// keyValue is a pair of key and value used when iterating.
type keyValue[K, V any] struct {
	k K
	v V
}

// InOrder returns an iterator that does an in-order traversal of the map.
func (m *Map[K, V]) InOrder() *Iterator[K, V] {
	sender, receiver := Ranger[keyValue[K, V]]()
	var f func(*node[K, V]) bool
	f = func(n *node[K, V]) bool {
		if n == nil {
			return true
		}
		// Stop sending values if sender.Send returns false,
		// meaning that nothing is listening at the receiver end.
		return f(n.left) &&
			sender.Send(keyValue[K, V]{n.k, n.v}) &&
			f(n.right)
	}
	go func() {
		f(m.root)
		sender.Close()
	}()
	return &Iterator[K, V]{receiver}
}

// Iterator is used to iterate over the map.
type Iterator[K, V any] struct {
	r *Receiver[keyValue[K, V]]
}

// Next returns the next key and value pair. The bool result reports
// whether the values are valid. If the values are not valid, we have
// reached the end.
func (it *Iterator[K, V]) Next() (K, V, bool) {
	kv, ok := it.r.Next()
	return kv.k, kv.v, ok
}

// Test ordered map
var m = New[int, string](func(i, j int) int { return i - j })

// Add adds the pair a, b to m.
func Add(a int, b string) {
	m.Insert(a, b)
}

func main() {
	// Set m to an ordered map from string to string,
	// using strings.Compare as the comparison function.
	Add(3, "three")
	Add(1, "one")
	Add(4, "four")
	Add(2, "two")
	iter := m.InOrder()
	for k, v, ok := iter.Next(); ok; k, v, ok = iter.Next() {
		fmt.Println(k, " --> ", v)
	}

}
