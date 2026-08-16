package main

import "fmt"

func main() {
	a := []int{1, 1, 3, 2, 3, 4, 5}
	b := []int{3, 4, 5, 6, 7}
	common := []int{}

	for _, va := range a {

		existsInCommon := false
		for _, vc := range common {
			if va == vc {
				existsInCommon = true
				break
			}
		}

		if existsInCommon {
			continue
		}

		for _, vb := range b {
			if va == vb {
				common = append(common, va)
			}
		}
	}

	fmt.Println(common)
}
