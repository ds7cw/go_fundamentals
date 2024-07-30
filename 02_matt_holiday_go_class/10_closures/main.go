package main

import "fmt"

// Scope is static, based on the code at compile time
// Lifetime depends on program execution (runtime)
// package xyz
// func doIt() *int {
// 	var b int . . .
// 	return &b
// }
// Var 'b' can only be seen inside doIt, but its value will live on
// The value (objet) will live so long as part of the program
// keeps a pointer to it

// A closure is when a func inside another func "closes over"
// one or more local vars of the outer func
func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return b
	}
}

// The inner func gets a reference to the outer func's var
// Those vars may end up with a much longer lifetime than expected
// as long as there's a reference to the inner func

func do(d func()) {
	d()
}

func main() {
	f := fib()
	// As long as this 'f' exists, the values of 'a' & 'b'
	// from fib will be reused every time 'f' is called
	for x := f(); x < 100; x = f() {
		fmt.Println(x) // 1 2 3 5 8 13 21 34 55 89
	}

	g, h := fib(), fib()
	fmt.Println(g(), g(), g(), g()) // 1 2 3 5
	fmt.Println(h(), h(), h(), h()) // 1 2 3 5

	// type kv struct {
	// 	key string
	// 	val int
	// }
	// var ss []kv
	// for k, v := range words {
	// 	ss = append(ss, kv{k, v})
	// }
	// sort.Slice(ss, func(i, j int) bool {
	// 	return ss[i].val > ss[j].val
	// })
	// ss = ss[:3]
	// for _, s := range ss {
	// 	fmt.Println(s.key, "appears", s.val, "times")
	// }

	for i := 0; i < 4; i++ {
		v := func() {
			fmt.Printf("%d @ %p\n", i, &i)
		}
		do(v)
		// 0 @ 0xc00000a118
		// 1 @ 0xc00000a150
		// 2 @ 0xc00000a158
		// 3 @ 0xc00000a160
	}

	s1 := make([]func(), 4)
	for i := 0; i < 4; i++ {
		s1[i] = func() {
			fmt.Printf("%d @ %p\n", i, &i)
		}
	}
	for i := 0; i < 4; i++ {
		s1[i]()
		// 0 @ 0xc000112120
		// 1 @ 0xc000112128
		// 2 @ 0xc000112130
		// 3 @ 0xc000112138
	}
}
