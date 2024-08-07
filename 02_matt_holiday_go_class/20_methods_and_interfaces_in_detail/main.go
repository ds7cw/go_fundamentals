package main

import (
	"fmt"
	"math"
)

// An interface variable is nil until initialized
// It really has two parts:
// 		A value or pointer of some type
// 		A pointer to type info so the correct actual
// 		method can be identified
// var r io.Reader // nil until initialized
// var b *bytes.Buffer // nil until initialized
// r = b // r is no longer nil!
// 		 // but it has a nil pointer to a Buffer
// This may confuse; an interface var is nil only if
// 	both parts are

// Error is really an interface
//	We called error a special type, but it's really an interface
// 		type error interface { func Error() string }
// 	We can compare it to nil unless we make a mistake
// 	The mistake is to store a nil pointer to a concrete type
// 	in the error variable

type errFoo struct {
	err  error
	path string
}

func (e errFoo) Error() string {
	return fmt.Sprintf("%s: %s", e.path, e.err)
}

func XYZ(a int) error {
	return nil
}

// Pointer vs value receivers
// A method can be defined on a pointer type
type Point struct {
	x, y float64
}

func (p *Point) Add(x, y float64) {
	p.x, p.y = p.x+x, p.y+y
}

func (p *Point) OffsetOf(p1 Point) (x float64, y float64) {
	x, y = p.x-p1.x, p.y-p1.y
	return
}

// The same method name may not be bound to both T and *T
// Pointer methods may be called on non-pointers and vice versa
// Go will automatically use * or & as needed
// p1 := new(Point) // *Point, at (0, 0)
// p2 := Point{1, 1}
// p1.OffsetOf(p2) // same as (*p1), OffsetOf(p2)
// p2.Add(3, 4) // same as (&p2). Add(3, 4)

// Consistency in receiver types
// 	If one method of a type takes a pointer receiver, then all
// 	its methods should take pointers*
// 	And in general objects of that type are probably not safe to copy
// type Buffer struct {
// 	buf []byte
// 	off int
// }
// func (b *Buffer) ReadString(delim byte) (string, error) { . . . }
// *Except when they shouldn't for other reasons

// Currying takes a func and reduces its argument count by one
// 	(one argument gets bound, and new func is returned)
// 	func Add(a, b int) int { return a+b }
// 	func AddToA(a int) func(int) int {
// 		return func(b int) int {
//			return Add(a, b)
// 		}
// 	}
// 	addTo1 := AddToA(1)
// 	fmt.Println(Add(1,2) == addTo1(2)) // true

// Method values
//
//	A selected method may be passed similar to a closure;
//	the receiver is closed over at that point
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

// Interfaces in practice
// 	1. Let consumers define interfaces
// 	(what minimal behavior do they require?)
// 	2. Re-use standard interfaces wherever possible
// 	(what minimal behavior do they require?)
// 	3. Keep interface declarations small
// 	(The bigger the interface, the weaker the abstraction)
//  4. Compose one-method interfaces into larger interfaces, if needed
// 	5. Avoid coupling interfaces to particular types/ implementations
// 	6. Accept interfaces, but return concrete types
// 	(let the consumer of the return type decide how to use it)

// Interfaces vs concrete values
//
//	"Be literal in what you accept, be conservative in what you return"
//	Put the least restriction on what parameters you accept
//	(the minimal interface)
//	Don't require ReadWriteCloser if you only need to read
//	Avoid restricting the use of your return type (the concrete value you
//	return might fit with many interfaces!)
//	Returning *os.File is less restrictive than returning io.ReadWriteCloser
//	because files have other useful methods

//	Returning error is a good example of an exception to this rule

// Empty interfaces
// 	The interface{} type has no methods
// 	So it is satisfied by anything!
// 	Empty interfaces are commonly used; they're how the formatted I/O
// 	routines can print any type
// 	func fmt.Printf(f string, args ...interface{})
// 	Reflection is needed to determine what the concrete type is

func main() {
	var err error = XYZ(1)

	if err != nil {
		fmt.Println("oops")
	} else {
		fmt.Println("OK!")
	}

	p := Point{1, 2}
	q := Point{4, 6}
	distanceFromP := p.Distance        // this is a method value
	fmt.Println(distanceFromP(q))      // 5
	fmt.Printf("%T\n", Point.Distance) // func(main.Point, main.Point)
	p = Point{2, 2}
	fmt.Println(distanceFromP(q)) // 5 stays the same unless you put ptr
	//	func (p *Point) Distance(q Point) float64 { . . . }
}
