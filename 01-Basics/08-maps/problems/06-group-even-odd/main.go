package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}
	numbersGroup := map[string][]int{
		"even": []int{},
		"odd":  []int{},
	}

	for _, num := range numbers {
		if num%2 == 0 {
			numbersGroup["even"] = append(numbersGroup["even"], num)
		} else {
			numbersGroup["odd"] = append(numbersGroup["odd"], num)
		}
	}

	for key, slice := range numbersGroup {
		fmt.Print(key, "-> ")
		for _, number := range slice {
			fmt.Print(number, " ")
		}
		fmt.Println()
	}
}
