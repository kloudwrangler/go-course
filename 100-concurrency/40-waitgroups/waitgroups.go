package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Before Numbers")
	for i := 1; i <= 10; i++ {
		go func(num int) {
			fmt.Printf("Running complex algorithm for %d\n", num)
			time.Sleep(time.Second * time.Duration(rand.Intn(3)+1))
			fmt.Printf("Done running complex algorithm for %d\n", num)
		}(i)
	}
	fmt.Println("After Numbers")
	fmt.Println("Terminating Program")
}
