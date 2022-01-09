// Package slices implements various slice algorithms.
package main

import "fmt"

// Map turns a []T1 to a []T2 using a mapping function.
// This function has two type parameters, T1 and T2.
// This works with slices of any type.
func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	r := make([]T2, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

// Reduce reduces a []T1 to a single value using a reduction function.
func Reduce[T1, T2 any](s []T1, initializer T2, f func(T2, T1) T2) T2 {
	r := initializer
	for _, v := range s {
		r = f(r, v)
	}
	return r
}

// Filter filters values from a slice using a filter function.
// It returns a new slice with only the elements of s
// for which f returned true.
func Filter[T any](s []T, f func(T) bool) []T {
	var r []T
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

func main() {
	s := []int{1, 2, 3, 4, 25}

	odds := Filter(s, func(i int) bool { return i%2 != 0 })
	fmt.Println(odds) // Now evens is []int{1, 3, 25}.

	floats := Map(odds, func(i int) float64 { return float64(i * i) })
	fmt.Println(floats) // Now floats is []float64{1, 9, 625}.

	sum := Reduce(floats, 0.0, func(i, j float64) float64 { return i + j })
	fmt.Println(sum) // Now sum is 635.

}
