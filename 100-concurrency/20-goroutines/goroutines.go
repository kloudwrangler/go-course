package main

import (
	"fmt"
	"time"
)

func printNumbers(i int) {
	fmt.Println(i)
}

func main() {
	fmt.Println("Before Numbers")
	for i := 1; i <= 10; i++ {
		go printNumbers(i)
	}
	fmt.Println("Before Sleep")
	time.Sleep(3 * time.Second)
	fmt.Println("After Numbers")
}
