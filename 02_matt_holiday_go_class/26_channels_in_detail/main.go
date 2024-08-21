package main

import (
	"fmt"
	"time"
)

// Channels block unless ready to read or write
// A channel is ready to write if:
// 	It has buffer space, or
// 	At least one reader is ready to read

// A channel is ready to read if:
// 	It has unread data in its buffer, or
// 	At least one writer is ready to write, or
// 	It is closed

// Channels are unidirectional, but have two ends
// 	(which can be passed separately as params)

// An end for writing & closing
// 	func get(url string, ch chan<- result) { . . . } // write-only end

// An end for reading
// 	func collect(ch chan<- result) map[string]int { . . . } // read-only end

// Closing a channel causes it to return the "zero" value
// We can receive a second value: is the channel closed?

// A channel can only be closed once (else it will panic)
// One of the main issues in working with goroutines is ending them
// 	An unbuffered channel requires a reader and writer
// 	(a writer blocked on a channel with no reader will "leak")
// 	Closing a channel is often a signal that work is done
// 	Only one goroutine can close a channel (not many)
// 	We may need some way to coordinate closing a channel or stopping
// 	goroutines (beyond the channel itself)

// Nil channels
// 	Reading or writing a channel that is nil always blocks*
// 	But a nil channel in a select block is ignored
// 	This can be a powerful tool:
// 		Use a channel to get input
// 		Suspend it by changing the channel variable to nil
// 		You can even un-suspend it again
// 		But close the channel if there really is no more input (EOF)

// Channel state reference
// State		Receive			Send	Close
// ------------+---------------+-------+---------------------+
// Nil			Block*			Block*	Panic
// Empty		Block			Write	Close
// Partly Full	Read			Write	Readable until empty
// Full			Read			Block	Readable until empty
// Closed		Default Value**	Panic	Panic
// ------------+---------------+-----------------------------+
// Receive-only	OK				Compile Error
// Send-only	Compile Error	OK
// ------------+---------------+-------+---------------------+
// *  select ignores a nil channel since it would always block
// ** Reading a closed channel returns (default-value, !ok)

// Rendezvous
// 	By default, channels are unbuffered (rendezvous model)
// 		the sender blocks until the receiver is ready (and vice versa)
// 		the send always happens before the receive
// 		the receive always returns before the send
// 		the sender & receiver are synchronized

// Buffering
// 	Buffering allows the sender to send without waiting
// 		the sender deposits its item and returns immediately
//		the sender blocks only if the buffer is full
// 		the receiver blocks only if the buffer is empty
// 		the sender & receiver run independently

func main() {
	ch1 := make(chan int, 1)
	ch1 <- 1

	b1, ok1 := <-ch1
	fmt.Println(b1, ok1) // 1 true

	// c1, ok2 := <-ch1
	// fmt.Println(c1, ok2) // 0 true

	// Buffering
	// Make a channel with buffer that holds 2 items
	messages := make(chan string, 2)

	// Now we can send twice without getting blocked
	messages <- "buffered"
	messages <- "channel"

	// And then receive both as usual
	fmt.Println(<-messages) // buffered
	fmt.Println(<-messages) // channel

	exampleFunc01()
	// {0 false}
	// {2 false}
	// {1 false}
	// {3 false}
	// {4 false}

	// Common uses of buffered channels:
	// 	avoid goroutine leaks (from abandoned channel)
	//	avoid randezvous pauses (performance improvement)
	// Don't buffer until it's needed: buffering may hide a race condition
	// Some testing may be required to find the right number of slots
	// Special uses of buffered channels:
	// 	counting semaphore pattern

}

type T struct {
	i byte
	b bool
}

func send(i int, ch chan<- *T) {
	t := &T{i: byte(i)}
	ch <- t

	t.b = true // UNSAFE, BAD PRACTICE

}

func exampleFunc01() {
	vs := make([]T, 5)
	ch := make(chan *T) // ch := make(chan *T, 5) all 5 return true

	for i := range vs {
		go send(i, ch)
	}

	time.Sleep(1 * time.Second) // all goroutines started

	// copy quickly
	for i := range vs {
		vs[i] = *<-ch
	}

	// print later
	for _, v := range vs {
		fmt.Println(v)
	}
}
