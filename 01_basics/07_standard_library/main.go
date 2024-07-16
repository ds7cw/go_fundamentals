package main

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func main() {
	// STRINGS PACKAGE
	var name1 string = "Geralt of Rivia"
	name2 := "Butcher of Blaviken"
	// Contains
	fmt.Println(strings.Contains(name1, "via"))     // returns True
	fmt.Println(strings.Contains(name2, "butcher")) // returns False
	// Replace
	fmt.Println(strings.ReplaceAll(name1, "a", "@")) // Ger@lt of Rivi@
	// ToUpper
	fmt.Println(strings.ToUpper(name1)) // GERALT OF RIVIA
	// Index
	fmt.Println(strings.Index(name2, "u")) // 1
	fmt.Println(strings.Index(name2, "B")) // 0; Returns the first match
	// Split
	fmt.Println(strings.Split(name2, " ")) // [Butcher of Blaviken]
	split_data_structure := strings.Split(name2, " ")
	fmt.Printf("Type of argument returned from Split method: %T", split_data_structure) // []string
	fmt.Println()
	fmt.Println(reflect.TypeOf(split_data_structure)) // []string

	// SORT
	// sort.Ints
	arr1 := []int{23, 11, 7, 37, 31, 57, 53, 17, 29}
	sort.Ints(arr1)
	fmt.Println(arr1) // [7 11 17 23 29 31 37 53 57]
	// SearchInts
	index := sort.SearchInts(arr1, 23)
	fmt.Println(index) // 3; looking for a non-existent element will return an inde equal to (slice length + 1)
	// sort.Strings
	names1 := []string{"TMac", "Kobe", "Shaq", "Iverson"}
	sort.Strings(names1)
	fmt.Println(names1) // [Iverson Kobe Shaq TMac]
	// SearchStrings
	fmt.Println(sort.SearchStrings(names1, "Kobe")) // 1

}
