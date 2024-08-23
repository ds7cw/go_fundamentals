package main

import (
	"fmt"
	"sync"
	"time"
)

// Concurrency problems
// #1: race conditions, where unprotected reads & writes overlap
// 	must be some data that is written to
// 	could be a read-modify-write operation
// 	and two goroutines can do it at the same time

// #2: deadlock, when no goroutine can make progress
// 	goroutines could all be blocked on empty channels
// 	goroutines could all be blocked waiting on a mutex
// 	GC could be prevented from running (busy loop)

// Go detects some deadlocks automatically; with -race it can find
// some data races

// #3: goroutine leak
// 	goroutine hangs on an empty or blocked channel
//  not deadlock; other goroutines make progress
// 	often found by looking at pprof output
// When you start a goroutine, always know how/when it will end

// #4: channel errors
// 	trying to send on a closed channel
// 	trying to send or receive on a nil channel
//  closing a nil channel
// 	closing a channel twice

// #5: other errors
//  closure capture
//  misuse of Mutex
//  misuse of WaitGroup
//  misuse of select
// Many of the errors are basic & should be found by review;
// maybe we'll get statis analysis tools to help find them

// Gotchas 1: Data race
// var nextID = 0

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "<h1>You got %v<h1>", nextID)
// 	nextID++ // unsafe - data race
// }

func main() {
	// Gotchas 1: Data race
	// http.HandleFunc("/", handler)
	// if err := http.ListenAndServe(":8080", nil); err != nil {
	// 	log.Fatal(err)
	// }

	// Deadlock
	// ch := make(chan bool)
	// go func(ok bool) {
	// 	fmt.Println("START")
	// 	if ok {
	// 		ch <- ok
	// 	}
	// }(false)

	// <-ch
	// fmt.println("DONE")

	//
	var m sync.Mutex
	done := make(chan bool)
	fmt.Println("START")
	go func() {
		m.Lock()
		defer m.Unlock() // don't forget to unlock!!!
	}()
	go func() {
		time.Sleep(1)
		m.Lock()
		defer m.Unlock()
		fmt.Println("SIGNAL")
		done <- true
	}()
	<-done
	fmt.Println("DONE")

}
