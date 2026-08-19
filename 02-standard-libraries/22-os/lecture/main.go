package main

import (
	"fmt"
	"os"
)

// The `os` package provides functions for interacting with the operating system.
//
// Common backend use cases:
// - Environment variables
// - Files
// - Command-line arguments
// - File information
// - Creating/removing files and directories

func main() {

	// ============================================================
	// 1. Environment Variables ⭐⭐⭐
	// ============================================================
	// `os.Getenv()` returns the value of an environment variable.
	//
	// If the variable does not exist, it returns an empty string.

	port := os.Getenv("port")

	fmt.Println(port)

	// ============================================================
	// 2. os.LookupEnv()
	// ============================================================
	// `os.LookupEnv()` returns:
	// - the value
	// - whether the variable exists
	//
	// This allows us to distinguish between:
	// - variable does not exist
	// - variable exists but contains an empty string

	value, exists := os.LookupEnv("PATH")

	fmt.Println(value)
	fmt.Println(exists)

	// ============================================================
	// 3. os.Args
	// ============================================================
	// `os.Args` contains the command-line arguments passed
	// to the program.
	//
	// Example:
	//
	// go run main.go hello 123
	//
	// os.Args:
	// [program-path hello 123]

	fmt.Println(os.Args)

	// ============================================================
	// 4. Reading a File
	// ============================================================
	// `os.ReadFile()` reads the entire file into a []byte.
	//
	// Convert []byte to string when you want to print text.

	data, err := os.ReadFile(
		"C:\\Users\\Mehdi\\Desktop\\go-crash\\02-standard-libraries\\22-os\\lecture\\test.txt",
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))

	// ============================================================
	// 5. Writing a File
	// ============================================================
	// `os.WriteFile()` writes data to a file.
	//
	// If the file does not exist, it creates it.
	//
	// If the file already exists, its contents are replaced.
	//
	// `0644` is the file permission.

	err = os.WriteFile(
		"C:\\Users\\Mehdi\\Desktop\\go-crash\\02-standard-libraries\\22-os\\lecture\\write.txt",
		[]byte("Hello\nMehdi"),
		0644,
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("File written")

	// ============================================================
	// 6. Creating a File
	// ============================================================
	// `os.Create()` creates a file and returns *os.File.
	//
	// If the file already exists, its contents are truncated.

	file, err := os.Create(
		"C:\\Users\\Mehdi\\Desktop\\go-crash\\02-standard-libraries\\22-os\\lecture\\md.md",
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(file)

	file.Close()

	// ============================================================
	// 7. Removing a File
	// ============================================================
	// `os.Remove()` removes a file or an empty directory.

	err = os.Remove(
		"C:\\Users\\Mehdi\\Desktop\\go-crash\\02-standard-libraries\\22-os\\lecture\\md.md",
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("File removed")

	// ============================================================
	// 8. File Information
	// ============================================================
	// `os.Stat()` returns information about a file or directory.
	//
	// The returned `os.FileInfo` provides information such as:
	// - Name
	// - Size
	// - IsDir
	// - Permissions
	// - Modification time

	info, err := os.Stat(
		"C:\\Users\\Mehdi\\Desktop\\go-crash\\02-standard-libraries\\22-os\\lecture\\test.txt",
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Name:", info.Name())
	fmt.Println("Size:", info.Size())
	fmt.Println("Is directory:", info.IsDir())
}
