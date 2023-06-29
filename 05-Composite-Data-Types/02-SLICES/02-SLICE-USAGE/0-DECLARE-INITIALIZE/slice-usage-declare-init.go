package main

import "fmt"

func main() {
	// Declaring a slice of strings by length
	var names []string
	names = make([]string, 3)
	names[0] = "Alice"
	names[1] = "Bob"
	names[2] = "Charlie"
	fmt.Println("Names:", names)

	// Declaring a slice of integers by length and capacity
	numbers := make([]int, 5, 10)
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	numbers[3] = 4
	numbers[4] = 5
	fmt.Println("Numbers:", numbers)

	// Compiler error setting capacity less than length
	// Uncomment the following line to see the compiler error
	// invalid slice index: 5 > 4 (out of bounds for 5-element array)
	// numbers[5] = 6

	// Declaring a slice with a slice literal
	fruits := []string{"apple", "banana", "orange"}
	fmt.Println("Fruits:", fruits)

	// Declaring a slice with index positions
	colors := []string{0: "red", 2: "blue", 1: "green"}
	fmt.Println("Colors:", colors)

	// Declaration differences between arrays and slices
	array := [3]int{1, 2, 3}
	arraySlice := array[:]
	fmt.Println("Array:", array)
	fmt.Println("Array Slice:", arraySlice)

	// Declaring a nil slice
	var data []int
	fmt.Println("Data (nil slice):", data)

	// Declaring an empty slice
	scores := []int{}
	fmt.Println("Scores (empty slice):", scores)
}
