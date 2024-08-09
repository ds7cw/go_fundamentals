package main

import (
	"fmt"
	"time"
)

// Concurrency definitions
// 	Execution happens in some non-deterministic order
// 	Undefined out-of-order execution
// 	Non-sequential execution
// 	Parts of a program execute out-of-order or in
// 	partial order

// Partial order
// 	Part 1 happens before parts 2 or 3
//  Both 2 and 3 complete before part 4
// 	The parts of 2 and 3 are ordered among themselves

// Non-deterministic
// 	Different behaviors on different runs, even
// 	with same input
// 	We don't necessarily mean different results, but
// 	a different trace of execution

// With the above in mind, concurrency definition:
// 	Parts of the program may execute independently
// 	in some non-deterministic (partial) order

// Parallelism
// 	Parts of a program execute independently at the same time
// 	You can have concurrency with a single-core cpu
// 	Parallelism can happen only on a multi-core cpu
// 	Concurrency doesn't make the program faster,
// 	parallelism does

// Concurrency vs Parallelism
// 	Concurrency is about dealing with things happening out-of-order
//  Parallelism is about things happening at the same time
// 	A single program won't have parallelism without concurrency
// 	We need concurrency to allow parts of the program to execute
// 	independently

// Race condition
// 	System behavior depends on the (non-deterministic) sequence
// 	or timing of parts of the program executing independently,
// 	where some possible behaviors (orders of execution) produce
// 	invalid results

// Ways to solve race conditions
//  Race conditions involve independent parts of the program
// 	changing things that are shared
// 	Don't share anything
// 	Make shared things read-only
// 	Allow only one writer to the shared things
// 	Make the read-modify-write operations atomic

// Concurrency is about dealing with multiple tasks at once.
// Think of it like a chef preparing multiple dishes at the
// same time. The chef might start cooking one dish, then while
// it's simmering, start preparing another dish. The chef is
// switching between tasks, but not necessarily doing them at
// the exact same time.

func cookDish(dish string) {
	for i := 1; i <= 3; i++ {
		fmt.Printf("Cooking %s: step %d\n", dish, i)
		time.Sleep(time.Second)
	}
}

// Parallelism, on the other hand, is about doing multiple tasks
// at the same time. Imagine two chefs in the kitchen, each
// cooking a different dish simultaneously.

func main() {
	go cookDish("Pasta")
	go cookDish("Salad")

	time.Sleep(4 * time.Second) // Wait for goroutines to finish
	fmt.Println("All dishes are prepared!")

	// Cooking Salad: step 1
	// Cooking Pasta: step 1
	// Cooking Pasta: step 2
	// Cooking Salad: step 2
	// Cooking Salad: step 3
	// Cooking Pasta: step 3
	// All dishes are prepared!
}
