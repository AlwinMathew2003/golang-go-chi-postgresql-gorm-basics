package main

import (
	"fmt"
)

// Square struct
type Square struct {
	side float64
}

// Pointer receiver modifies the original object
func (s *Square) DoubleSide() {
	s.side *= 2
}

// Circle struct
type Circle struct {
	radius float64
}

// Value receiver cannot modify the original object
func (c Circle) DoubleRadius() {
	c.radius *= 2
}

func main() {
	shape := &Square{side: 5.0}
	shape.DoubleSide()             // Modifies the original object
	fmt.Println("Square side:", shape.side) // Prints 10.0

	shape2 := Circle{radius: 3.0}
	shape2.DoubleRadius()            // Operates on a copy
	fmt.Println("Circle radius:", shape2.radius) // Still prints 3.0

}
