package main

import "fmt"

// 1. var
// 2. :=
// 3. const
// 4. string
// 5. int
// 6. float64
// 7. bool
// 8. type inference
// 9. type conversion
// 10. arithmetic

func main() {
	fmt.Println("1. var")
	var age int = 25
	fmt.Println(age)

	fmt.Println("2. Short declaration :=")
	name := "Mahdi"
	height := 1.80
	developer := true
	fmt.Println(name)
	fmt.Println(height)
	fmt.Println(developer)

	fmt.Println("3. const")
	const pi = 3.14159
	fmt.Println(pi)
}
