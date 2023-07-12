package main

import (
	"fmt"
	"math"

	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	// Create cpu profiling
	// Perform intensive calculations
	go intenseMemoryUsage()
	go intenseCpuUsage()
	fmt.Printf("Starting http server at http://localhost:8080\nFor profiling info http://localhost:8080/debug/pprof/\n")	
	http.ListenAndServe("localhost:8080", nil)
}


func intenseCpuUsage() {
	fmt.Println("Starting IntenseCpuUsage")
	for i := 0; i < 1000000000; i++ {
		_ = math.Sqrt(float64(i))
	}
}

func intenseMemoryUsage() {
	fmt.Println("Starting IntenseMemoryUsage")
	var data [][]byte

	for i := 0; i < 10; i++ {
		// Allocate a large block of memory
		block := make([]byte, 1024*1024) // 1 MB

		// Append the block to the data slice
		data = append(data, block)

		// Introduce some variability to simulate memory usage
		time.Sleep(10 * time.Millisecond)
	}

	// Use the data slice to prevent optimizations
	for _, block := range data {
		_ = len(block)
	}
}
