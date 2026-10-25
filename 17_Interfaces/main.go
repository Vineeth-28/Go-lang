package main

import (
	"fmt"
	"math"
)

// Definition of the interface
type Shape interface {
	Area() float64
}

// Concrete type 1: Circle
type Circle struct {
	Radius float64
}

// Circle implements Shape implicitly by having an Area() method
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Concrete type 2: Rectangle
type Rectangle struct {
	Width, Height float64
}

// Rectangle also implements Shape implicitly
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// A generic function that accepts the interface
func printArea(s Shape) {
	fmt.Printf("Area: %f\n", s.Area())
}

func main() {
	c := Circle{Radius: 5}
	r := Rectangle{Width: 10, Height: 5}

	printArea(c) // Works!
	printArea(r) // Works!
}
