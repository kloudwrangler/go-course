package main

import "fmt"

func multiply(a, b int) int {
	return a * b
}
func sum(a, b int) int {
	return a + b
}
func applyOperation(operation func(int, int) int, a int, b int) int {
	return operation(a, b)
}
func main() {
	resultMult := applyOperation(multiply, 4, 5)
	resultSum := applyOperation(sum, 4, 5)
	fmt.Println("Result of applyOperation multiply:", resultMult)
	fmt.Println("Result of applyOperation sum:", resultSum)
}
