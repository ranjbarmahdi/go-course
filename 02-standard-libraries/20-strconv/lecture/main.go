package main

import (
	"fmt"
	"strconv"
)

// ============================================================
// strconv package
// ============================================================
//
// The strconv package is used for converting between strings
// and basic data types.
//
// Common functions:
//
// strconv.Atoi()
// strconv.Itoa()
// strconv.ParseFloat()
// strconv.FormatFloat()
// strconv.ParseInt()
// strconv.ParseBool()
//

func main() {

	// ============================================================
	// 1. String → Integer
	// ============================================================
	// Atoi means:
	//
	// ASCII to Integer
	//
	// Returns:
	// - parsed int
	// - error

	res, err := strconv.Atoi("-10")

	fmt.Println(res)
	fmt.Println(err)

	// ============================================================
	// 2. Integer → String
	// ============================================================
	// Itoa means:
	//
	// Integer to ASCII
	//
	// Returns a string.

	res1 := strconv.Itoa(1)

	fmt.Println(res1)

	// ============================================================
	// 3. String → Float
	// ============================================================
	// ParseFloat converts a string into a float.
	//
	// The second argument specifies the bit size:
	//
	// 32 -> float32
	// 64 -> float64

	res2, err2 := strconv.ParseFloat("1.99", 64)

	fmt.Println(res2)
	fmt.Println(err2)

	// ============================================================
	// 4. Float → String
	// ============================================================
	// FormatFloat converts a floating-point number into a string.
	//
	// Arguments:
	//
	// value
	// format
	// precision
	// bitSize
	//
	// 'f' -> decimal notation
	// 4   -> 4 digits after the decimal point
	// 64  -> float64

	res3 := strconv.FormatFloat(
		2.9826123,
		'f',
		4,
		64,
	)

	fmt.Println(res3)

	// ============================================================
	// 5. ParseInt
	// ============================================================
	// ParseInt converts a string into an integer.
	//
	// Arguments:
	//
	// value
	// base
	// bitSize
	//
	// Base:
	//
	// 10 -> decimal
	// 2  -> binary
	// 8  -> octal
	// 16 -> hexadecimal

	res4, err4 := strconv.ParseInt(
		"1019",
		10,
		64,
	)

	fmt.Println(res4)
	fmt.Println(err4)

	// Hexadecimal → Integer
	//
	// "FF" in base 16 = 255

	res5, err5 := strconv.ParseInt(
		"FF",
		16,
		64,
	)

	fmt.Println(res5)
	fmt.Println(err5)

	// ============================================================
	// 6. String → Boolean
	// ============================================================
	// ParseBool converts a string into a bool.
	//
	// Example:
	//
	// "true"  -> true
	// "false" -> false
	//
	// It returns an error if the string is not a valid
	// boolean representation.

	res6, err6 := strconv.ParseBool("true")

	fmt.Println(res6)
	fmt.Println(err6)

	// ============================================================
	// 7. Conversion Errors
	// ============================================================
	// Parsing can fail when the string does not represent
	// the requested type.

	number, err7 := strconv.Atoi("hello")

	fmt.Println(number)
	fmt.Println(err7)

	floatValue, err8 := strconv.ParseFloat("abc", 64)

	fmt.Println(floatValue)
	fmt.Println(err8)

	boolValue, err9 := strconv.ParseBool("hello")

	fmt.Println(boolValue)
	fmt.Println(err9)
}
