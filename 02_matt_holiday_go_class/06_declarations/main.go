package main

import "fmt"

func main() {
	// Declarations
	// There are six ways to introduce a name:
	// 1. Constatnt declaration with const
	// 2. Type declaration using type
	// 3. Variable declaration with var
	// (must have type or initial value, sometimes both)
	// 4. Short, initialized variable declaration of any type :=
	// (only inside a function)
	// 5. Function declaration with func
	// (methods may only be declared at package level)
	// 6. Formal parameters and named returns of a function

	var (
		x, y int
		z    float64
		s    string
	)
	fmt.Println(x, y, z, s) // 0 0 0

	// Short declaration operator := rules:
	// 1. It can't be used outside of a function
	// 2. It must be used (instead of var) in a control statement (if, etc)
	// 3. It must declare at least one new variable

	// err := doSomething();
	// err := doSomethingElse(); // WRONG
	// x, err := getSomeValue(); // OK; err is not redeclared

	// 4. It won't re-use an existing declaration from an outer scope

	// Shadowing short declarations
	// func main() {
	// 	n, err := fmt.Println("Hello, playground")
	// 	if _, err := fmt.Println(n); err != nil {
	// 		fmt.Println(err)
	// 	}
	// }
	// Compile error: the first 'err' is unused
	// This follows from the scoping rules, because := is a declaration
	// and the second err is in the scope of the if statement

	// func BadRead(f *os.File, buf []byte) error {
	// 	var err error
	// 	for {
	// 		n, err := f.Read(buf) // shadows 'err' above
	// 		if err != nil {
	// 			break			// causes return of WRONG value
	// 		}
	// 		foo(buf)
	// 	}
	// 	return err // will always be nil
	// }

	// Structural typing
	// It's the same type if it has the same structure or behavior
	a := [...]int{1, 2, 3}
	b := [3]int{}
	fmt.Println("Array A:", a, "Array B:", b) // [1 2 3] [0 0 0]
	a = b
	fmt.Println("Array A:", a, "Array B:", b) // [0 0 0] [0 0 0]
	// c := [4]int{}
	// a = c // can't use c (var of type [4]int) as [3]int value
	// in assignmentcompilerIncompatibleAssign

	// Go uses structural typing in most cases
	// It's the same type if it has the same structure or behavior:
	// arrays of the same size and base type
	// slices with the same base type
	// maps of the same key and value types
	// structs with the same sequence of field names/ types
	// functions with the same parameter & return types

	// Named typing
	// It's only the same type if it has the same declared type name

	// type x int // declared outside the main() function
	// func main() {
	// 	var a x  // x is a defined type; base int
	// 	b := 12  // b defaults to int
	// 	a = b    // TYPE MISSMATCH
	// 	a = 12   // OK, untyped literal
	// 	a = x(b) // OK, type conversion
	// }
	// Go uses named typing for non-function user-declared types

	// Basic operators
	// Arithmetic: numbers only except + on string
	//   +  -  *  /  %  ++  --
	// Comparison: only numbers/ strings support order
	//   ==  !=  <  >  <=  >=
	// Boolean: only booleans, with shortcut evaluation
	//   ! &&  ||
	// Bitwise: operate on integers
	//   &  |  ^  <<  >>  &^
	// Assignment: as above for binary operations
	//    =  +=  -=  *=    /=    %=
	//   &=  |=  ^=  <<==  >>==  &^=

	// Operator precedence
	// There are only five levels of precedence, otherwise left-to-right:
	// Operators like multiplication: * / % << >>  &^
	// Operators like addition: + - | ^
	// Comparison operators: == != < <= > >=
	// Logical and: &&
	// Logical or: ||

}
