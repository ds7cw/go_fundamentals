package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func exerciseFunc() {
	scan := bufio.NewScanner(os.Stdin)
	words := make(map[string]int)

	scan.Split(bufio.ScanWords)
	for scan.Scan() {
		words[scan.Text()]++
	}
	fmt.Println(len(words), "unique words")

	type kv struct {
		key string
		val int
	}

	var ss []kv

	for k, v := range words {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].val > ss[j].val
	})

	for _, s := range ss[:3] {
		fmt.Println(s.key, "appears", s.val, "times")
	}
}

func main() {
	// You can read from a nil map, but inserting will panic
	var m1 map[string]int      // map[]; nil, no storage
	p1 := make(map[string]int) // map[]; non-nil but empty

	a1 := p1["the"] // returns 0
	b1 := m1["the"] // returns 0

	fmt.Println(a1)
	fmt.Println(b1)
	// m1["and"] = 1 //PANIC - nil map, map var is nil, no hash map behind it
	m1 = p1
	m1["and"]++ // map[and:1]; OK, same map as p1 now
	c1 := p1["and"]
	fmt.Println(m1, p1, c1) // map[and:1] map[and:1] 1
	// Maps are passed by reference, no copying, updating OK
	// The type used for the key must have == and != defined
	// Slices, maps and funcs cannot be keys

	var m2 = map[string]int{
		"and": 1,
		"the": 1,
		"or":  2,
	}

	var n2 map[string]int
	// b2 := m2 == n2 // SYNTAX ERROR maps have no == operator
	c2 := n2 == nil // true
	d2 := len(m2)   // 3
	// e2 := cap(m) // TYPE MISMATCH
	fmt.Println(c2, d2) // true 3

	// Maps have a special 2-result lookup func
	// The second var tells you if the key was there
	p3 := map[string]int{} // non-nil but empty
	a3 := p3["the"]        // returns 0
	b3, okb := p3["and"]   // 0, false
	p3["the"]++
	c3, okc := p3["the"] // 1, true
	fmt.Printf("p3 %v; a3 %v; b3 %v okb %v; c3 %v okc %v\n", p3, a3, b3, okb, c3, okc)
	// p3 map[the:1]; a3 0; b3 0 okb false; c3 1 okc true;

	// if w, ok := p["the"]; ok {
	// 	// we know w is not the default value
	// 	. . .
	// }

	// len(s)				string		string len
	// len(a), cap(a)		array		array len, capacity (constant)
	// make(T, x)			slice		slice of type T with len x and capacity x
	// make(T, x, y)		slice		slice of type T with len x and capacity y
	// copy(c, d)			slice		copy from d to c; # = min of the 2 lengths
	// c = append(c, d)		slice		append d to c and return a new slice result
	// len(s), cap(d)		slice		slice length and capacity
	// make(T)				map			map of type T
	// make(T, x)			map			map of type T with space hint for x elements
	// delete(m, k)			map			delete key k (if present, else no change)
	// len(m)				map			map length

	// Nil is a type of 0: indicates absence of something
	// Many built-ins are safe: len, cap, range
	var s4 []int
	var m4 map[string]int
	l4 := len(s4)                    // len of a nil slice is 0
	i4, oki := m4["int"]             // 0, false for any missing key
	fmt.Println(s4, m4, l4, i4, oki) // [] map[] 0 0 false

	// for _, v := range s4 {
	// 	. . . // skip if s4 is nil or empty
	// }

	exerciseFunc() // go run main.go < test.txt
	// 31 unique words
	// to appears 3 times
	// is appears 2 times
	// some appears 2 times

}
