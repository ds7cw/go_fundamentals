package main

import "fmt"

func main() {
	age := 27
	fmt.Println("age <= 30:", age <= 30) // true
	fmt.Println("age >= 30:", age >= 30) // true
	fmt.Println("age == 27:", age == 27) // true
	fmt.Println("age != 30:", age != 30) // true

	if age < 20 {
		fmt.Println("Age is less than 20.")
	} else if age < 27 {
		fmt.Println("Age is less than 27.")
	} else {
		fmt.Println("Age is not less than 27.")
	}

	movies := []string{"Full Metal Jacket (1987)", "The Shining (1980)", "Dr. Strangelove (1964)", "Barry Lyndon (1975)"}
	for idx, title := range movies {
		if idx == 1 {
			fmt.Println("'continue' key word used @ index", idx)
			continue
		} else if idx > 2 {
			fmt.Println("'break' key word used @ index", idx)
			break
		}

		fmt.Printf("Current Index %v, Current movie title: %v\n", idx, title)
	}

}
