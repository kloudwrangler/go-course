---
marp: true
theme: uncover
---
# Functions in Go: An Introduction and Best Practices

---
- Welcome to the Functions in Go lesson
- Target audience: Experienced developers with no prior knowledge of Go
- Objectives: Understand the basics of functions in Go and learn best practices
- Topics covered: Functions, Multiple Return Values, Named Return Values, Variadic Functions, Anonymous Functions & Closures, Recursion, First-class Functions, Higher Order Functions

---
- Functions are essential building blocks of Go programs
- Syntax: `func functionName(parameters) returnType { ... }`
- Example: `func add(a, b int) int { return a + b }`
- Functions can have parameters and a return type
- Encourage the use of meaningful names for functions and parameters

---
- Go supports returning multiple values from a function
- Example: `func calculate(a, b int) (int, int) { return a+b, a-b }`
- Demonstrate how to use and assign multiple return values

---
- Go allows naming return values in function signatures
- Named return values are initialized with zero values by default
- Example: `func calculate(a, b int) (sum int, diff int) { sum = a + b; diff = a - b; return }`
- Benefits of using named return values: improves code readability

---
- Variadic functions accept a variable number of arguments
- Syntax: `func functionName(parameters ...type) { ... }`
- Example: `func sum(numbers ...int) int { ... }`
- Demonstrate the usage of variadic functions with examples

---
- Anonymous functions are functions without names
- They can be assigned to variables or invoked directly
- Example: `add := func(a, b int) int { return a + b }`
- Closures: Anonymous functions that access variables from their surrounding scope
- Show an example of a closure using an anonymous function

---
- Recursion is the process of a function calling itself
- Base case: condition that stops the recursive calls
- Example: `func factorial(n int) int { if n == 0 { return 1 }; return n * factorial(n-1) }`
- Explain the concept of recursion and demonstrate its usage

---
- In Go, functions are first-class citizens
- Functions can be assigned to variables, passed as arguments, and returned as values
- Explain the benefits and use cases of first-class functions

---
- Higher-order functions take one or more functions as arguments or return a function
- Example: `func applyOperation(operation func(int) int, num int) int { return operation(num) }`
- Show examples of higher-order functions to illustrate their usage

---
- Use meaningful names for functions, parameters, and return values
- Keep functions small and focused on a single task (single responsibility principle)
- Avoid deep levels of nesting in functions
- Document functions using comments to clarify their purpose, parameters, and return values
- Write unit tests for functions to ensure their correctness and maintainability

---
- Recap the main topics covered: Functions, Multiple Return Values, Named Return Values, Variadic Functions, Anonymous Functions & Closures, Recursion, First-class Functions, Higher Order Functions
- Emphasize the importance of understanding and applying these concepts effectively
- Encourage further practice and exploration of Go functions

---
- Open the floor for any questions or discussion
Certainly! Here's the expanded code block with more examples for each topic:

```go
package main

import "fmt"

// Functions
func add(a, b int) int {
	return a + b
}

// Multiple Return Values
func calculate(a, b int) (int, int) {
	return a + b, a - b
}

// Named Return Values
func calculateNamed(a, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b
	return
}

// Variadic Functions
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// Anonymous Functions & Closures
func anonymousFunc() {
	add := func(a, b int) int {
		return a + b
	}

	result := add(2, 3)
	fmt.Println("Result of anonymous function:", result)

	// Closure example
	x := 5
	increment := func() int {
		x++
		return x
	}
	fmt.Println("Closure result:", increment())
	fmt.Println("Closure result:", increment())
}

// Recursion
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// First-class Functions
func multiply(a, b int) int {
	return a * b
}

// Higher Order Functions
func applyOperation(operation func(int, int) int, a, b int) int {
	return operation(a, b)
}

func main() {
	// Functions
	fmt.Println("Sum:", add(2, 3))

	// Multiple Return Values
	sumResult, diffResult := calculate(5, 3)
	fmt.Println("Sum:", sumResult)
	fmt.Println("Difference:", diffResult)

	// Named Return Values
	sumResultNamed, diffResultNamed := calculateNamed(5, 3)
	fmt.Println("Sum (named):", sumResultNamed)
	fmt.Println("Difference (named):", diffResultNamed)

	// Variadic Functions
	fmt.Println("Total:", sum(1, 2, 3, 4, 5))

	// Anonymous Functions & Closures
	anonymousFunc()

	// Recursion
	fmt.Println("Factorial of 5:", factorial(5))

	// First-class Functions
	fmt.Println("Multiplication:", multiply(2, 3))

	// Higher Order Functions
	result := applyOperation(multiply, 4, 5)
	fmt.Println("Result of applyOperation:", result)
}
```

This code block demonstrates the usage of functions, multiple return values, named return values, variadic functions, anonymous functions, closures, recursion, first-class functions, and higher-order functions. It includes examples and output statements to showcase the functionality of each concept. Feel free to modify and enhance the code as needed for your class.


---


```go
package main

import "fmt"

// Functions
func add(a, b int) int {
	return a + b
}

func main() {
	// Functions
	fmt.Println("Sum:", add(2, 3))
}
```

```go
package main

import "fmt"

// Multiple Return Values
func calculate(a, b int) (int, int) {
	return a + b, a - b
}

func main() {
	// Multiple Return Values
	sumResult, diffResult := calculate(5, 3)
	fmt.Println("Sum:", sumResult)
	fmt.Println("Difference:", diffResult)
}
```

```go
package main

import "fmt"

// Named Return Values
func calculateNamed(a, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b
	return
}

func main() {
	// Named Return Values
	sumResultNamed, diffResultNamed := calculateNamed(5, 3)
	fmt.Println("Sum (named):", sumResultNamed)
	fmt.Println("Difference (named):", diffResultNamed)
}
```

```go
package main

import "fmt"

// Variadic Functions
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	// Variadic Functions
	fmt.Println("Total:", sum(1, 2, 3, 4, 5))
}
```

```go
package main

import "fmt"

// Anonymous Functions & Closures
func anonymousFunc() {
	add := func(a, b int) int {
		return a + b
	}

	result := add(2, 3)
	fmt.Println("Result of anonymous function:", result)

	// Closure example
	x := 5
	increment := func() int {
		x++
		return x
	}
	fmt.Println("Closure result:", increment())
	fmt.Println("Closure result:", increment())
}

func main() {
	// Anonymous Functions & Closures
	anonymousFunc()
}
```

```go
package main

import "fmt"

// Recursion
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	// Recursion
	fmt.Println("Factorial of 5:", factorial(5))
}
```

```go
package main

import "fmt"

// First-class Functions
func multiply(a, b int) int {
	return a * b
}

func main() {
	// First-class Functions
	fmt.Println("Multiplication:", multiply(2, 3))
}
```

```go
package main

import "fmt"

// Higher Order Functions
func applyOperation(operation func(int, int) int, a, b int) int {
	return operation(a, b)
}

func main() {
	// Higher Order Functions
	result := applyOperation(multiply, 4, 5)
	fmt.Println("Result of applyOperation:", result)
}
```

You can use these individual code blocks to explain each concept separately in your slides.