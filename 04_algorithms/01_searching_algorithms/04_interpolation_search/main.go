package main

import "fmt"

// interpolationSearch performs a interpolation search on a sorted array.
// It returns the index of the target element if found, otherwise -1.
func interpolationSearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right && target >= arr[left] && target <= arr[right] {
		if left == right {
			if arr[left] == target {
				return left
			}

			return -1
		}

		fmt.Println("left: ", left, "\nright:", right)

		pos := left + int(float64(right-left)*(float64(target-arr[left])/float64(arr[right]-arr[left])))

		fmt.Println("pos:", pos)

		if arr[pos] == target {
			return pos
		}

		if arr[pos] < target {
			left = pos + 1
		} else {
			right = pos - 1
		}
	}

	return -1
}

func main() {
	// Test interpolation search
	arr := []int{10, 12, 13, 16, 18, 19, 20, 21, 22, 23, 24, 33, 35, 42, 47}
	target := 18
	index := interpolationSearch(arr, target)

	if index != -1 {
		fmt.Printf("Element %d is at index %d\n", target, index)
	} else {
		fmt.Printf("Element %d not found", target)
	}
}
