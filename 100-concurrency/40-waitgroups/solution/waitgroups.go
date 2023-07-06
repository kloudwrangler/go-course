package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	fmt.Println("Before Numbers")
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			fmt.Printf("Running complex algorithm for %d\n", num)
			time.Sleep(time.Second * time.Duration(rand.Intn(3)+1))
			fmt.Printf("Done running complex algorithm for %d\n", num)
		}(i)
	}
	fmt.Println("After Numbers")
	wg.Wait()
	fmt.Println("Terminating Program")
}
