package main

import (
	"fmt"
	"strings"
)

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string
	for _, name := range names {
		initials = append(initials, name[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"

}

func printInitials(f, s string) {
	fmt.Printf("First Name: %v\nSecond Name: %v\n", f, s)

}

func main() {
	f1, s1 := getInitials("ivan locke")
	f2, s2 := getInitials("Tarantino")

	printInitials(f1, s1)
	printInitials(f2, s2)

}
