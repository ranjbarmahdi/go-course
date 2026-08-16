package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

func main() {
	rectangle := Rectangle{
		Width:  10,
		Height: 5,
	}

	area := rectangle.Height * rectangle.Width
	perimeter := 2 * (rectangle.Height + rectangle.Width)

	fmt.Println("Area:", area)
	fmt.Println("Perimeter:", perimeter)
}
