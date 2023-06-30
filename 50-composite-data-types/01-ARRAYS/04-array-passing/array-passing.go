package main

import "fmt"

// Function to modify an array by value
func modifyArrayByValue(array [5]int) {
	array[0] = 10
}

// Function to modify an array by pointer
func modifyArrayByPointer(array *[5]int) {
	(*array)[0] = 10
}

func main() {
	// Declaring and initializing an array
	array := [5]int{1, 2, 3, 4, 5}

	fmt.Println("Original array:", array)

	// Passing the array by value to the modifyArrayByValue function
	modifyArrayByValue(array)
	fmt.Println("After modifyArrayByValue:", array)

	// Passing the array by pointer to the modifyArrayByPointer function
	modifyArrayByPointer(&array)
	fmt.Println("After modifyArrayByPointer:", array)
}
