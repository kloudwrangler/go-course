// This sample program demonstrates how to create race
// conditions in our programs. We don't want to do this.
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	// counter is a variable incremented by all goroutines.
	counter int64

	// wg is used to wait for the program to finish.
	wg sync.WaitGroup
)

// main is the entry point for all Go programs.
func main() {
	// Add a count of two, one for each goroutine.
	wg.Add(2)

	// Create two goroutines.
	go incCounterByTwo(1)
	go incCounterByTwo(2)

	// Wait for the goroutines to finish.
	wg.Wait()
	fmt.Println("Final Counter:", counter)
	if counter != 4 {
		fmt.Printf("Error: Was expecting 4 and got %d", counter)
	} else {
		fmt.Println("Success: Got the result I was expecting")
	}

}

// incCounter increments the package level counter variable.
func incCounterByTwo(id int) {
	// Schedule the call to Done to tell main we are done.
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// Safely Add One To Counter.
		atomic.AddInt64(&counter, 1)
		// Yield the thread and be placed back in queue.
		runtime.Gosched()
	}
}
