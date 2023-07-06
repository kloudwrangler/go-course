package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func printArea(s Shape) {
	fmt.Printf("Area: %f\n", s.Area())
}

func main() {
	rect := Rectangle{width: 10, height: 5}
	circ := Circle{3}
	printArea(rect)
	printArea(circ)
}
