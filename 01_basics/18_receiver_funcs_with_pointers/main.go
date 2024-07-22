package main

import "fmt"

func main() {
	myBill := newBill("My First Bill :)")

	// Update tip
	fmt.Println(myBill.tip) // 0
	myBill.updateTip(8)
	fmt.Println(myBill.tip) // 8

	// Add item
	myBill.addItem("Garlic bread", 3.75)
	myBill.addItem("Cottage pie", 8.50)
	myBill.addItem("Green tea", 2.99)

	fmt.Println(myBill.format())
	// Bill breakdown:
	// Garlic bread:             ...$3.75
	// Cottage pie:              ...$8.50
	// Green tea:                ...$2.99
	// tip:                      ...$8
	// total:                    ...$23.24

}
