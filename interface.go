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
	return math.Pi * c.radius * c.radius
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func main() {
	c := Circle{radius: 5}
	fmt.Println("Area of circle:", c.Area())

	r := Rectangle{width: 10, height: 5}
	fmt.Println("Area of rectangle:", r.Area())
}
