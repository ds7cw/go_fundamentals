package main

import (
	"fmt"
)

// Funcs are "first class" objects, you can:
// define them - even inside another function
// create anonymous function literals
// pass them as function params/ return values
// store them in variables
// store them in slices and maps (but not as keys)
// store them as fields of a structure type
// send and receive them in channels
// write methods against a function type
// compare a function var against nil

// The signature of a function is the order & type
// of its parameters and return values
// It doesn't depend on the names of those params/ returns

// var try func(string, int) string
// func Do(a string, b int) string {}

// A func declaration lists formal params
// A func call has actual params (a.k.a. "arguments")

// A param is passed by value if the function gets a copy;
// the caller can't see changes to the copy
// A param is passed by reference if the func can modify
// the actual parameter such that the caller sees the changes

func arrFunc(b [2]int) int {
	b[0] = 0
	return b[1]
}

func mapFunc(m1 map[int]int) {
	m1[3] = 0              // updates the original map
	m1 = make(map[int]int) // new map stored in 'm1' only, not in 'm'
	m1[4] = 4
	fmt.Println("m1", m1) // m1 map[4:4]
}

func main() {
	a := [2]int{1, 2}
	v := arrFunc(a)
	fmt.Println("a:", a, "v:", v) // a: [1 2] v: 2
	// If 'a' is a slice, the func would update its values

	m := map[int]int{4: 1, 7: 2, 8: 3}
	mapFunc(m)
	fmt.Println("m", m) // m map[3:0 4:1 7:2 8:3]

	// Actually, all params are passed by copying something i.e. by value
	// If the thing copied is a pointer or descriptor, then the shared
	// backing store (array, hash table, etc.) can be changed through it

	// Recursion
	// A function may call itself; the trick is knowing when to stop
	// func walk(node *tree.T) int {
	// 	if node == nil {
	// 		return 0
	// 	}
	// 	return node.value + walk(node.left) + walk(node.right)
	// }
	// This works because each function call adds context to the stack
	// and unwinds it when done
	// The program will crash, if you don't have good stopping criteria

	// Deferred execution
	// How do we make sure something gets done?
	// close a file we opened
	// close a socket/ HTTP request we made
	// unlock a mutex we locked
	// make sure something gets saved before we're done ...

	// The defer statement captures a function call to run later
	// We need to ensure the file closes
	// f, err := os.Open("my_file.txt")
	// if err != nil {
	// 	. . .
	// }
	// defer f.Close() // and do something with the file
	// The call to Close is guaranteed to run at function exit
	// (don't defer closing the file until we know it really opened!)

	// defer operates on a function scope, not on a block scope
	// defer happens when the function exits
	// func main() {
	// f := os.Stdin
	// if len(os.Args) > 1 {
	// 	if f, err := os.Open(os.Args[1]); err != nil {
	// 		. . .
	// 	}
	// 	defer f.Close()
	// }
	// And do something with the file
	// }
	// The defer will not execute when we leave the if block

	// Unlike a closure, defer copies args to the deferred call
	defVar := 10
	defer fmt.Println(defVar, &defVar)
	defVar = 11
	defer fmt.Println(defVar, &defVar)
	// 11 0xc0000a6100
	// 10 0xc0000a6100
	// Param 'a' gets copied at the defer statement (not a reference)

	// A defer statement runs before the return of a function
	// func do() (a int) { // named return ('a' int)
	// 	defer func() {
	// 		a = 2
	// 	} () // notice the call after declaration
	// 	a = 1
	// 	return // naked return still returns 'a'
	// } // returns 2
	// The deferred anonymous func can update that var

}
