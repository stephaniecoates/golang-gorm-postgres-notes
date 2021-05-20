package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Square struct {
	width float64
}

type RightTriangle struct {
	aLength float64
	bLength float64
}

func (r *Rectangle) area() float64 {
	return r.width * r.height
}

func (s *Square) area() float64 {
	return math.Pow(2, s.width)
}

func (t *RightTriangle) area() float64 {
	// obviously this isn't how to get triangle area, fudged for example purposes
	return t.aLength * t.bLength
}

func (r *Rectangle) perimeter() float64 {
	return (2 * r.width) + (2 * r.height)
}

func (s *Square) perimeter() float64 {
	return s.width * 4
}

func (t *RightTriangle) perimeter() float64 {
	cLength := t.getCLength()
	return t.aLength + t.bLength + cLength
}

func (t *RightTriangle) getCLength() float64 {
	cSqrd := math.Pow(t.aLength, 2) + math.Pow(t.bLength, 2)

	return math.Sqrt(cSqrd)
}

// can pass interfaces in as function parameters too! and access their methods there.
func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	sqr := Square{4}
	rect := Rectangle{2, 5}
	tri := RightTriangle{3, 3}

	fmt.Println(sqr, rect, tri)

	// because interface methods now accept pointers, must change these args to pointers, rather than values themselves
	// this would be necessary if we were modifying data on the objs, rather than just reading it
	// that said, it's always best practice to pass the pointer so we have access to the original value if we need it.
	// never hurts to pass the pointer
	shapes := []Shape{&sqr, &rect, &tri}
	fmt.Println(shapes)

	for _, shape := range shapes {
		fmt.Println(getArea(shape))
	}
}
