package main

import (
	"log"
	"net/http"
	"time"
)

// Channels
// 	A channel is a one way communications pipe
// 	Things go in one end, come out the other
// 	In the same order they went in
// 	Until the channel is closed
//  Multiple readers & writers can share it safely

// Sequential process
// 	Looking at a single independent part of the program
// 	it appears to be sequential
// 	for {
// 		read()
// 		process()
// 		write()
// 	}
// 	This is perfectly natural if we think of reading & writing
// 	files or network sockets

// Communicating sequential processes (CSP)
// 	Put the parts together with channels to communicate
// 	Each part is independent
// 	All they share are the channels between them
// 	The parts can run in parallel as the hardware allows

// 	Concurrency is always hard
// 	CSP provides a model for thinking about it that makes it
//  less hard (take the program apart and make the pieces talk
// 	to each other)
// 	Go doesn't force devs to embrace the asynchronous ways of
// 	event-driven programming. That lets you write asynchronous
//  code in a synchronous style. As people, we're much better
// 	suited to writing about things in a synchronous style.

// Goroutines
// 	A goroutine is a unit of independent execution (coroutine)
// 	It's easy to start a goroutine: put 'go' in front of a func call
// 	The trick is knowing how the goroutine will stop:
// 		You have a well-defined loop terminating condition, or
// 		You signal completion through a channel or context, or
// 		You let it run until the program stops
// 	But you need to make sure it doesn't get blocked by mistake

// Channels
// 	A channel is like a one-way socket or a Unix pipe
// 	(except it allows multiple readers & writers)
// 	It's a method of synchronization as well as communication
// 	We know that a send (write) always happens before a receive (read)
// 	It's also a vehicle for transferring ownership of data, so that
// 	only one goroutine at a time is writing the data (avoid race conditions)
// 	Don't communicate by sharing memory; instead, share memory by
// 	communicating - Rob Pike

type result struct {
	url     string
	err     error
	latency time.Duration
}

func get(url string, ch chan<- result) {
	start := time.Now()

	if resp, err := http.Get(url); err != nil {
		ch <- result{url, err, 0}
	} else {
		t := time.Since(start).Round(time.Millisecond)
		ch <- result{url, err, t}
		resp.Body.Close()
	}
}

func main() {
	results := make(chan result)
	list := []string{
		"https://amazon.co.uk",
		"https://nba.com",
		"https://nytimes.com",
	}

	for _, url := range list {
		go get(url, results)
		// a go routine wants a function call
	}

	for range list {
		r := <-results

		if r.err != nil {
			log.Printf("%-20s %s\n", r.url, r.err)
		} else {
			log.Printf("%-20s %s\n", r.url, r.latency)
		}
	}
	// time go run main.go
	// 2024/08/11 07:53:51 https://amazon.co.uk 609ms
	// 2024/08/11 07:53:51 https://nytimes.com  622ms
	// 2024/08/11 07:53:52 https://nba.com      1.977s

	// real    0m8.151s
	// user    0m0.015s
	// sys     0m0.153s

}
