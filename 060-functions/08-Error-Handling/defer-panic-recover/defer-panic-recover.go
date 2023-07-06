package main

import (
	"fmt"
)

func main() {
	fmt.Println("Program Started")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	// Example 1: Defer
	fmt.Println("Example 1: Defer")
	defer fmt.Println("Deferred statement 1")
	defer fmt.Println("Deferred statement 2")

	// Example 2: Panic
	fmt.Println("Example 2: Panic")
	performPanic()

	// Example 3: Recover
	fmt.Println("Example 3: Recover")
	performRecover()

	fmt.Println("Program Completed")
}

func performPanic() {
	fmt.Println("Performing panic...")
	panic("Something went wrong!")
}

func performRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic in performRecover:", r)
		}
	}()

	fmt.Println("Performing panic in performRecover...")
	panic("Another panic occurred!")
}
