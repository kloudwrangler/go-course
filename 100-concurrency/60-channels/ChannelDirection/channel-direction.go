package main

import (
	"fmt"
)

func ping(pings chan<- string, msg string) {
	pings <- msg
	fmt.Println("Ping")

}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
	fmt.Println("Pong")
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	go ping(pings, "passed message")
	go pong(pings, pongs)
	fmt.Println(<-pongs)
}
