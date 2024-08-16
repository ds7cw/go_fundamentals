package main

import (
	"fmt"
	"math"
)

// Binary search function
func binarySearch(arr []int, left, right, target int) int {
	if right >= left {
		mid := left + (right-left)/2
		fmt.Printf("left: %d, mid: %d, right: %d\n", left, mid, right)

		// Check if target is present at mid
		if arr[mid] == target {
			return mid
		}

		// If target is smaller, ignore right half
		if arr[mid] > target {
			return binarySearch(arr, left, mid-1, target)
		}

		// If target is larger, ignore left half
		return binarySearch(arr, mid+1, right, target)
	}

	// Element is not present in array
	return -1
}

// Exponential search function
func exponentialSearch(arr []int, n, target int) int {
	// If target is present at first location
	if arr[0] == target {
		return 0
	}

	// Find range for binary search by repeated doubling
	i := 1
	for i < n && arr[i] <= target {
		i = i * 2
		fmt.Println("i value:", i)
	}

	// Call binary search for the found range
	return binarySearch(arr, i/2, int(math.Min(float64(i), float64(n-1))), target)
}

func main() {
	// Test exponential search
	arr := []int{2, 3, 4, 10, 40}
	n := len(arr)
	target := 3
	result := exponentialSearch(arr, n, target)

	if result == -1 {
		fmt.Printf("Element %d is not present in array", target)
	} else {
		fmt.Printf("Element %d is present at index %d", target, result)
	}
}
