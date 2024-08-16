package main

import (
	"fmt"
)

// func min(x, y int) int {
// 	if x < y {
// 		return x
// 	}
// 	return y
// }

func fibonacciSearch(arr []int, target int) int {
	n := len(arr)
	fibMMm2 := 0              // (m-2)'th Fibonacci No.
	fibMMm1 := 1              // (m-1)'th Fibonacci No.
	fibM := fibMMm2 + fibMMm1 // m'th Fibonacci

	// fibM is going to store the smallest Fibonacci
	// number greater than or equal to n
	for fibM < n {
		fibMMm2 = fibMMm1
		fibMMm1 = fibM
		fibM = fibMMm2 + fibMMm1
	}

	offset := -1

	// while there are elements to be inspected.
	// Note that we compare arr[fibMMm2] with target.
	// When fibM becomes 1, fibMMm2 becomes 0
	for fibM > 1 {
		// Check if fibMMm2 is a valid location
		i := min(offset+fibMMm2, n-1)

		// If target is greater than the value at index fibMMm2,
		// cut the subarray array from offset to i
		if arr[i] < target {
			fibM = fibMMm1
			fibMMm1 = fibMMm2
			fibMMm2 = fibM - fibMMm1
			offset = i
		} else if arr[i] > target {
			// If target is less than the value at index fibMMm2,
			// cut the subarray after i+1
			fibM = fibMMm2
			fibMMm1 = fibMMm1 - fibMMm2
			fibMMm2 = fibM - fibMMm1
		} else {
			return i
		}
	}

	// comparing the last element with target
	if fibMMm1 == 1 && arr[offset+1] == target {
		return offset + 1
	}

	// element not found
	return -1
}

func main() {
	// Test fibonacci search
	arr := []int{10, 22, 35, 40, 45, 50, 80, 82, 85, 90, 100}
	target := 85
	index := fibonacciSearch(arr, target)

	if index >= 0 {
		fmt.Printf("Found element %d at index: %d\n", target, index)
	} else {
		fmt.Printf("Element %d Not found", target)
	}
}
