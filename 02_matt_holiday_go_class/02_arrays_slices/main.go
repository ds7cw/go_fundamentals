package main

import "fmt"

func do(a [3]int, b []int) []int {
	// a = b // SYNTAX ERROR
	a[0] = 4 // w unchanged
	b[0] = 3 // x changed

	d := a
	fmt.Println(d, &d[0])

	c := make([]int, 5) // []int{0,0,0,0,0}
	c[4] = 42
	copy(c, b) // copies only 3 elements

	return c
}

func main() {
	// Arrays are passed by value, thys elements are copied
	// var a [3]int
	b := [3]int{0, 1, 0}
	// var c [...]int // sized by initializer

	var d [3]int
	fmt.Println(d, &d[1]) // [0 0 0] 0xc000016188
	d = b
	fmt.Println(d, &d[1]) // [0 1 0] 0xc000016188

	// Slices are passed by reference, no copying, updating OK
	var slice_a []int         // []
	var slice_b = []int{1, 2} // [1 2]

	slice_a = append(slice_a, 1) // [1]
	slice_b = append(slice_b, 3) // [1 2 3]

	slice_d := make([]int, 5)             // []int{0, 0, 0, 0, 0}
	slice_e := slice_a                    // same storage (alias)
	fmt.Println(&slice_a[0], &slice_e[0]) // 0xc00000a0e0 0xc00000a0e0
	fmt.Println(slice_d)                  // [0, 0, 0, 0, 0]
	fmt.Println(slice_e[0] == slice_b[0]) // true

	// Slice						Array
	// Variable length				Length fixed at compile time
	// Passed by reference			Passed by value
	// Not comparable				Comparable (==)
	// Can't be used as map key		Can be used as map key
	// Has copy & append helpers	-
	// Useful as func params		Useful as "pseudo" constants

	var w = [...]int{1, 2, 3}
	var x = []int{0, 0, 0}

	y := do(w, x)
	fmt.Println(w, x, y) // [1 2 3] [3 0 0] [3 0 0 0 42]

}
