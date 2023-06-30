package main

import "fmt"

func main() {
	// Exercise 1: Iterating over a slice of integers
	intSlice := []int{10, 20, 30, 40, 50}
	fmt.Println("Exercise 1: Iterating over a slice of integers")
	for _, value := range intSlice {
		fmt.Println(value)
	}

	fmt.Println("--------------------")

	// Exercise 2: Iterating over a slice of strings with indices
	strSlice := []string{"apple", "banana", "cherry", "date"}
	fmt.Println("Exercise 2: Iterating over a slice of strings with indices")
	for index, value := range strSlice {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}

	fmt.Println("--------------------")

	// Exercise 3: Iterating over a slice of floats and calculating the sum
	floatSlice := []float64{3.14, 1.618, 2.718, 0.577}
	fmt.Println("Exercise 3: Iterating over a slice of floats and calculating the sum")
	sum := 0.0
	for i := 0; i < len(floatSlice); i++ {
		sum += floatSlice[i]
	}
	fmt.Println("Sum:", sum)
}
