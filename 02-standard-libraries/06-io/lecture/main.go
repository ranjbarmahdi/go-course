package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// ============================================================
// 1. io.Reader ⭐⭐⭐
// ============================================================
// `io.Reader` represents something that can read data.
//
// Interface:
//
// type Reader interface {
//     Read(p []byte) (n int, err error)
// }
//
// A Reader produces data.
//
// Examples:
// - Files
// - HTTP response bodies
// - Strings
// - Network connections
//
// Important:
// You usually don't need to implement Reader yourself.
// Many Go types already implement it.

func main() {

	fmt.Println("1. io.Reader")

	file, err := os.Open(
		"C:\\Users\\Mehdi\\Desktop\\go-crash\\02-standard-libraries\\23-io\\lecture\\test.txt",
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	data, err := io.ReadAll(file)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))

	// ============================================================
	// 2. io.Writer ⭐⭐⭐
	// ============================================================
	// `io.Writer` represents something that can receive data.
	//
	// Interface:
	//
	// type Writer interface {
	//     Write(p []byte) (n int, err error)
	// }
	//
	// A Writer receives data.

	fmt.Println("\n2. io.Writer")

	file2, err := os.Create(
		"C:\\Users\\Mehdi\\Desktop\\go-crash\\02-standard-libraries\\23-io\\lecture\\write.txt",
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file2.Close()

	_, err = file2.Write([]byte("Hello Go"))

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Data written successfully")

	// ============================================================
	// 3. io.Copy() ⭐⭐⭐
	// ============================================================
	// `io.Copy()` copies data from a Reader to a Writer.
	//
	// Reader → Writer
	//
	// io.Copy(destination, source)

	fmt.Println("\n3. io.Copy()")

	source, err := os.Open(
		"C:\\Users\\Mehdi\\Desktop\\go-crash\\02-standard-libraries\\23-io\\lecture\\test.txt",
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer source.Close()

	destination, err := os.Create(
		"C:\\Users\\Mehdi\\Desktop\\go-crash\\02-standard-libraries\\23-io\\lecture\\copy.txt",
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer destination.Close()

	written, err := io.Copy(destination, source)

	fmt.Println("Bytes copied:", written)
	fmt.Println("Error:", err)

	// ============================================================
	// 4. Reader → Writer
	// ============================================================
	// The most important concept in the `io` package:
	//
	//                  io.Copy()
	//
	// Reader ──────────────────────→ Writer
	//
	// Example:
	//
	// test.txt
	//    ↓
	// Reader
	//    ↓
	// io.Copy()
	//    ↓
	// Writer
	//    ↓
	// copy.txt
	//
	// This same pattern works with many different sources
	// and destinations.

	fmt.Println("\n4. Reader → Writer")

	// Examples:
	//
	// File    → File
	// File    → Terminal
	// HTTP    → File
	// Memory  → File
	// Network → File

	// ============================================================
	// 5. os.Stdout is an io.Writer
	// ============================================================
	// `os.Stdout` represents the program's standard output.
	//
	// It implements io.Writer.
	//
	// Therefore, we can write directly to the terminal.

	fmt.Println("\n5. os.Stdout")

	_, err = os.Stdout.Write(
		[]byte("Hello from os.Stdout\n"),
	)

	if err != nil {
		fmt.Println(err)
	}

	// ============================================================
	// 6. io.WriteString()
	// ============================================================
	// `io.WriteString()` writes a string to an io.Writer.
	//
	// Instead of:
	//
	// writer.Write([]byte("Hello"))
	//
	// We can use:
	//
	// io.WriteString(writer, "Hello")

	fmt.Println("\n6. io.WriteString()")

	file3, err := os.Create(
		"C:\\Users\\Mehdi\\Desktop\\go-crash\\02-standard-libraries\\23-io\\lecture\\string.txt",
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file3.Close()

	_, err = io.WriteString(
		file3,
		"Hello from io.WriteString()\n",
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("String written successfully")

	// ============================================================
	// 7. strings.Reader
	// ============================================================
	// `strings.NewReader()` creates a Reader from a string.
	//
	// This is useful when you want to treat a string
	// as an io.Reader.

	fmt.Println("\n7. strings.Reader")

	reader := strings.NewReader(
		"Hello from strings.Reader",
	)

	data, err = io.ReadAll(reader)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))

	// ============================================================
	// 8. io.LimitReader()
	// ============================================================
	// `io.LimitReader()` creates a Reader that reads
	// at most N bytes from another Reader.
	//
	// io.LimitReader(reader, n)

	fmt.Println("\n8. io.LimitReader()")

	reader = strings.NewReader(
		"Hello Golang Backend Developer",
	)

	limitedReader := io.LimitReader(
		reader,
		5,
	)

	data, err = io.ReadAll(limitedReader)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))

	// Output:
	//
	// Hello
	//
	// Only the first 5 bytes are available.

	// ============================================================
	// 9. Reader Position / EOF ⭐⭐⭐
	// ============================================================
	// Readers have a current position.
	//
	// When you read from a file, the position moves forward.
	//
	// After reading the entire file, the Reader is at EOF
	// (End Of File).
	//
	// Reading again does not automatically start from the beginning.

	fmt.Println("\n9. Reader position")

	file4, err := os.Open(
		"C:\\Users\\Mehdi\\Desktop\\go-crash\\02-standard-libraries\\23-io\\lecture\\test.txt",
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file4.Close()

	// First read

	data, err = io.ReadAll(file4)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("First read:")
	fmt.Println(string(data))

	// The Reader is now at the end of the file.

	data, err = io.ReadAll(file4)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Second read:")
	fmt.Println(string(data))

	// The second read is empty because the Reader
	// is already at the end.

	// ============================================================
	// 10. Reopen File To Read Again
	// ============================================================
	// Reopening the file creates a new Reader position
	// starting from the beginning.

	fmt.Println("\n10. Reopen file")

	file5, err := os.Open(
		"C:\\Users\\Mehdi\\Desktop\\go-crash\\02-standard-libraries\\23-io\\lecture\\test.txt",
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file5.Close()

	_, err = io.Copy(
		os.Stdout,
		file5,
	)

	if err != nil {
		fmt.Println(err)
	}

	// ============================================================
	// 11. Important Mental Model ⭐⭐⭐
	// ============================================================
	//
	// Reader:
	//     produces data
	//
	// Writer:
	//     receives data
	//
	// io.Copy():
	//     Reader → Writer
	//
	// Think about data flowing from one place to another:
	//
	//     File ─────────→ File
	//
	//     File ─────────→ Terminal
	//
	//     HTTP ─────────→ File
	//
	//     Memory ───────→ File
	//
	//     Network ──────→ File
	//
	// The important part is not the concrete type.
	//
	// The important part is:
	//
	//     Does it implement io.Reader?
	//
	//     Does it implement io.Writer?

	fmt.Println("\n11. Important mental model")

	fmt.Println(`
	Reader:
		produces data

	Writer:
		receives data

	io.Copy():
		Reader → Writer

	Examples:

		File → File
		File → Terminal
		HTTP → File
		File → HTTP
		Memory → File
		Network → File
	`)
}
