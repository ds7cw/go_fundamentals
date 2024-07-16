package main

import "fmt"

func main() {
	// INTEGERS
	var num1 int = 20
	num2 := 30
	var num3 int
	fmt.Println("Multiply:")
	num3 = 7
	fmt.Println(num1, num2, num3)
	fmt.Println("Result:")
	fmt.Println(num1 * num2 * num3)

	// uint8  - Unsigned 8-bit integers 0 to 255
	// uint16 - Unsigned 16-bit integers 0 to 65535
	// uint32 - Unsigned 32-bit integers 0 to 4294967295
	// uint64 - Unsigned 64-bit integers 0 to 18446744073709551615
	// int8   - Signed 8-bit integers −128 to 127
	// int16  - Signed 16-bit integers −32768 to 32767
	// int32  - Signed 32-bit integers −2147483648 to 2147483647
	// int64  - Signed 64-bit integers −9223372036854775808 to 9223372036854775807

	// BYTES
	// byte - same as uint8
	// rune - same as int32
	// uint - 32 or 64 bits
	// int - same size as uint
	// uintptr - an unsigned integer to store the uninterpreted bits of a pointer value

	// FLOATS
	// float32 - IEEE-754 32-bit floating-point numbers
	// float64 - IEEE-754 64-bit floating-point numbers
	// complex64 - Complex numbers with float32 real and imaginary parts
	// complex128 - Complex numbers with float64 real and imaginary parts

	// EXAMPLES
	var intbit8 int8 = -128
	var intbit16 int16 = -3200
	var uintbit8 uint8 = 255
	fmt.Println("\nint8:", intbit8, "\nint16:", intbit16, "\nuint8:", uintbit8)

}
