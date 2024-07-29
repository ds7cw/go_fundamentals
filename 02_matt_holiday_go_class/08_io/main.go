package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Package os has functions to open or create files,
// list directories, etc. and hosts the File type
// 'os' has utilities to read and write;
// 'bufio' provides the buffed I/O scanners, etc.

// Package io/ioutil has extra utilities such as reading
// an entire file to memory, or writing it out all at once

// Package strconv has utilities to convert to/ from
// string representations

func exercise1() {
	for _, fname := range os.Args[1:] {
		file, err := os.Open(fname)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		if _, err := io.Copy(os.Stdout, file); err != nil {
			fmt.Fprint(os.Stderr, err)
		}

		file.Close()

	}
}

func exercise2() {
	for _, fname := range os.Args[1:] {
		file, err := os.Open(fname)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		data, err := io.ReadAll(file)

		if err != nil {
			fmt.Fprint(os.Stderr, err)
			continue
		}

		fmt.Println("The file has", len(data), "bytes")
		file.Close()

	}
}

func exercise3() {
	for _, fname := range os.Args[1:] {
		var lc, wc, cc int

		file, err := os.Open(fname)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		scan := bufio.NewScanner(file)

		for scan.Scan() {
			s := scan.Text()

			wc += len(strings.Fields(s))
			cc += len(s)
			lc++
		}

		fmt.Printf("%7d %7d %7d %s\n", lc, wc, cc, fname)
		file.Close()

	}
}

func main() {

	// We often call funcs whose 2nd return val is a possible error
	// func Open(name string) (*File, error)
	// where the error can be compared to nil, meaning no error
	// Always check the error - the file might not really be open!

	// go run main.go a.txt
	exercise1() // a file line of text
	exercise2() // The file has 21 bytes
	exercise3() //       1       5      19 a.txt

}
