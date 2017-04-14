package main

import "fmt"
import "math"

func main() {
	rect := Rectangle{20, 50}
	circ := Circle{4}

	fmt.Println("Rectangle =", getArea(rect))
	fmt.Println("Circle =", getArea(circ))
}

type Shape interface {
	area() float64
}

type Rectangle struct {
	height float64
	width  float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

// This pulls the interface and  struct together
// This happens automatically
func getArea(shape Shape) float64 {
	return shape.area()
}
