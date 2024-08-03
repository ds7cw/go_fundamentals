package main

import "fmt"

// Pointers - shared, not copied
// Values - copied, not shared

// Value semantics lead to higher integrity, particularly with concurrency
// Pointer semantics may be more efficient

// Common uses of pointers:
// Some objects can't be copied safely (mutex)
// Some objects are too large to copy efficiently
// (consider pointers if size > 64 bytes)
// Some methods need to change (mutate) the receiver [later!]
// When decoding protocol data into an object
// (JSON, etc.; often in a variable argument list)
// var Response
// err := json.Unmarshal(j, &r)
// When using a pointer to signal a "null" object

// Any struct with a mutex must be passed by reference:
// type Employee struct {
// 	mu sync.Mutex
// 	Name string
// }
// func do(emp *Employee) {
// 	emp.mu.Lock()
// 	defer emp.mu.Unlock()
// }

// Copying is okay
// Any small struct under 64 bytes probably should be copied:
// type Widget struct {
// 	ID int
// 	Count int
// }
// func Expend(w Widget) Widget {
// 	w.Count--
// 	return w
// }
// Note that Go routinely copies string & slice descriptors

// If something is to be shared, then always pass a pointer
// Employee relocation
// f1(emp *Employee)
// f2(emp *Employee)
// f3(emp Employee)		passes a copy
// f4(emp *Employee)	changes are LOST

// Stack allocation is more efficient
// Accessing a var directly is more efficient than following a pointer
// Accessing a dense sequence of data is more efficient than sparse
// data (an array is faster than a linked list, etc.)

// Heap allocation
// Go would prefer to allocate on the stack, but sometimes can't
// a function returns a pointer to a local object
// a local object is captured in a function closure
// a pointer to a local object is sent via a channel
// any object is assigned in an interface
// any object whose size is variable at runtime (slices)

// The use of new has nothing to do with it
// Build with the flag -gcflags -m=2 to see the escape analysis

// For loops
// The value returned by range is always a copy
// for i, thing := range things { thing is a copy}
// Use the index if you need to mutate the element
// for i := range things { things[i].which = whatever }
// the change in the 2nd loop will be visible after the loop

// Slice safety
// Anytime a func mutates a slice that's passed in,
// we must return a copy
// func update(things []thing) []thing {
// 	things = append(things, x)
// 	return things
// }
// That's because the slice's backing array may be
// reallocated to grow

// Keeping a pointer to an element of a slice is risky
type user struct {
	name  string
	count int
}

func addTo(u *user) { u.count++ }

func main() {
	users := []user{{"alice", 0}, {"bob", 0}}
	alice := &users[0] // risky
	amy := user{"amy", 1}
	users = append(users, amy)

	addTo(alice)       // alice is likely a stale pointer
	fmt.Println(users) // so alice's count will be 0
	// [{alice 0} {bob 0} {amy 1}]

	items := [][2]byte{{1, 2}, {3, 4}, {5, 6}}
	a := [][]byte{}

	for _, item := range items {
		a = append(a, item[:])
	}
	fmt.Println(items) // [[1 2] [3 4] [5 6]]
	fmt.Println(a)     // [[1 2] [3 4] [5 6]]

}

// Taking the address of a mutating loop variable is wrong
// func (r OfferResolver) Changes() []ChangeResolver {
// 	var result []ChangeResolver
// 	// wrong
// 	for _, change := range r.d.Status.Changes {
// 		result = append(result, ChangeResolver{&change}) // WRONG
// 	}
// 	return result
// }
// Wrong: all the returned resolvers point to the last change
// in the list
