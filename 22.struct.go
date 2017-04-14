package main

import "fmt"

func main() {
	rect1 := Rectangle{leftX: 1, topY: 50, width: 100, height: 200}
	rect2 := Rectangle{10, 10, 100, 100}

	fmt.Println(rect1)
	fmt.Println(rect1.leftX)
	fmt.Println(rect1.area())
	fmt.Println(rect2)
}

// New Data Type
type Rectangle struct {
	leftX  float64
	topY   float64
	height float64
	width  float64
}

// We attach the function to the type
// Also known as assigning a method to structs
func (rect *Rectangle) area() float64 {
	return rect.width * rect.height
}
