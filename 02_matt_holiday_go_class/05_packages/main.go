package main // Every standalone program has a main package

import "fmt"

// Nothing is 'global'; it's either in your package or in another
// It's either at package scope or function
// The goal of packages is to achieve a level of encapsulation

// Every name that's capitalized is exported
// func GetAll(space, name string) (map[string]string, error) { ... }
// That means another package in the program can import it
// (within a package, everything is visible even across files)

// Each source file in your package must import what it needs
// package secrets
// import (
// 	"encoding/base64"
// 	"encoding/json"
// 	"fmt"
// 	"os"
// 	"strings"
// )
// It may only import what it needs; unused imports are an error
// Generally, flies of the same package live together in  dir

// A package 'A' can't import a package 'B' that imports A
// Move common dependencies to a 3rd package, or eliminate them

// Items within a packaage get initialized before main

// A package should embed deep functionality behind a simple API
// package os
// func Create(name string) (*File, error)
// func Open(name string) (*File, error)

// func (f *File) Read(b []byte) (n int, err error)
// func (f *File) Write(b []byte) (n int, err error)
// func (f *File) Close() error

// The Unix file API is perhaps the best example of this model
// Roughly five functions hide a lot of complexity from the user

func main() {
	fmt.Println("Hello")
}
