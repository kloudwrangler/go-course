package main

import (
	"fmt"
	"log"
	"math"

	// "net/http"
	// _ "net/http/pprof"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

func main() {
	// Create cpu profiling
	// Perform intensive calculations
	cpuProfiling()
	memoryProfiling()
	// Uncomment if you want to serve the profiling through an http server
	// http.ListenAndServe("localhost:8080", nil)
}
func memoryProfiling() {
	file, err := os.OpenFile("mem.prof", os.O_CREATE|os.O_RDWR, 0655)
	if err != nil {
		fmt.Println("Failed to create memory profiling file:", err)
		return
	}
	defer file.Close()

	// Allocate memory profiles
	runtime.GC()
	intenseMemoryUsage()
	if err := pprof.WriteHeapProfile(file); err != nil {
		fmt.Println("Failed to write memory profile:", err)
		return
	}
}

func cpuProfiling() {
	fmt.Println("Starting CPU-intensive task...")
	start := time.Now()

	file, err := os.OpenFile("cpu.prof", os.O_CREATE|os.O_RDWR, 0655)
	if err != nil {
		log.Fatal("open file error", err)
	}
	defer file.Close()
	pprof.StartCPUProfile(file)
	defer pprof.StopCPUProfile()

	intenseCpuUsage()

	elapsed := time.Since(start)
	fmt.Printf("Finished CPU-intensive task. Elapsed time: %s\n", elapsed)
}

func intenseCpuUsage() {
	for i := 0; i < 1000000000; i++ {
		_ = math.Sqrt(float64(i))
	}
}

func intenseMemoryUsage() {
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
