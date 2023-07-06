package main

import (
	"errors"
	"fmt"
)

type NoNameProvided struct{}

func (e NoNameProvided) Error() string {
	return "No Name Provided"
}

func sayHello(s string) (string, error) {
	if s == "" {
		return "", &NoNameProvided{}
	}
	return fmt.Sprintf("Hello %s", s), nil
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {
	numerator := 10
	denominator := 0

	result, err := divide(numerator, denominator)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	} else {
		fmt.Printf("Result: %d\n", result)
	}

	res, err := sayHello("")
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	} else {
		fmt.Printf(res)
	}
}
