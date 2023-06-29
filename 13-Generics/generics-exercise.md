Exercise: Implementing a Generic Stack

In this exercise, you will implement a generic stack data structure using the concepts of generics in Go. The goal is to reinforce the understanding of generic types, type parameters, and type constraints.

Instructions:
1. Create a new Go file called `stack.go`.
2. Define a generic type `Stack[T]` using a type parameter `T`.
3. Apply a type constraint on `T` to ensure that it supports the equality comparison (`==`) by creating an interface called `Comparable` with a single method `Equals(other T) bool`.
4. Implement a function `NewStack[T]() *Stack[T]` that returns a new instance of the stack.
5. Implement the following methods for the `Stack[T]` type:
   - `Push(item T)` - Adds an item to the top of the stack.
   - `Pop() T` - Removes and returns the top item from the stack.
   - `IsEmpty() bool` - Returns `true` if the stack is empty, `false` otherwise.

Note: For the purpose of this exercise, assume that the stack has an unlimited capacity.

Solution:

```go
// stack.go

package main

import "fmt"

// Comparable interface for type constraint
type Comparable interface {
	Equals(other T) bool
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
func (s string) Equals(other string) bool {
	return s == other
}

// Implement the Equals method for type int
func (i int) Equals(other int) bool {
	return i == other
}

func main() {
	stack := NewStack[string]()
	stack.Push("Alice")
	stack.Push("Bob")
	stack.Push("Charlie")

	for !stack.IsEmpty() {
		item := stack.Pop()
		fmt.Println(item)
	}
}
```

Explanation:
- We define a generic type `Stack[T]` using a type parameter `T`. The type parameter `T` is constrained to implement the `Comparable` interface.
- The `Comparable` interface declares a single method `Equals(other T) bool`, which will be used to ensure the type supports equality comparison.
- The `NewStack[T]` function creates a new instance of the stack by initializing an empty slice of type `T`.
- The `Push` method adds an item of type `T` to the top of the stack by appending it to the slice.
- The `Pop` method removes and returns the top item from the stack. It uses the length of the slice to determine the index of the top item and then reslices the slice accordingly.
- The `IsEmpty` method returns `true` if the length of the slice is 0, indicating that the stack is empty.
- In the `main` function, we create a stack of strings using `NewStack[string]()`. We push three strings onto the stack

 and then pop and print them until the stack is empty.
- For demonstration purposes, we implement the `Equals` method for the `string` and `int` types. In a real scenario, these types would already support equality comparison.

By implementing this exercise, students will gain hands-on experience with generics in Go by creating a generic stack data structure. They will understand how to define generic types, apply type constraints, and use type parameters effectively.