package main

import (
	"fmt"
	"time"
)

func worker1(c chan<- string) {
	time.Sleep(2 * time.Second)
	c <- "Worker 1: Task completed"
}

func worker2(c chan<- string) {
	time.Sleep(4 * time.Second)
	c <- "Worker 2: Task completed"
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go worker1(ch1)
	go worker2(ch2)

	select {
	case result := <-ch1:
		fmt.Println(result)
	case result := <-ch2:
		fmt.Println(result)
	}

	select {
	case ch1 <- "Message from main to Worker 1":
		fmt.Println("Message sent to Worker 1")
	default:
		fmt.Println("Worker 1 is busy")
	}

	select {
	case <-ch1:
		fmt.Println("Received data from Worker 1")
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout! Worker 1 took too long to respond")
	}
}
