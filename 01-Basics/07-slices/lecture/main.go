package main

import "fmt"

func main() {

	// ============================================================
	// 1. Creating a Slice
	// ============================================================
	// A slice is a flexible view over an underlying array.
	//
	// Unlike arrays, the length of a slice is NOT part of its type.
	//
	// []int is the type of all int slices.
	// [3]int and [5]int are different array types.

	numbers := []int{10, 20, 30}

	fmt.Println(numbers)

	// ============================================================
	// 2. append
	// ============================================================
	// append adds elements to a slice.
	//
	// append may return a new slice, so we normally assign
	// the result back to the slice.

	arr := []int{10, 20, 30}

	arr = append(arr, 40)
	arr = append(arr, 50)

	fmt.Println(arr)

	// ============================================================
	// 3. Append Multiple Values
	// ============================================================
	// You can append multiple values at once.

	arr2 := []int{10, 20, 30}

	arr2 = append(arr2, 40, 50)

	fmt.Println(arr2)

	// ============================================================
	// 4. Slice Expressions
	// ============================================================
	// Syntax:
	//
	// slice[start:end]
	//
	// start is included.
	// end is NOT included.

	arr3 := []int{10, 20, 30, 40, 50}

	fmt.Println(arr3[1:3]) // [20 30]
	fmt.Println(arr3[1:])  // [20 30 40 50]
	fmt.Println(arr3[:3])  // [10 20 30]
	fmt.Println(arr3[:])   // [10 20 30 40 50]

	// ============================================================
	// 5. len vs cap
	// ============================================================
	// len = number of elements currently in the slice.
	//
	// cap = number of elements that can be used from the
	// slice's starting position before a new underlying array
	// may be required.

	arr4 := []int{10, 20, 30, 40, 50}

	fmt.Println("len:", len(arr4))
	fmt.Println("cap:", cap(arr4))

	// ============================================================
	// 6. make
	// ============================================================
	// make can create a slice with a specific length and capacity.
	//
	// make([]int, length, capacity)

	arr5 := make([]int, 3, 4)

	fmt.Println(arr5)

	fmt.Println("len:", len(arr5))
	fmt.Println("cap:", cap(arr5))

	arr5 = append(arr5, 40)

	fmt.Println(arr5)
	fmt.Println("len:", len(arr5))
	fmt.Println("cap:", cap(arr5))

	arr5 = append(arr5, 50)

	fmt.Println(arr5)
	fmt.Println("len:", len(arr5))
	fmt.Println("cap:", cap(arr5))

	// ============================================================
	// 7. Slice Sharing
	// ============================================================
	// A slice does NOT contain the actual array data itself.
	//
	// It points to an underlying array.
	//
	// Therefore, creating a slice from another slice can make
	// both slices share the same underlying array.

	arr6 := []int{10, 20, 30, 40, 50}

	a := arr6[1:4]

	fmt.Println("a:", a)

	a[0] = 1000

	// arr6 is also changed because a and arr6 share
	// the same underlying array.

	fmt.Println("arr6:", arr6)
	fmt.Println("a:", a)

	// ============================================================
	// 8. len vs cap of a Sub-slice
	// ============================================================
	// When creating:
	//
	// b := arr7[1:5]
	//
	// len(b) = 5 - 1 = 4
	//
	// cap(b) starts from index 1 and continues to the end
	// of the underlying array.
	//
	// Therefore:
	//
	// cap(b) = len(arr7) - 1

	arr7 := []int{
		10, 20, 30, 40, 50, 60, 70,
	}

	b := arr7[1:5]

	fmt.Println("arr7 len:", len(arr7))
	fmt.Println("arr7 cap:", cap(arr7))

	fmt.Println("b len:", len(b))
	fmt.Println("b cap:", cap(b))

	// ============================================================
	// 9. append and Shared Underlying Arrays
	// ============================================================
	// c initially shares the underlying array with arr8.

	arr8 := []int{10, 20, 30, 40, 50}

	c := arr8[1:3]

	fmt.Println("c:", c)
	fmt.Println("arr8:", arr8)

	// If c has enough capacity, append may write into the
	// same underlying array.
	//
	// If c runs out of capacity, Go may allocate a new
	// underlying array.

	c = append(c, 22)
	c = append(c, 33)
	c = append(c, 44)
	c = append(c, 55)
	c = append(c, 66)
	c = append(c, 77)

	fmt.Println("c:", c)
	fmt.Println("arr8:", arr8)

	// ============================================================
	// 10. copy
	// ============================================================
	// copy copies elements from one slice to another.
	//
	// Syntax:
	//
	// copy(destination, source)
	//
	// The destination must already have enough length
	// to receive the elements.

	source := []int{10, 20, 30, 40, 50}

	destination := make([]int, len(source))

	copy(destination, source)

	fmt.Println("source:", source)
	fmt.Println("destination:", destination)
}
