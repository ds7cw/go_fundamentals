package main

import "fmt"

func main() {
	myBill := newBill("My First Bill :)")
	fmt.Println(myBill)       // {My First Bill :) map[] 0}
	fmt.Println(myBill.name)  // My First Bill :)
	fmt.Println(myBill.items) // map[]
	fmt.Println(myBill.tip)   // 0

	fmt.Println(myBill.format())
	// Bill breakdown:
	// pie: ...$5.99
	// cake: ...$3.99
	// total: ...$9.98

}
