package main

import "fmt"

type Animal interface {
	Sound()
}

type Dog struct {
	Name string
}

func (d Dog) Sound() {
	fmt.Println("Woof")
}

type Cat struct {
	Name string
}

func (c Cat) Sound() {
	fmt.Println("Mew")
}

func MakeSound(a Animal) {
	a.Sound()
}

func main() {
	cat := Cat{
		Name: "Cat",
	}

	dog := Dog{
		Name: "Dog",
	}

	MakeSound(cat)
	MakeSound(dog)
}
