package main

import (
	"fmt"
	"path/filepath"
	"sort"
)

// Composition is not Inheritance
// 	The fields of an embedded struct are promoted to the
// 	level of the embedding structure

type Pair struct {
	Path string
	Hash string
}

type PairWithLength struct {
	Pair
	Length int
}

func (p Pair) String() string {
	return fmt.Sprintf("Hash of %s is %s", p.Path, p.Hash)
}

func exercise1() {
	pl := PairWithLength{Pair{"/usr", "0xfdfe"}, 121}
	fmt.Println(pl.Path, pl.Length) // not pl.Pair.Path
}

func exercise2() {
	p := Pair{"/usr", "0xfdfe"}
	fmt.Println(p) // Hash of /usr is 0xfdfe
}

func exercise3() { // only String method for Pair
	pl := PairWithLength{Pair{"/usr/lib", "0xfgfg"}, 120}
	fmt.Println(pl) // Hash of /usr/lib is 0xfgfg
}

func (pwl PairWithLength) String() string {
	return fmt.Sprintf("Hash of %s is %s; length %d", pwl.Path, pwl.Hash, pwl.Length)
}

func exercise4() { // String method for PairWithLength
	pl := PairWithLength{Pair{"/usr/lib", "0xfgfg"}, 120}
	fmt.Println(pl) // Hash of /usr/lib is 0xfgfg; length 120
}

func (p Pair) Filename() string {
	return filepath.Base(p.Path)
}

type Filenamer interface {
	Filename() string
}

func exercise5() {
	p := Pair{"/usr", "0xfex5"}
	fmt.Println(p.Filename()) // usr
}

func exercise6() {
	var fn Filenamer = PairWithLength{Pair{"/usr/lib", "0xfex6"}, 120}
	fmt.Println(fn.Filename()) // lib
}

// Composition with pointer types
//
//	A struct can embed a pointer to another type; promotion
//	of its fields and methods works the same way
type Fizgig struct {
	*PairWithLength
	Broken bool
}

func exercise7() {
	fg := Fizgig{
		&PairWithLength{Pair{"/usr", "0xfex7"}, 120},
		false,
	}
	fmt.Println(fg) // Hash of /usr is 0xfex7; length 120
}

// Sortable interface
// 	sort.Interface is defined as:
// type Interface interface {
// 	// Len is the number of elements in the collection
// 	Len() int

// 	// Less reports whether the element with
// 	// idx i should sort before the element with idx j
// 	Less(i, j int) bool

// 	// Swap swaps the elements with indexes i and j
// 	Swap(i, j int)
// }
// //	and sort.Sort as
// func Sort(data Interface)

// Sortable built-ins
// Slices of strings can be sorted using StringsSlice
// 		defined in the sort package
//  	type StringSlice []string
// entries := []string{"charlie", "able", "dog", "baker"}
// sort.Sort(sort.StringSlice(entries)) // [able baker charlie dog]

type Organ struct {
	Name   string
	Weight int
}

type Organs []Organ

func (s Organs) Len() int      { return len(s) }
func (s Organs) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func exercise8() {
	s := []Organ{
		{"brain", 1340},
		{"liver", 1494},
		{"spleen", 162},
		{"pancreas", 131},
		{"heart", 290},
	}
	fmt.Println(s)
	// [{brain 1340} {liver 1494} {spleen 162} {pancreas 131} {heart 290}]
}

type ByName struct{ Organs }
type ByWeight struct{ Organs }

func (s ByName) Less(i, j int) bool {
	return s.Organs[i].Name < s.Organs[j].Name
}

func (s ByWeight) Less(i, j int) bool {
	return s.Organs[i].Weight < s.Organs[j].Weight
}

func exercise9() {
	s := []Organ{
		{"brain", 1340}, {"liver", 1494},
		{"spleen", 162}, {"pancreas", 131}, {"heart", 290},
	}
	sort.Sort(ByWeight{s})
	fmt.Println(s)
	// [{pancreas 131} {spleen 162} {heart 290} {brain 1340} {liver 1494}]
	sort.Sort(ByName{s})
	fmt.Println(s)
	// [{brain 1340} {heart 290} {liver 1494} {pancreas 131} {spleen 162}]
}

// Make the nil value useful
type StringsStack struct{ data []string }

func (s *StringsStack) Push(x string) {
	s.data = append(s.data, x)
}

func (s *StringsStack) Pop() string {
	if l := len(s.data); l > 0 {
		t := s.data[l-1]
		s.data = s.data[:l-1]
		return t
	}
	panic("pop from empty stack")
}

// Nil as a receiver value
//
//	Nothing in Go prevents calling a method with a nil receiver
type IntList struct {
	Value int
	Tail  *IntList
}

func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}

func main() {
	exercise1() // /usr 121
	exercise2() // Hash of /usr is 0xfdfe
	exercise3() // Hash of /usr/lib is 0xfgfg
	exercise4() // Hash of /usr/lib is 0xfgfg; length 120
	exercise5() // usr
	exercise6() // lib
	exercise7() // Hash of /usr is 0xfex7; length 120
	exercise8() // [{brain 1340} {liver 1494} {spleen 162} {pancreas 131} {heart 290}]
	exercise9()
}
