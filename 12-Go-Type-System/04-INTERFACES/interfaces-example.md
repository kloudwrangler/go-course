Certainly! Here's an example program that covers the topics of interfaces, type embedding, type assertions, and type switches:

```go
package main

import (
	"fmt"
)

// Interface definition
type Shape interface {
	Area() float64
}

type Perimeter interface {
	Perimeter() float64
}

// Type embedding
type Rectangle struct {
	width, height float64
}

// Method implementations
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func main() {
	// Creating an instance of Rectangle
	rect := Rectangle{width: 10, height: 5}

	// Demonstrating interface usage
	var s Shape
	s = rect
	fmt.Printf("Area: %.2f\n", s.Area()) // Output: Area: 50.00

	var p Perimeter
	p = rect
	fmt.Printf("Perimeter: %.2f\n", p.Perimeter()) // Output: Perimeter: 30.00

	// Demonstrating type assertion
	rect2, ok := s.(Rectangle)
	if ok {
		fmt.Printf("Type assertion: %v\n", rect2) // Output: Type assertion: {10 5}
	}

	// Demonstrating type switch
	switch shape := s.(type) {
	case Rectangle:
		fmt.Printf("Type switch: Rectangle with area %.2f\n", shape.Area())
	default:
		fmt.Println("Type switch: Unknown shape")
	}
}
```

This program showcases the following topics:

1. Interfaces: The `Shape` and `Perimeter` interfaces define the behavior expected from implementing types.

2. Type Embedding: The `Rectangle` struct embeds the `Shape` and `Perimeter` interfaces, extending its behavior to fulfill those interfaces.

3. Method Implementations: The `Rectangle` struct provides method implementations for both the `Area()` and `Perimeter()` methods.

4. Interface Usage: The `rect` instance of `Rectangle` is assigned to variables of type `Shape` and `Perimeter`, demonstrating how a single type can be treated as multiple interfaces.

5. Type Assertion: The `s` interface is type asserted to the concrete type `Rectangle`, and the resulting value is printed if the assertion is successful.

6. Type Switch: The `s` interface is type switched to handle different types, and in this case, the `Rectangle` case is executed, showcasing the use of type switches.

You can run this program to see the output that illustrates the concepts covered in the lesson. Feel free to modify or expand the program as needed for your teaching purposes.