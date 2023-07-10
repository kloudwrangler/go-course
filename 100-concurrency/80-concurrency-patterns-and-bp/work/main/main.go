// This sample program demonstrates how to use the work package
// to use a pool of goroutines to get work done.
package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/aldorperez1/go-course/100-concurrency/80-concurrency-patterns-and-bp/work"
)

// names provides a set of names to display.
var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}

// namePrinter provides special support for printing names.
type namePrinter struct {
	name string
	iter int
}

// Task implements the Worker interface.
func (m *namePrinter) Task() {
	log.Println(m.name, m.iter)
	time.Sleep(time.Second)
}

// main is the entry point for all Go programs.
func main() {
	// Create a work pool with 2 goroutines.
	p := work.New(2)

	var wg sync.WaitGroup
	wg.Add(2 * len(names))

	for i := 0; i < 2; i++ {
		// Iterate over the slice of names.
		for _, name := range names {
			// Create a namePrinter and provide the
			// specific name.
			np := namePrinter{
				name: name,
				iter: i,
			}
			fmt.Printf("%+v\n", np)
			go func() {
				// Submit the task to be worked on. When RunTask
				// returns we know it is being handled.
				fmt.Printf("Before RUN - Iter: %d, Name: %s\n", np.iter, np.name)
				p.Run(&np)
				fmt.Printf("After RUN - Iter: %d, Name: %s\n", np.iter, np.name)
				wg.Done()
			}()
		}
	}
	fmt.Printf("About to Wait start waiting\n")
	wg.Wait()

	// Shutdown the work pool and wait for all existing work
	// to be completed.
	p.Shutdown()
}
