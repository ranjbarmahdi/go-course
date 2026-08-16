package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func main() {
	rectangle := Rectangle{
		Width:  10,
		Height: 5,
	}

	fmt.Println("Area:", rectangle.Area())
	fmt.Println("Perimeter:", rectangle.Perimeter())
}
