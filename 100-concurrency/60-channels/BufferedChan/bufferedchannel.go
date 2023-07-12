package main

import "fmt"

func main() {
	// Channel buffering and blocking
	bufferedCh := make(chan string, 2) // Buffered channel with capacity 2

	bufferedCh <- "Hello" // Sending data without blocking
	bufferedCh <- "World"

	// Uncomment the line below to demonstrate blocking behavior
	// bufferedCh <- "Blocked"

	fmt.Println("Buffered channel values:")
	fmt.Println(<-bufferedCh) // Receiving buffered values
	fmt.Println(<-bufferedCh)

	// Uncomment the line below to demonstrate blocking behavior
	fmt.Println(<-bufferedCh)
}
