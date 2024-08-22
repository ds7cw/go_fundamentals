package main

import (
	"container/heap"
	"fmt"
)

// This code defines a min-heap using Go’s container/heap package and then uses it to sort an array.

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// The heapSort function pushes all elements of the slice into the heap and then pops them out in sorted order.
func heapSort(slice []int) []int {
	h := &IntHeap{}
	heap.Init(h)
	for _, v := range slice {
		heap.Push(h, v)
	}
	sorted := make([]int, 0, len(slice))
	for h.Len() > 0 {
		sorted = append(sorted, heap.Pop(h).(int))
	}
	return sorted
}

func main() {
	// Test heap sort
	slice := []int{5, 5, 3, 8, 4, 2, 7, 1, 10}
	sortedslice := heapSort(slice)

	fmt.Println("Sorted Slice:", sortedslice)
	// Sorted Slice: [1 2 3 4 5 5 7 8 10]

}
