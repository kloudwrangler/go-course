Exercise: Shape Operations

Objective: Practice implementing methods with value and pointer receivers in Go.

Instructions:
1. Create a Go program that models different shapes: `Circle` and `Rectangle`.
2. Each shape should have the following methods:
   - `Area()` that calculates and returns the area of the shape.
   - `Resize(factor float64)` that resizes the shape by a given factor. For `Circle`, the factor should be applied to the radius, and for `Rectangle`, it should be applied to both the width and height.
3. Create instances of `Circle` and `Rectangle` and test the methods on them.
4. Experiment with calling the methods using both value and pointer receivers.
5. Observe the behavior and results of the methods based on the receiver type.

Solution:

```go
package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

type Rectangle struct {
	width, height float64
}

// Method with value receiver for Circle
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// Method with pointer receiver for Circle
func (c *Circle) Resize(factor float64) {
	c.radius *= factor
}

// Method with value receiver for Rectangle
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Method with pointer receiver for Rectangle
func (r *Rectangle) Resize(factor float64) {
	r.width *= factor
	r.height *= factor
}

func main() {
	// Creating a circle and rectangle instances
	circle := Circle{radius: 5}
	rectangle := Rectangle{width: 4, height: 3}

	// Calling methods on circle using value receiver
	circleArea := circle.Area()
	fmt.Println("Circle area:", circleArea)

	// Calling methods on rectangle using value receiver
	rectangleArea := rectangle.Area()
	fmt.Println("Rectangle area:", rectangleArea)

	// Calling methods on circle using pointer receiver
	circle.Resize(2)
	circleAreaAfterResize := circle.Area()
	fmt.Println("Circle area after resize:", circleAreaAfterResize)

	// Calling methods on rectangle using pointer receiver
	rectangle.Resize(1.5)
	rectangleAreaAfterResize := rectangle.Area()
	fmt.Println("Rectangle area after resize:", rectangleAreaAfterResize)
}
```

Explanation:
- In this exercise, we define two shapes: `Circle` and `Rectangle`.
- Both shapes have an `Area()` method, which calculates and returns their respective areas.
- The `Resize(factor float64)` method is implemented for both shapes to resize them by a given factor.
- We create instances of `Circle` and `Rectangle` and demonstrate the usage of methods on them.
- First, we call the `Area()` method on both shapes using value receivers.
- Then, we call the `Resize(factor)` method on both shapes using pointer receivers.
- Finally, we call the `Area()` method again on both shapes to observe the updated areas after resizing.

Expected Output:
```
Circle area: 78.53981633974483
Rectangle area: 12
Circle area after resize: 314.1592653589793
Rectangle area after resize: 18
```

Note: Encourage participants to experiment with different values and try calling the methods with both value and pointer receivers to observe the behavior and understand the distinction between them.