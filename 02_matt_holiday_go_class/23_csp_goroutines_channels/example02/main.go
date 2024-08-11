package main

import "fmt"

func generate(limit int, ch chan<- int) {
	for i := 2; i < limit; i++ {
		ch <- i
	}

	close(ch)
}

func filter(src <-chan int, dst chan<- int, prime int) {
	for i := range src {
		if i%prime != 0 {
			dst <- i
		}
	}

	close(dst)
}

func sieve(limit int) {
	ch := make(chan int)

	go generate(limit, ch)

	for {
		prime, ok := <-ch

		if !ok {
			break
		}

		ch1 := make(chan int)

		go filter(ch, ch1, prime)

		ch = ch1

		fmt.Print(prime, " ")
	}

	fmt.Println()
}

func main() {
	sieve(100)
	// 2 3 5 7 11 13 17 19 23 29 31
	// 37 41 43 47 53 59 61 67 71
	// 73 79 83 89 97
}
