package main

import "fmt"

func main() {
	// Declaring an array using an array literal
	arr := [5]int{1, 2, 3, 4, 5}

	// Taking the slice of a slice
	slice := arr[2:4]

	// How length and capacity are calculated
	length := len(slice)
	capacity := cap(slice)

	fmt.Println("Original Array:", arr)
	fmt.Println("Original Slice:", slice)
	fmt.Println("Length:", length)
	fmt.Println("Capacity:", capacity)

	// Calculating the new length and capacity
	newSlice := slice[1:3]
	newLength := len(newSlice)
	newCapacity := cap(newSlice)

	fmt.Println("New Slice:", newSlice)
	fmt.Println("New Length:", newLength)
	fmt.Println("New Capacity:", newCapacity)

	// Potential consequence of making changes to a slice
	slice[0] = 10

	fmt.Println("Modified Slice:", slice)
	fmt.Println("Original Array (after modification):", arr)

	// Runtime error showing index out of range
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Runtime Error:", r)
		}
	}()
	element := slice[10]
	fmt.Println("Accessing index 10 of the slice:", element)
}
