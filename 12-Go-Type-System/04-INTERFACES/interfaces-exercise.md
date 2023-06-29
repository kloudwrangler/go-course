Exercise: Implementing Interfaces and Polymorphism

Objective: The exercise aims to reinforce the concepts of implementing interfaces and achieving polymorphism through type embedding.

Instructions:
1. Create an interface named `Shape` with two methods: `Area() float64` and `Perimeter() float64`.
2. Implement the `Shape` interface for two different shapes: `Rectangle` and `Circle`.
   - The `Rectangle` struct should have width and height fields and provide implementations for `Area()` and `Perimeter()`.
   - The `Circle` struct should have a radius field and provide implementations for `Area()` and `Perimeter()`.
3. Create a function named `PrintShapeDetails` that accepts a parameter of type `Shape` and prints its area and perimeter.
4. In the `main` function, create instances of `Rectangle` and `Circle`.
5. Call the `PrintShapeDetails` function with the instances of `Rectangle` and `Circle`.
   - Verify that the area and perimeter of each shape are correctly printed.

Solution:

```go
package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func PrintShapeDetails(shape Shape) {
	fmt.Printf("Area: %.2f\n", shape.Area())
	fmt.Printf("Perimeter: %.2f\n", shape.Perimeter())
}

func main() {
	rect := Rectangle{width: 5, height: 3}
	circle := Circle{radius: 2.5}

	PrintShapeDetails(rect)
	PrintShapeDetails(circle)
}
```

Explanation:
- The code defines an interface `Shape` with two methods, `Area()` and `Perimeter()`.
- Two structs, `Rectangle` and `Circle`, are created to implement the `Shape` interface.
- The `PrintShapeDetails` function accepts a parameter of type `Shape` and uses it to call the `Area()` and `Perimeter()` methods, printing the results.
- In the `main` function, instances of `Rectangle` and `Circle` are created.
- The `PrintShapeDetails` function is called with these instances to verify that the area and perimeter are correctly calculated and printed.

This exercise allows students to practice implementing interfaces, defining methods, and achieving polymorphism through type embedding. It also reinforces the understanding of how different shapes can be treated interchangeably when they implement the same interface.