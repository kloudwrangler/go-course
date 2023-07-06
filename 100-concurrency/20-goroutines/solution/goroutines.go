package main

import (
	"fmt"
	"time"
)

func printNumber(num int) {
	fmt.Printf("%d in goroutine\n", num)
}

func main() {
	fmt.Println("Before Numbers")
	for i := 1; i <= 10; i++ {
		go printNumber(i)
	}
	fmt.Println("After Numbers")
	time.Sleep(2 * time.Second)
}
