package main

import "fmt"

func greatMovie(title string) {
	fmt.Println("A great movie:", title)
}

func badMovie(title string) {
	fmt.Println("A bad movie:", title)
}

func cycleTitles(movies []string, f func(string)) {
	for _, title := range movies {
		f(title)
	}
}

func squareArea(side float32) float32 {
	result := side * side
	return result
}

func main() {
	// Function takes a string argument
	greatMovie("Goodfellas (1990)")
	badMovie("Sex Lives of the Potato Men (2004)")

	// Function takes another function as the argument
	m := []string{"Full Metal Jacket (1987)", "Nightcrawler (2014)", "The Prestige (2006)"}
	cycleTitles(m, greatMovie)

	// Function returns a value
	fmt.Println(squareArea(4.5))

}
