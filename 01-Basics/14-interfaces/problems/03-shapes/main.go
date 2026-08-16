package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

func CalculateArea(shape Shape) float64 {
	return shape.Area()
}

func main() {
	rectangle := Rectangle{
		Width:  10,
		Height: 50,
	}

	circle := Circle{
		Radius: 10,
	}

	fmt.Println(CalculateArea(rectangle))
	fmt.Println(CalculateArea(circle))

}

var _ Shape = (*Circle)(nil)
var _ Shape = (*Rectangle)(nil)
