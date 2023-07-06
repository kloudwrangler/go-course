package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func complexAlgorithm(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Running complex algorithm for %d\n", i)
	time.Sleep(time.Second * time.Duration(rand.Intn(3)+1))
	fmt.Printf("Done running complex algorithm for %d\n", i)
}
func main() {
	var wg sync.WaitGroup
	fmt.Println("Before Numbers")
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go complexAlgorithm(i, &wg)
	}
	fmt.Println("After Numbers")
	wg.Wait()
	fmt.Println("Terminating Program")
}
