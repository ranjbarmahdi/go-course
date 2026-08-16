package main

import "fmt"

type Counter struct {
	Value int
}

func (c *Counter) Increment() {
	c.Value++
}

func main() {
	counter := Counter{}
	counter.Increment()
	counter.Increment()
	counter.Increment()
	counter.Increment()
	fmt.Println(counter)
}
