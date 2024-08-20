package main

import (
	"fmt"
)

// quickSort is a recursive function that sorts a slice using the Quick Sort algorithm.
// It selects a pivot element, partitions the slice around the pivot, and recursively sorts the sub-slices.
func quickSort(slice []int, left, right int) {
	if left < right {
		// Partition the slice and get the pivot index
		p := partition(slice, left, right)
		// Recursively sort elements before and after partition
		quickSort(slice, left, p-1)
		quickSort(slice, p+1, right)
	}
}

// partition is a helper function that partitions the slice around a pivot element.
// It places all elements smaller than the pivot to the left and all greater elements to the right.
func partition(slice []int, left, right int) int {
	pivot := slice[right]
	i := left
	for j := left; j < right; j++ {
		if slice[j] < pivot {
			slice[i], slice[j] = slice[j], slice[i]
			i++
		}
	}
	slice[i], slice[right] = slice[right], slice[i]
	return i
}

func main() {
	// Test quick sort
	unsorted := []int{25, 17, 31, 13, 2, 8, 24, 19, 30, 15, 7,
		1, 12, 5, 21, 18, 6, 14, 9, 3, 11, 4, 20, 10, 16}

	quickSort(unsorted, 0, len(unsorted)-1)

	fmt.Println("Sorted slice:", unsorted)

}
