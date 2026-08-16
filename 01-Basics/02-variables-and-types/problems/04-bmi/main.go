package main

import "fmt"

func main() {
	weight := 80.0
	height := 1.80

	bmi := weight / (height * height)

	fmt.Println("Weight:", weight, "kg")
	fmt.Println("Height:", height, "m")
	fmt.Printf("BMI: %.2f", bmi)
}
