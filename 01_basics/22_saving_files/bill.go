package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// Make new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

// Format bill
// The '()' after the 'func' keyword sets the receiver
func (b *bill) format() string {
	fs := "Bill breakdown:\n"
	var total float64

	// List items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%0.2f\n", k+":", v)
		total += v
	}

	// Add tip
	fs += fmt.Sprintf("%-25v ...$%v\n", "tip:", b.tip)

	// Total
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total+b.tip)

	return fs
}

// Update tip
func (b *bill) updateTip(t float64) {
	// '(*b).tip' can be used to de-reference a struct
	// but it's not needed when using receiver funcs
	b.tip = t

}

// Add item to bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

// Save bill
func (b *bill) saveBill() {
	data := []byte(b.format())

	err := os.WriteFile("./bills/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Bill was saved to file")
}
