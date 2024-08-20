package main

import (
	"fmt"
)

// mergeSort is a recursive function that sorts a slice using the Merge Sort algorithm.
// It divides the slice into two halves, recursively sorts them, and then merges the sorted halves.
func mergeSort(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}

	mid := len(slice) / 2
	left := mergeSort(slice[:mid])
	right := mergeSort(slice[mid:])
	return merge(left, right)
}

// merge is a helper function that merges two sorted slices into a single sorted slice.
// It takes two sorted slices as input and returns a single sorted slice.
func merge(left, right []int) []int {
	result := []int{}
	i, j := 0, 0

	// Merge the two slices while maintaining order
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Append any remaining elements from the left slice
	for i < len(left) {
		result = append(result, left[i])
		i++
	}

	// Append any remaining elements from the right slice
	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	return result
}

func main() {
	// Test merge sort
	unsorted := []int{25, 17, 31, 13, 2, 8, 24, 19, 30, 15, 7,
		1, 12, 5, 21, 18, 6, 14, 9, 3, 11, 4, 20, 10, 16}

	sorted := mergeSort(unsorted)

	fmt.Println("Sorted slice:", sorted)

}
