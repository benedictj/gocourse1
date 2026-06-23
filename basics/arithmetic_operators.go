package basics

import (
	"fmt"
	"math"
)

func main1() {
	// Variable declarations
	var a, b int = 10, 3
	var result int

	// Addition
	result = a + b
	fmt.Println("Addition:", result) // Output: 13

	// Subtraction
	result = a - b
	fmt.Println("Subtraction:", result) // Output: 7

	// Multiplication
	result = a * b
	fmt.Println("Multiplication:", result) // Output: 30

	// Division
	result = a / b
	fmt.Println("Division:", result) // Output: 3

	// Modulus
	result = a % b
	fmt.Println("Modulus:", result) // Output: 1

	const p float64 = 22 / 7.0
	fmt.Println("Value of p:", p) // Output: 3.0

	// Overflow with signed integers
	var maxInt int64 = 9223372036854775807 // This is the maximum value for a 64-bit(int64) signed integer
	fmt.Println(maxInt)

	maxInt++
	fmt.Println(maxInt) // Output: -9223372036854775808 (overflow)

	// Underflow with signed integers
	var uMaxInt uint64 = 18446744073709551615 // This is the maximum value for a 64-bit unsigned(uint) integer
	fmt.Println(uMaxInt)

	uMaxInt++
	fmt.Println(uMaxInt) // Output: 0 (underflow)

	// Underflow with floating-point numbers
	var smallFloat float64 = 1.0e-323
	fmt.Println(smallFloat) // Output: 1e-323

	smallFloat = smallFloat / math.MaxFloat64
	fmt.Println(smallFloat) // Output: 0 (underflow)

}
