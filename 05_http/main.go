//Write a program that creates two custom struct types called 'triangle' and 'square'
// The 'square' type shjould be a struct with a field called 'sideLength' of type float64
// The 'triangle' type should be a struct with a field called 'height' of type float64 and a field 'base' of type float64
// Both types should have function called 'getArea' that returns the calculated area of the square or triangle
// Add a 'shape' interface that defines a function called 'printArea'.

package main

import "fmt"

type shape interface {
	getArea() float64
	printArea()
}
type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

func main() {
	s := square{
		sideLength: 4,
	}

	t := triangle{
		base:   2,
		height: 3,
	}

	s.printArea()
	fmt.Println()
	t.printArea()
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (s square) printArea() {
	fmt.Printf("Square area is: %v", s.getArea())
}

func (t triangle) getArea() float64 {
	return t.base * t.height * 0.5
}

func (t triangle) printArea() {
	fmt.Printf("Triangle area is: %v", t.getArea())
}
