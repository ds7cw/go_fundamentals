package main

import "fmt"

func main() {
	i := 0
	for i < 4 {
		fmt.Println("Value of 'i':", i)
		i++
	}

	for j := 0; j < 4; j++ {
		fmt.Println("Value of 'j':", j)
	}

	names := []string{"Pulp Fiction", "Death Proof", "Reservoir Dogs"}
	for k := 0; k < len(names); k++ {
		fmt.Println(names[k])
	}

	fmt.Println()
	for idx, name := range names {
		fmt.Printf("Value at index %v : %v\n", idx, name)
	}
	// User 'for _, name := range names' if you don't need the index

}
