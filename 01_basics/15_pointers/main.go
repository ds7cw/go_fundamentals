package main

import "fmt"

func updateNamePointer(x *string) {
	*x = "updated"
}

func main() {
	// Non-Pointer Values:
	// strings, ints, floats, booleans, arrays, structs
	myValue := "Tracy McGrady"
	fmt.Println("Memory address (&myValue):", &myValue)            // 0xc000028070
	fmt.Println("Value behind memory address (myValue):", myValue) // Tracy McGrady

	// Store memory address of a variable
	myValueAddress := &myValue                                                          // this is a pointer
	fmt.Println("Memory address of myValue:", myValueAddress)                           // 0xc000028070
	fmt.Println("Memory address of myValueAddress (&myValueAddress):", &myValueAddress) // 0xc00005a030
	fmt.Println("Vlue at memory address (*myValueAddress):", *myValueAddress)           // Tracy McGrady

	fmt.Println("My value before function call (myValue):", myValue) // Tracy McGrady
	updateNamePointer(myValueAddress)
	fmt.Println("My value after function call (myValue):", myValue) // updated

	// Pointer Values:
	// slices, maps, functions
}
