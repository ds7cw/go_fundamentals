package main

import (
	"fmt"
	"image/color"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// An interface specifies abstract behaviour in terms of methods
// type Stringer interface { // in "fmt"
// 	String() string
// }
// Concrete types offer methods that satisfy the interface

// A method is a special type of function:
//
//	it has a receiver param before the func name param
type IntSlice []int

func (is IntSlice) String() string {
	var strs []string

	for _, v := range is {
		strs = append(strs, strconv.Itoa(v))
	}
	return "[" + strings.Join(strs, ";") + "]"
}

// Without interfaces, we'd have to write (many) funcs
// 	for (many) concrete types, possibly coupled to them
// 		func OutputToFile(f *File, . . .) {. . .}
// 		func OutputToBuffer(b *Buffer, . . .) {. . .}
// 		func OutputToSocket(s *Socket, . . .) {. . .}
// 	Better - we want to define our func in terms of
// 		abstract behaviour
// type Writer interface {
// 	Write([]byte) (int, error)
// }
// func OutputTo(w io.Writer, . . .) {. . .}

// An interface specifies required behaviour as a method set
// Any type that implements that method set satisfies the
// 	interface. This is known as structural typing ("duc" typing)

// A method may be defined on any user-declared (named) type*
// type MyInt int
// func (i MyInt) String() string {. . .}
// The same method name may be bound to different types
// *Some rules and restrictions apply, see package insert

// A method may take a pointer or value receiver, but not both
type Point struct {
	X, Y float64
}

func (p Point) Offset(X, Y float64) Point {
	return Point{p.X + X, p.Y + Y}
}

func (p *Point) Move(X, Y float64) {
	p.X += X
	p.Y += Y
}

// Taking a pointer allows the method to change the receiver
// 	(original object)

func exercise1() {
	var c ByteCounter

	f1, _ := os.Open("a.txt")
	f2 := &c

	n, _ := io.Copy(f2, f1)
	fmt.Println("copied", n, "bytes") // copied 32 bytes
	fmt.Println(c)                    // 32
}

type ByteCounter int

func (b *ByteCounter) Write(p []byte) (int, error) {
	l := len(p)
	*b += ByteCounter(l)
	return l, nil
}

// All the methods must be present to satisfy the interface
// var w io.Writer
// var rwc io.ReadWriteCloser
// w = os.Stdout 			// OK: *os.File has Write method
// w = new(bytes.Buffer)	// OK: *bytes.Buffer has Write method
// w = time.Second			// ERROR: no Write method
// rwc = os.Stdout			// OK: *os.File has all 3 methods
// rwc = new(bytes.Buffer)	// ERROR: no Close method
// w = rwc					// OK: io.ReadWrite
// rwc = w					// ERROR: no Close method
// Which is why it pays to keep interfaces small

// The receiver must be of the right type (pointer or value)

// INTERFACE COMPOSITION
// io.ReadWriter is actually defined by Go as two interfaces
// type Reader interface {
// 	Read(p []byte) (n int, err error)
// }
// type Writer interface {
// 	Write(p []byte) (n int, err error)
// }
// type ReadWriter interface {
// 	Reader
// 	Writer
// }
// Small interfaces with composition where needed are more flexible

// All methods for a given type must be declared in the same
// 	package where the type is declared
// 	This allows a package importing the type to know all the methods
// 	at compile time
// 	But we can always extend the type in a new package through embedding:

type Line struct {
	Begin, End Point
}

func (l Line) Distance() float64 {
	return math.Hypot(l.End.X-l.Begin.X, l.End.Y-l.Begin.Y)
}

type Path []Point

func (p Path) Distance() (sum float64) {
	for i := 1; i < len(p); i++ {
		sum += Line{p[i-1], p[i]}.Distance()
	}
	return sum
}

func exercise2() {
	side := Line{Point{1, 2}, Point{4, 6}}
	// fmt.Println(side.Distance()) // 5
	PrintDistance(side) // 5
}

func exercise3() {
	perimeter := Path{{1, 1}, {5, 1}, {5, 4}, {1, 1}}
	// fmt.Println(perimeter.Distance()) // 12
	PrintDistance(perimeter) // 12
}

type Distancer interface {
	Distance() float64
}

func PrintDistance(d Distancer) {
	fmt.Println(d.Distance())
}

func (l Line) ScaleBy(f float64) Line {
	l.End.X += (f - 1) * (l.End.X - l.Begin.X)
	l.End.Y += (f - 1) * (l.End.Y - l.Begin.Y)
	return Line{l.Begin, Point{l.End.X, l.End.Y}}
} // side := Line{Point{1,2}, Point{4,6}} // s2 := side.ScaleBy(2.5)

func exercise4() {
	fmt.Println(Line{Point{1, 2}, Point{4, 6}}.ScaleBy(2).Distance()) // 10
}

type ColoredPoint struct {
	Point
	Color color.RGBA
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func exercise5() {
	p := Point{1, 1}
	q := ColoredPoint{Point{4, 5}, color.RGBA{255, 0, 0, 255}}
	l1 := q.Distance(p)
	l2 := p.Distance(q.Point) // OK
	fmt.Println(l1, l2)       // 5 5
}

func main() {
	var v IntSlice = []int{1, 2, 3}
	var s fmt.Stringer = v

	for i, x := range v {
		fmt.Printf("%d: %d\n", i, x) // 0: 1 // 1: 2 // 2: 3
	}

	fmt.Printf("%T %[1]v\n", s) // main.IntSlice [1;2;3]
	// Uses String() method (below)
	fmt.Printf("%T %[1]v\n", v) // main.IntSlice [1;2;3]

	exercise1()
	exercise2() // 5
	exercise3() // 12
	exercise4() // 10
	exercise5() // 5 5
}
