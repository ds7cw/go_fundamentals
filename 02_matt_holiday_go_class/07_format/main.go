package main

import (
	"fmt"
	"os"
)

// Unix has the notion of 3 standard I/O streams:
// Standard input
// Standard output
// Standard error (output)

func main() {
	fmt.Println("printing to standard output")
	fmt.Fprintln(os.Stderr, "printing to error output")

	// Always os.Stdout
	// fmt.Println(...interface{}) (int, error)
	// fmt.Printf(string, ...interface{}) (int, error)

	// Print to anything that has the correct Write() method
	// fmt.Fprintln(io.Writer, ...interface{}) (int, error)
	// fmt.Fprintf(io.Writer, string,...interface{}) (int, error)

	// Return a string
	// fmt.Sprintln(...interface{}) string
	// fmt.Sprintf(string, ...interface{}) string

	// %s  the uninterpreted bytes of the string or slice
	// %q  a double-quoted string safely escaped with Go syntax
	// %c  the character represented by the corresponding Unicode code point

	// %d  base 10
	// %x  base 16, with lower-case letters for a-f
	// %f  decimal point but no exponent, e.g. 123,456 (%.2f)
	// %t  the word true or false

	// %v  the value in a default format
	//     when printing structs, the + flag (%+v) adds field names

	// %#v a Go-syntax representation of the value
	// %T  a Go-syntax representation of the type

	// %%  a literal percent sign; consumes no value [escape]

	a, b := 12, 345
	c, d := 1.2, 3.45
	fmt.Printf("%d %d\n", a, b)       // 12 345
	fmt.Printf("%X %X\n", a, b)       // C 159
	fmt.Printf("%#x %#x\n", a, b)     // 0xc 0x159
	fmt.Printf("%f %.2f\n", c, d)     // 1.200000 3.45
	fmt.Printf("|%6d|%6d|\n", a, b)   // |    12|   345|
	fmt.Printf("|%06d|%06d|\n", a, b) // |000012|000345|
	fmt.Printf("|%-6d|%-6d|\n", a, b) // |12    |345   |

	// Slices
	s1 := []int{1, 2, 3}
	fmt.Printf("%T\n", s1)  // []int
	fmt.Printf("%v\n", s1)  // [1 2 3]
	fmt.Printf("%#v\n", s1) // []int{1, 2, 3}

	// Arrays
	a1 := [3]rune{'a', 'b', 'c'}
	fmt.Printf("%T\n", a1)  // [3]int32
	fmt.Printf("%v\n", a1)  // [97 98 99]
	fmt.Printf("%#v\n", a1) // [3]int32{97, 98, 99}
	fmt.Printf("%q\n", a1)  // ['a' 'b' 'c']

	// Maps
	m1 := map[string]int{"and": 1, "or": 2}
	fmt.Printf("%T\n", m1)  // map[string]int
	fmt.Printf("%v\n", m1)  // map[and:1 or:2]
	fmt.Printf("%#v\n", m1) // map[string]int{"and":1, "or":2}

	// String
	str1 := "a string"
	b1 := []byte(str1)
	fmt.Printf("%T\n", str1)  // string
	fmt.Printf("%v\n", str1)  // a string
	fmt.Printf("%#v\n", str1) // "a string"
	fmt.Printf("%q\n", str1)  // "a string"
	fmt.Printf("%v\n", b1)    // [97 32 115 116 114 105 110 103]

}
