package main

import "fmt"

func main() {
	// ARRAYS
	var arr1 [3]int = [3]int{11, 23, 29}
	fmt.Println(arr1)
	fmt.Println("Array element at idx 2:", arr1[2])
	names := [4]string{"Bruce", "Peter", "Tony", "Clark"}
	fmt.Println("Array length:", len(names), "Array contents", names)

	// SLICES
	var slice1 []int = []int{7, 13, 31}
	fmt.Println(slice1)
	slice1 = append(slice1, 43)
	// Slice Range
	range1 := slice1[0:2]
	range2 := slice1[2:]
	fmt.Println(range1)
	fmt.Println(range2)
	range2 = append(range2, 113) // Change only applied to range2
	fmt.Println(range2)
	fmt.Println(slice1)

}
