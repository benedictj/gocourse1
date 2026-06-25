package basics

import "fmt"

func loopbasics() {
	//
	// Simple iteration over a range of numbers using a for loop
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Iterate over collection (slice) using a for loop
	numbers := []int{1, 2, 3, 4, 5, 6}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue // Skip even numbers
		}
		fmt.Println("Odd number:", i)
		if i == 5 {
			break // Exit the loop when i equals 5
		}
	}

	rows := 5

	// Outer loop for rows
	for i := 1; i <= rows; i++ {
		// Inner loop for spaces before the stars
		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}
		// Inner loop for stars
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println() // Move to the next line after each row
	}

	for i := range 10 {
		fmt.Println(10 - i)
	}
	fmt.Println("We have lift off!")
}
