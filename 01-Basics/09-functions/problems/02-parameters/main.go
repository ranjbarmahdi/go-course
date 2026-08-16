package main

import "fmt"

func printPerson(name string, age int) {
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
}

func main() {
	printPerson("Mahdi", 28)
}
