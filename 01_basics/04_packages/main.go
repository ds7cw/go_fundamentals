package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	// QUOTE PACKAGE
	fmt.Println("\n" + quote.Go())
	fmt.Println(quote.Glass())

	// FMT PACKAGE
	// fmt.Print() does not add a new line after printing the output
	fmt.Print("Hello,")
	fmt.Print(" World!")

}
