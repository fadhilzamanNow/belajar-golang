package main

import (
	"fmt"
	"math"
)

type geometry interface {
	findArea() float64
	findPerimeter() float64
}

type Rect struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rect) findArea() float64 {
	return r.width * r.height
}

func (r Rect) findPerimeter() float64 {
	return 2 * (r.height + r.width)
}

func (c Circle) findArea() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) findPerimeter() float64 {
	return 2 * math.Pi * c.radius
}

func detectGeometry(g geometry) {
	if c, ok := g.(Circle); ok {
		fmt.Println("The radius of circle is", c.radius)
	}

	if r, ok := g.(Rect); ok {
		fmt.Println("The width of rect is", r.width, "and the height is", r.height)
	}
}

func main() {
	r := Rect{width: 10, height: 50}
	fmt.Println("The area is", r.findArea(), "and the parameter is", r.findPerimeter())

	detectGeometry(r)
}
