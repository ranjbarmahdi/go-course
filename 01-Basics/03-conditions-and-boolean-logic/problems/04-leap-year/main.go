package main

import "fmt"

func main() {
	year := 2023

	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		fmt.Println(year, "--> leap year")
	} else {
		fmt.Println(year, "--> not leap year")
	}
}
