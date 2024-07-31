package main

import (
	"fmt"
	"time"
)

type Employee struct {
	Name   string
	Number int
	Boss   *Employee
	Hired  time.Time
}

func main() {
	c := map[string]*Employee{}

	c["Kobe"] = &Employee{"Kobe", 2, nil, time.Now()}

	// var e Employee
	// e := Employee{Name: "Shaq", Number: 1, Hired: time.Now()}
	// Another way of declaring (above)
	c["Shaq"] = &Employee{
		Name:   "Shaq",
		Number: 1,
		Boss:   c["Kobe"],
		Hired:  time.Now(),
	}

	fmt.Printf("%T %[1]v\n", c["Shaq"])
	// main.Employee { 0 <nil> 0001-01-01 00:00:00 +0000 UTC}
	// *main.Employee &{Shaq 1 0xc00003c040 2024-07-31 06:53:47}

	fmt.Printf("%T %+[1]v\n", c["Kobe"])
	// *main.Employee &{Name:Kobe Number:2 Boss:<nil> Hired:2024-07-31 06:53:47}

	fmt.Printf("%T %+[1]v\n", c)
	// map[string]*main.Employee map[Kobe:0xc00003c040 Shaq:0xc00003c080]

	c["Kobe"].Number++            // Update from 2 to 3
	fmt.Println(c["Kobe"].Number) // 3

	var album = struct {
		title string
	}{
		"The White Album",
	}
	fmt.Println(album) // {The White Album The Beatles 1968 100000000}

	type album1 struct {
		title string
	}
	type album2 struct {
		title string
	}
	a1 := album1{"The White Album"}
	a2 := album2{"The Black Album"}
	// a1 = a2 // can't use a2 (type album2) as type album1 in assignment
	a1 = album1(a2) // type conversion from one struct to another
	fmt.Println("a1:", a1, "a2:", a2)
	// a1: {The Black Album} a2: {The Black Album}

	// 2 struct types are compatible if
	// the fields have the same types and names
	// the fields are in the same order
	// with the same tags (*)

	// A struct may be copied or passed as a param in its entirety
	// A struct is comparable if all its fields are comparable
	// The zero value for a struct is "zero" for each field in turn

	// Make the zero value useful
	// It's usually desireable that the 0 val be a natural or sensible default
	// For example, in bytes. Buffer, the initial val of the struct is ready-
	// to-use empty buffer.
	// type Buffer struct {
	// 	buf      []byte // contents are the bytes buf[off: len(buf)]
	// 	off      int    // read at &buf[off], write at &buff[len(buf)]
	// 	lastRead read0p // last read operation, so that Unread* can work correctly
	// }
	// which has a nil slice we can append to, and off starts as 0;
	// the 0 value for read0p is opInvalid

	// r := &Response{Page: 1, Words: []string{"up", "in", "out"}}
	// j, _ := json.Marshal(r)
	// fmt.Printf("%#v\n", r)
	// // &main.Response{Page:1, Words:[]string{"up", "in", "out"}}
	// fmt.Println(string(j)) // {"Page":1,"Words":["up","in","out"]}

	// var r2 Response
	// _ = json.Unmarshal(j, &r2)
	// fmt.Printf("%#v\n", r2)
}

// type Response struct {
// 	Page  int      `json:page`
// 	Words []string `json:words,omitempty`
// }

// Field names should start with capital letter for json export to work
