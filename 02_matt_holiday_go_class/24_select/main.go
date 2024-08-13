package main

import (
	"log"
	"time"
)

// Select
// 	Allows any "ready" alternative to proceed among:
// 		a channel we can read from
// 		a channel we can write to
// 		a default action that's always ready
// 	Most often select runs in a loop so we keep trying
// 	We can put a timeout or done channel into the select
// 	We can compose channels as synchronization primitives
// 	Traditional primitives (mutex, condition var) can't be
// 	composed

func tick() {
	log.Println("start")

	const tickRate = 2 * time.Second

	stopper := time.After(5 * tickRate)
	ticker := time.NewTicker(tickRate).C

loop:
	for {
		select {
		case <-ticker:
			log.Println("tick")
		case <-stopper:
			break loop
		}
	}

	log.Println("finish")
	// 2024/08/13 19:17:01 start
	// 2024/08/13 19:17:03 tick
	// 2024/08/13 19:17:05 tick
	// 2024/08/13 19:17:07 tick
	// 2024/08/13 19:17:09 tick
	// 2024/08/13 19:17:11 tick
	// 2024/08/13 19:17:11 finish
}

// Select: default
// 	In a select block, the default case is always ready and
// 	will be chosen if no other case is:
// func sendOrDrop(data []byte) {
// 	select {
// 	case ch <- data;
// 		// sent ok; do nothing
// 	default:
// 		log.Printf("overflow: drop %d bytes", len(data))
// 	}
// }
// 	Dont use default inside a loop, the select will busy wait
// 	and waste CPU

func main() {
	chans := []chan int{
		make(chan int),
		make(chan int),
	}

	for i := range chans {
		go func(i int, ch chan<- int) {
			for {
				time.Sleep(time.Duration(i) * time.Second)
				ch <- i
			}
		}(i+1, chans[i])
	}

	for i := 0; i < 12; i++ {
		select {
		case m0 := <-chans[0]:
			log.Println("received", m0)
		case m1 := <-chans[1]:
			log.Println("received", m1)
		}
		// 2024/08/13 19:01:07 received 1
		// 2024/08/13 19:01:08 received 2
		// 2024/08/13 19:01:08 received 1
		// 2024/08/13 19:01:09 received 1
		// 2024/08/13 19:01:10 received 2
		// 2024/08/13 19:01:10 received 1
		// 2024/08/13 19:01:11 received 1
		// 2024/08/13 19:01:12 received 2
		// 2024/08/13 19:01:12 received 1
		// 2024/08/13 19:01:13 received 1
		// 2024/08/13 19:01:14 received 2
		// 2024/08/13 19:01:14 received 1
	}

	tick()

}
