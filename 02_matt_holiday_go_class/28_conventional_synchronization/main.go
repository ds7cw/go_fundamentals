package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Package sync, for example:
// 	Mutex
// 	Once
// 	Pool
// 	RWMutex
// 	WaitGroup

// Package sync/atomic for atomic scalar reads & writes

// Mutual exclusion
// 	What if multiple goroutines must read & write some data?
// 	We must make sure only one of them can do so at any instant
// 	(in the so-called "critical section")

// We accomplish this with some type of lock
// 	acquire the lock before accessing the data
// 	any other goroutine will block waiting to get the lock
// 	release the lock when done

func example01() int {
	var n int64
	var w sync.WaitGroup

	for i := 0; i < 1000; i++ {
		w.Add(1)

		go func() {
			n++ // DATA RACE
			w.Done()
		}()
	}

	w.Wait()
	return int(n)
}

func example01Update() int {
	// m := make(chan bool, 1)
	var m sync.Mutex

	var n int64
	var w sync.WaitGroup

	for i := 0; i < 1000; i++ {
		w.Add(1)

		go func() {
			// m <- true
			m.Lock()

			n++ // DATA RACE
			// <-m
			m.Unlock()

			w.Done()
		}()
	}

	w.Wait()
	return int(n)
}

// Mutexes in action
type SafeMap struct {
	sync.Mutex // not safe to copy
	m          map[string]int
}

// so methods must take a pointer, not a value
func (s *SafeMap) example02(key string) {
	s.Lock()
	defer s.Unlock() // using defer is a good habit

	// only one goroutine can execute this
	// code at the same time, guaranteed
	s.m[key]++
}

// RWMutexes in action
type InfoClient struct {
	mu        sync.RWMutex
	token     string
	tokenTime time.Time
	TTL       time.Duration
}

// CheckToken func
func (i *InfoClient) example03Check() (string, time.Duration) {
	i.mu.RLock()
	defer i.mu.RUnlock()

	return i.token, i.TTL - time.Since(i.tokenTime)
}

// func (i *InfoClient) example03Replace(ctx context.Context) (string, error) {
// 	token, ttl, err := i.getAccessToken(ctx)

// 	if err != nil {
// 		return "", err
// 	}

// 	i.mu.Lock()
// 	defer i.mu.Unlock()

// 	i.token = token
// 	i.tokenTime = time.Now()
// 	i.TTL = time.Duration(ttl) * time.Second

// 	return token, nil
// }

// Atomic primitives
func example04() int {
	var n int64
	var w sync.WaitGroup

	for i := 0; i < 1000; i++ {
		w.Add(1)

		go func() {
			atomic.AddInt64(&n, 1) // fxed
			w.Done()
		}()
	}

	w.Wait()
	return int(n)
}

// Only-once execution
// 	A sync.Once obj allows us to ensure a func runs only once
// 	(only the first call to Do will call the func passed in)
// var once sync.Once
// var x *singleton
// func initialize() {
// 	x = NewSingleton()
// }
// func handle(w http.ResponseWriter, r *http.Request) {
// 	once.Do(initialize)
// 	. . . // checking x == nil in the handler is UNSAFE
// }

// Pool
//
//	A Pool provides for efficient & safe reuse of objects,
//	but it's a container of interface{}
// var bufPool = sync.Pool{
// 	New: func() interface{} {
// 		return new(bytes.Buffer)
// 	},
// }

// func example05Log(w io.Writer, key, val string) {
// 	b := bufPool.Get().(*bytes.Buffer) // more reflection
// 	b.Reset()
// 	// write to it
// 	w.Write(b.Bytes())
// 	bufPool.Put(b)
// }

// Other primitives
// 	Condition variable
// 	Map (safe container; uses interface{})
// 	WaitGroup

func main() {
	fmt.Println(example01())       // 983 // different every time
	fmt.Println(example01Update()) // 1000 // always 1000
	fmt.Println(example04())       // 1000 // always 1000
}
