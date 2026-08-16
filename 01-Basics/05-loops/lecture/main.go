package main

import "fmt"

func main() {
	fmt.Println("1. Classic for")
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("2. While loop")
	i := 0
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	fmt.Println("3. Infinite loop")
	j := 0
	for {
		fmt.Println(j)
		j++
		if j == 5 {
			break
		}
	}

	fmt.Println("4. Continue")
	for i := 0; i <= 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
}
