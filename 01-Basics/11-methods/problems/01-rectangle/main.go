package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

func main() {
	rectangle := Rectangle{
		Width:  10,
		Height: 5,
	}
	fmt.Println("Area:", rectangle.Area())
	fmt.Println("Perimeter:", rectangle.Perimeter())
}
