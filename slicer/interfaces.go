package slicer

import (
	"fmt"
	"math"
)

type Circle struct {
	width float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Shape interface {
	area() float64
}

func (c Circle) area() float64 {
	return c.width * 2 * math.Pi
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func HelloInterfaces() {
	c1 := Circle{width: 12}
	r1 := Rectangle{
		width:  12,
		height: 6,
	}

	shapes := []Shape{c1, r1}

	for _, v := range shapes {
		fmt.Println("The area of shape is ", v.area())
	}

	fmt.Println(shapes)
}
