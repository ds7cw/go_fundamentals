package main

import (
	"fmt"
	"log"
	"net/http"
)

//  OPTION 1
// var nextID = make(chan int)

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "<h1>You got %d<h1>", <-nextID)
// }

// func counter() {
// 	for i := 0; ; i++ {
// 		nextID <- i
// 	}
// }

// OPTION 2
type nextCh chan int

func (ch nextCh) handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>You got %d<h1>", <-ch)
}

func counter(ch chan<- int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func main() {
	//  OPTION 1
	// go counter()
	// http.HandleFunc("/", handler)

	// OPTION 2
	var nextID nextCh = make(chan int)
	go counter(nextID)
	http.HandleFunc("/", nextID.handler)

	// Used in both cases
	log.Fatal(http.ListenAndServe(":8080", nil))

}
