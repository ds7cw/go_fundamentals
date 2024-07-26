package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "not enough args")
		os.Exit(-1)
	}

	old, new := os.Args[1], os.Args[2]
	scan := bufio.NewScanner(os.Stdin)

	for scan.Scan() {
		s := strings.Split(scan.Text(), old)
		t := strings.Join(s, new)

		fmt.Println(t)
	}
	// $ go run main.go matt ed < test.txt

	// matt went to greece
	// where did matt go
	// alan went to rome
	// matt didn't go there

	// ed went to greece
	// where did ed go
	// alan went to rome
	// ed didn't go there

	// // Strings vs runes
	// s := "élite"
	// b := []byte(s)

	// fmt.Printf("%8T %[1]v %d\n", s, len(s)) // string élite 6
	// fmt.Printf("%8T %[1]v\n", []rune(s))    // []int32 [233 108 105 116 101]
	// fmt.Printf("%8T %[1]v %d\n", b, len(b)) // []uint8 [195 169 108 105 116 101] 6

	// str := "the quick brown fox"
	// a := len(str)
	// b1 := str[:3]
	// c := str[4:9]
	// d := str[:4] + "slow" + str[9:]

	// // str[5] = "a" // SYNTAX ERROR
	// fmt.Println("Address of 'str':", &str) // Address of 'str': 0xc00008a040
	// str += "es"
	// fmt.Println("Address of updated 'str':", &str) // Address of updated 'str': 0xc00008a040

	// fmt.Println(a)   // 19
	// fmt.Println(b1)  // the
	// fmt.Println(c)   // quick
	// fmt.Println(d)   // the slow brown fox
	// fmt.Println(str) // the quick brown foxes

	// fmt.Println(strings.Contains(str, "w"))  // true
	// fmt.Println(strings.Contains(str, "j"))  // false
	// fmt.Println(strings.HasPrefix(str, "t")) // true
	// fmt.Println(strings.Index(str, "h"))     // 1

	// str = strings.ToUpper(str)
	// fmt.Println(str)  // THE QUICK BROWN FOXES
	// fmt.Println(&str) // 0xc000032080

}
