package main

import "fmt"

func increment(ptr *int) {
	// Increment the value pointed to by ptr by 1
	*ptr = *ptr + 1
}

func main() {
	// Task 1: Declare and Initialize Pointers
	var ptr *int
	num := 42
	ptr = &num

	// Task 2: Dereference and Modify Pointers
	fmt.Println("Value of num:", *ptr)
	*ptr = 99

	// Task 3: Pass Pointers to Functions
	increment(ptr)
	fmt.Println("Value of num after increment:", num)

	// Task 4: Reflection
	// We pass the address of num to the increment function because we want to modify the actual value of num,
	// and not a copy of it. By passing the address (pointer), the function can directly modify the value at that memory location.
}
