package main

import (
	"fmt"
)

type Speaker interface {
	Speak()
}

type Dog struct {
	Name string
}

func (d Dog) Speak() {
	fmt.Println("Woof")
}

type Cat struct {
	Name string
}

func (c Cat) Speak() {
	fmt.Println("Meow")
}

func makeSound(s Speaker) {
	s.Speak()
}

func print(value interface{}) {
	fmt.Println(value)
}

func check(value any) {

	switch v := value.(type) {

	case int:
		fmt.Println("Integer", v)

	case string:
		fmt.Println("String", v)

	case bool:
		fmt.Println("Boolean", v)
	case Dog:
		fmt.Println("Dog", v)
	default:
		fmt.Println("Unknown")

	}

}

func main() {
	fmt.Println("1. Basic")
	dog := Dog{
		Name: "Buddy",
	}

	cat := Cat{
		Name: "Kitty",
	}

	makeSound(dog)
	makeSound(cat)

	fmt.Println("2. Empty Interface")
	print("Hello World")
	print(2)
	print(true)

	fmt.Println("3. Type Assertion")
	var value any
	value = 100
	number, ok := value.(int)
	if ok {
		fmt.Println(number)
	}

	value = "hello"
	number2, ok2 := value.(int)
	if ok2 {
		fmt.Println(number2)
	}

	fmt.Println("4. Type Switch")
	check("123")
	check(1)
	check(true)

	dog2 := Dog{
		Name: "Black",
	}
	check(dog2)
}
