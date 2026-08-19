package main

import "fmt"

func GenerateWelcome(name string) string {
	return fmt.Sprintf("Welcome %s", name)
}

func main() {
	name := "Mahdi"
	fmt.Println(GenerateWelcome(name))
}
