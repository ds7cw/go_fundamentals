package main

import "fmt"

func main() {
	menu := map[string]float64{
		"salad":      4.99,
		"fries":      5.99,
		"pasta":      7.99,
		"cheesecake": 6.99,
	}
	fmt.Println(menu)
	fmt.Println(menu["pasta"])

	// Iterate over maps
	for k, v := range menu {
		fmt.Println(k, "~", v)
	}

	playersMap := map[int]string{
		1: "Kobe Bryant",
		2: "Allen Iverson",
		3: "LeBron James",
	}

	fmt.Println(playersMap)
	fmt.Println(playersMap[2])

	// Update a map value
	playersMap[3] = "Kevin Durant"
	fmt.Println(playersMap)

}
