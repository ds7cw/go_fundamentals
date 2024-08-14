package main

import "fmt"

// Returns the index at which the target was found
// Returns -1 if the target isn't found
func binarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1 // target not found

}

func main() {
	// Test binary search
	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 2
	result := binarySearch(arr1, target)

	if result != -1 {
		fmt.Printf("Element '%d' found at index %d.\n", target, result)
	} else {
		fmt.Printf("Element '%d' not found in the array.\n", target)
	}

}
