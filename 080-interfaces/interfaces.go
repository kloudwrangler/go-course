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
