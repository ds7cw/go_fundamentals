package main

import "fmt"

func main() {
	myBill := newBill("My First Bill :)")
	fmt.Println(myBill)       // {My First Bill :) map[] 0}
	fmt.Println(myBill.name)  // My First Bill :)
	fmt.Println(myBill.items) // map[]
	fmt.Println(myBill.tip)   // 0

}
