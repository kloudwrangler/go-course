package main

import "fmt"

func doubleSlice(slice []int) {
	for i := range slice {
		slice[i] *= 2
	}
}

func modifySlice(slice []int) []int {
	slice = append(slice, 5)
	return slice
}

func main() {
	// Create a slice of integers
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)
	// Call the doubleSlice function
	doubleSlice(numbers)
	fmt.Println(numbers)
	// Call the modifySlice function
	numbers = modifySlice(numbers)
	fmt.Println(numbers)
}
