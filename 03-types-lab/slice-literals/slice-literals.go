package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q) // [2 3 5 7 11 13]

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r) // [true false true true false true]

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	for i, v := range s {
		fmt.Printf("%d -> (%v, %v)\n", i, v.i, v.b) // [{2 true} {3 false} {5 true} {7 true} {11 false} {13 true}]
	}

}
