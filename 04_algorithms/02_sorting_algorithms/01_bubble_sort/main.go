package main

import "fmt"

// bubbleSort function takes a slice of integers and sorts it.
// The function returns the sorted slice.
func bubbleSort(slice []int) []int {
	n := len(slice)

	for i := 0; i < n; i++ {
		swapped := false

		for j := 0; j < n-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
				swapped = true
			}
		}

		// If no elements were swapped in the inner loop, break
		if !swapped {
			fmt.Println("i counter value:", i)
			break
		}
	}

	return slice

}

func main() {
	// Test bubble sort
	unsorted := []int{10, 6, 2, 1, 5, 8, 3, 4, 7, 9}
	sorted := bubbleSort(unsorted)

	fmt.Println("Sorted slice:", sorted)

}
