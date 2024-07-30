package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var s1 []int
	t1 := []int{}
	u1 := make([]int, 5)
	v1 := make([]int, 0, 5) // len 0, cap 5
	fmt.Printf("%d, %d, %T, %5t, %#[3]v\n",
		len(s1), cap(s1), s1, s1 == nil) // 0, 0, []int,  true, []int(nil)
	fmt.Printf("%d, %d, %T, %5t, %#[3]v\n",
		len(t1), cap(t1), t1, t1 == nil) // 0, 0, []int, false, []int{}
	fmt.Printf("%d, %d, %T, %5t, %#[3]v\n",
		len(u1), cap(u1), u1, u1 == nil) // 5, 5, []int, false, []int{0, 0, 0, 0, 0}
	fmt.Printf("%d, %d, %T, %5t, %#[3]v\n",
		len(v1), cap(v1), v1, v1 == nil) // 0, 5, []int, false, []int{}

	// Slices (and maps) encoding differently in JSON when nil
	var a1 []int
	j1, _ := json.Marshal(a1)
	fmt.Println(string(j1)) // null

	b1 := []int{}
	j2, _ := json.Marshal(b1)
	fmt.Println(string(j2)) // []

	a2 := [3]int{1, 2, 3}
	b2 := a2[:1]
	fmt.Println("a2:", a2, "&a2[0]:", &a2[0]) // a2: [1 2 3] &a2[0]: 0xc000016180
	fmt.Println("b2:", b2, "&b2[0]:", &b2[0]) // b2: [1]     &b2[0]: 0xc000016180

	c2 := b2[0:2]
	fmt.Println("c2:", c2, "&c2[0]:", &c2[0]) // c2: [1 2]    &c2[0]: 0xc000016180]

	fmt.Println(len(b2), cap(b2)) // len:1  cap:3
	fmt.Println(len(c2), cap(c2)) // len:2  cap:3

	d2 := a2[0:1:1]
	fmt.Println("len:", len(d2), "cap:", cap(d2), "d2:", d2, "&d2[0]", &d2[0])
	// len: 1 cap: 1 d2: [1] &d2[0] 0xc000016180
	// e2 := d2[0:2] // slice bounds out of range [:2] with capacity 1

	fmt.Printf("a2[%p] = %v\n", &a2, a2) // a2[0xc0000ac030] = [1 2 3]
	fmt.Printf("b2[%p] = %[1]v\n", b2)   // b2[0xc0000ac030] = [1]
	fmt.Printf("c2[%p] = %[1]v\n", c2)   // c2[0xc0000ac030] = [1 2]

}
