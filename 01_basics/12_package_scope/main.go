package main

import (
	"fmt"
)

// score variable is used by showScore() which is defined in greetings.go
// if we place 'score' inside main(), then greetings.go would not be able to access it
// anything inside the main() func is outside the package scope
var score = 35.99

func main() {
	sayHello("Quentin Tarantino")

	for _, v := range points {
		fmt.Println(v)
	}

	showScore()

}
