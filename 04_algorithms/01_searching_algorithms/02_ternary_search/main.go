package main

import "fmt"

// TernarySearch performs a ternary search on a sorted array.
// It returns the index of the target element if found, otherwise -1.
func ternarySearch(arr []int, left, right, target int) int {
	if right >= left {
		mid1 := left + (right-left)/3
		mid2 := right - (right-left)/3

		if arr[mid1] == target {
			return mid1
		}
		if arr[mid2] == target {
			return mid2
		}

		if target < arr[mid1] {
			return ternarySearch(arr, left, mid1-1, target)
		} else if target > arr[mid2] {
			return ternarySearch(arr, mid2+1, right, target)
		} else {
			return ternarySearch(arr, mid1+1, mid2-1, target)
		}
	}

	return -1

}

func main() {
	// Test ternary search
	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 9
	result := ternarySearch(arr1, 0, len(arr1)-1, target)

	if result != -1 {
		fmt.Printf("Element '%d' found at index %d.\n", target, result)
	} else {
		fmt.Printf("Element '%d' not found in the array.\n", target)
	}

}
