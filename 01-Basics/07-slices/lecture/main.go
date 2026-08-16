package main

import "fmt"

func main() {
	fmt.Println("1. Creating a slice")
	numbers := []int{10, 20, 30}
	fmt.Println(numbers)
	//the slice doesn't encode its length in its type.

	fmt.Println("2. append")
	arr := []int{10, 20, 30}
	arr = append(arr, 40)
	arr = append(arr, 50)
	fmt.Println(arr)

	fmt.Println("3. Append multiple values")
	arr2 := []int{10, 20, 30}
	arr2 = append(arr2, 40, 50)
	fmt.Println(arr2)

	fmt.Println("4. Slice expressions")
	arr3 := []int{10, 20, 30, 40, 50}
	fmt.Println(arr3[1:3])
	fmt.Println(arr3[1:])
	fmt.Println(arr3[:3])
	fmt.Println(arr3[:])

	fmt.Println("5. len vs cap")
	arr4 := []int{10, 20, 30, 40, 50}
	fmt.Println(len(arr4))
	fmt.Println(cap(arr4))
	// len: How many elements are currently in the slice?
	// cap: How much capacity does the slice currently have before it needs a new underlying array?

	fmt.Println("6. make")
	arr5 := make([]int, 3, 4)
	fmt.Println(arr5)

	fmt.Println(len(arr5))
	fmt.Println(cap(arr5))

	arr5 = append(arr5, 40)
	fmt.Println(cap(arr5))

	arr5 = append(arr5, 50)
	fmt.Println(cap(arr5))

	// Deep in slice

	// Important
	fmt.Println("7. Slice sharing")
	arr6 := []int{10, 20, 30, 40, 50}
	a := arr6[1:4] // reference coy
	a[0] = 1000
	fmt.Println(arr6)

	// Important
	fmt.Println("8. len vs cap")
	arr7 := []int{10, 20, 30, 40, 50, 60, 70}

	b := arr7[1:5]

	fmt.Println("arr7 len:", len(arr7))
	fmt.Println("arr7 cap:", cap(arr7))

	fmt.Println("b len:", len(b))
	fmt.Println("b cap:", cap(b))

	// Important
	fmt.Println("9. Append and shared arrays")
	arr8 := []int{10, 20, 30, 40, 50}
	c := arr8[1:3]

	fmt.Println(c)
	fmt.Println(arr8)

	c = append(c, 22)
	c = append(c, 33)
	c = append(c, 44)
	c = append(c, 55)
	c = append(c, 66)
	c = append(c, 77)

	fmt.Println(c)
	fmt.Println(arr8)

	fmt.Println("10. Copy")
	source := []int{10, 20, 30, 40, 50}
	destination := make([]int, len(source))

	copy(destination, source)

	fmt.Println(destination)
}
