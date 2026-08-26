/*
Problem:

Create these interfaces:

type Reader interface {
    Read()
}

type Closer interface {
    Close()
}

Create a combined interface:

type ReadCloser interface {
    Reader
    Closer
}

Create a File struct.

Implement both methods:

    Read()
    Close()

Requirements:

- File must satisfy ReadCloser.
- Create a File value.
- Store it in a ReadCloser variable.
- Call Read().
- Call Close().

Expected output:

Reading...
Closing...

Important:

ReadCloser embeds both Reader and Closer.

Therefore, any type assigned to ReadCloser
must implement BOTH methods:

    Read()
    Close()
*/

package main

import "fmt"

type Reader interface {
	Read()
}

type Closer interface {
	Close()
}

type ReadCloser interface {
	Reader
	Closer
}

type File struct{}

func (f File) Read() {
	fmt.Println("Reading...")
}

func (f File) Close() {
	fmt.Println("Closing...")
}

func readAndClose(r ReadCloser) {
	r.Read()
	r.Close()
}

func main() {
	file := File{}
	readAndClose(file)
}
