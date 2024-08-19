package main

import "fmt"

// insertionSort function takes a slice of integers and sorts it.
// The function returns the sorted slice.
func insertionSort(slice []int) []int {
	n := len(slice)

	for i := 1; i < n; i++ {
		key := slice[i]
		j := i - 1
		sorted := true

		// Move elements of slice[0..i-1], greater than key,
		// to one position ahead of their current position
		for j >= 0 && slice[j] > key {
			slice[j+1] = slice[j]
			j--
			sorted = false
		}

		slice[j+1] = key

		// If no elements were moved, slice is already sorted
		if sorted {
			break
		}
	}

	return slice
}

func main() {
	// Test insertion sort
	unsorted := []int{10, 6, 2, 1, 5, 8, 3, 4, 7, 9}
	sorted := insertionSort(unsorted)

	fmt.Println("Sorted slice:", sorted)

}
