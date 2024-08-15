package main

import (
	"fmt"
	"math"
)

// jumpSearch performs a jump search on a sorted array.
// It returns the index of the target element if found, otherwise -1.
func jumpSearch(arr []int, target int) int {
	n := len(arr)
	step := int(math.Sqrt(float64(n)))
	prev := 0

	fmt.Println("Array size:", n, "\nStep value:", step)

	// Jump in steps
	for arr[int(math.Min(float64(step), float64(n)))-1] < target {
		prev = step
		step += int(math.Sqrt(float64(n)))
		fmt.Println("prev:", prev)
		fmt.Println("step:", step)

		if prev >= n {
			return -1
		}
	}

	fmt.Println("Linear search within the block")
	// Linear search within the block
	for arr[prev] < target {
		prev++
		fmt.Println("prev:", prev)
		if prev == int(math.Min(float64(step), float64(n))) {
			return -1
		}
	}

	fmt.Println("prev:", prev)
	if arr[prev] == target {
		return prev
	}

	return -1

}

func main() {
	// Test jump search
	arr := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610}
	target := 55
	index := jumpSearch(arr, target)

	if index != -1 {
		fmt.Printf("Element %d is at index %d\n", target, index)
	} else {
		fmt.Printf("Element %d not found", target)
	}
}
