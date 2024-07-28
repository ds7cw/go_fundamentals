package main

import "fmt"

func main() {
	// If statements
	// They can start with a short declaration or statement

	// if err := doSomething(); err != nil {
	// 	return err
	// }

	// The loop control structure provides automatic repetition
	// There is only for (no 'do' or 'while') but with options

	for i := 0; i < 4; i++ {
		fmt.Printf("(%d, %d)\n", i, i*i)
	} // (0, 0) (1, 1) (2, 4) (3, 9)

	// Three loop parts, all optional (initialize, check, increment)

	// Implicit control through the range operator for arrays/ slices
	// one var: i is an index 0, 1, 2, ...
	// for i := range myArray {
	// 	fmt.Println(i, myArray)
	// }

	// two vars: i is the index, v is a value
	// for i, v := range myArray {
	// 	fmt.Println(i, v)
	// }

	// Implicit control through the range operator for maps
	// one var: k is key
	// for k := range myMap {
	// 	fmt.Println(k, myMap[k])
	// }

	// two vars: k is the key, v is a value
	// for k, v := range myMap {
	// 	fmt.Println(k, v)
	// }

	// An infinite loop with an explicit break
	i, j := 0, 3
	for { // This is similar to while True in python
		i, j = i+50, j*j
		fmt.Println(i, j)

		if j > i {
			break // when i = 150, j = 6561
		}
	}

	// If you only want range values, you need the blank identifier
	// Two vars: _ is the index (ignored),
	// 				v is the value
	// for _, v := range myArray {
	// 	fmt.Println(v)
	// }

	// for k := range testItemsMap { 		// keys
	// 	for _, v := range returnedData { 	// values in list
	// 		if k == v.ID { 					// found it
	// 			continue outer
	// 		}
	// 	}
	// 	t.Errorf("key not found: %s", k)
	// }

	// Switch
	// It is a shortcut replacing a series of if-then statements

	// switch a := f.Get(); a {
	// case 0, 1, 2:
	// 	fmt.Println("underflow possible")
	// case 3, 4, 5, 6, 7, 8:
	// default:
	// 	fmt.Println("warning: overload")
	// }

	// Arbitrary comparisons may be made for an switch with no argument

	// a := f.Get()
	// switch {
	// case a <= 2:
	// 	fmt.Println("underflow possible")
	// case a <= 8:
	// 	// evaluated in order
	// default:
	// 	fmt.Println("warning: overload")
	// }
}
