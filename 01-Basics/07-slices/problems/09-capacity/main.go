package main

import "fmt"

func main() {
	numbers := make([]int, 2, 5)
	fmt.Println("numbers:", numbers)
	fmt.Println("Len:", len(numbers))
	fmt.Println("Cap:", cap(numbers))

	for i := range 20 {
		fmt.Println("============================")
		numbers = append(numbers, (i+1)*10)
		fmt.Println("numbers:", numbers)
		fmt.Println("Append:", (i+1)*10)
		fmt.Println("Len:", len(numbers))
		fmt.Println("Cap:", cap(numbers))
	}

}
