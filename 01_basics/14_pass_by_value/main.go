package main

import "fmt"

func updateName(x string) string {
	fmt.Println("&x:", &x)
	x = "updated"
	return x
}

func updatePlayerMap(y map[int]string) {
	y[3] = "Chris Paul"
}

func main() {
	// Go makes "copies" of values when passed into functions
	// Group A types -> strings, ints, floats, arrays, structs
	name := "initial"
	fmt.Println("&name:", &name) // this address will be different from the one on line 6

	name = updateName(name)
	fmt.Println(name)

	// Group A types -> slices, maps, functions
	playersMap := map[int]string{
		1: "Kobe Bryant",
		2: "Allen Iverson",
		3: "LeBron James",
	}

	fmt.Println(playersMap)
	updatePlayerMap(playersMap)
	fmt.Println(playersMap)

}
