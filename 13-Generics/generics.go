package main

import "fmt"

// Comparable interface for type constraint
type Comparable interface {
	Equals(other interface{}) bool
}

// Stack represents a generic stack
type Stack[T Comparable] []T

// NewStack creates a new instance of Stack[T]
func NewStack[T Comparable]() *Stack[T] {
	stack := make(Stack[T], 0)
	return &stack
}

// Push adds an item to the top of the stack
func (s *Stack[T]) Push(item T) {
	*s = append(*s, item)
}

// Pop removes and returns the top item from the stack
func (s *Stack[T]) Pop() T {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	index := len(*s) - 1
	item := (*s)[index]
	*s = (*s)[:index]
	return item
}

// IsEmpty returns true if the stack is empty, false otherwise
func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}

// Implement the Equals method for type string
func (s string) Equals(other interface{}) bool {
	if val, ok := other.(string); ok {
		return s == val
	}
	return false
}

// Implement the Equals method for type int
func (i int) Equals(other interface{}) bool {
	if val, ok := other.(int); ok {
		return i == val
	}
	return false
}

func main() {
	stack := NewStack[string]()
	stack.Push("Alice")
	stack.Push("Bob")
	stack.Push("Charlie")

	fmt.Println("Stack elements:")
	for !stack.IsEmpty() {
		item := stack.Pop()
		fmt.Println(item)
	}

	numStack := NewStack[int]()
	numStack.Push(10)
	numStack.Push(20)
	numStack.Push(30)

	fmt.Println("Numeric Stack elements:")
	for !numStack.IsEmpty() {
		item := numStack.Pop()
		fmt.Println(item)
	}
}
