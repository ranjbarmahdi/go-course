/*
Problem:

Create a generic Stack:

type Stack[T any] struct {
    values []T
}

Create these methods:

func (s *Stack[T]) Push(value T)

func (s *Stack[T]) Pop() (T, bool)

Requirements:

Push():

- Add a value to the stack.

Pop():

- Return the last value.
- Remove it from the stack.
- Return true if a value exists.
- If the stack is empty:
    - return the zero value of T
    - return false

In main():

Create Stack[int].

Push:

10
20
30

Then call Pop() twice.

Expected output:

30 true
20 true
*/

package main

import "fmt"

type Stack[T any] struct {
	values []T
}

func (s *Stack[T]) Push(value T) {
	s.values = append(s.values, value)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.values) == 0 {
		var zero T
		return zero, false
	}
	res := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return res, true
}

func main() {
	stack := Stack[int]{
		values: []int{1, 2, 3, 4},
	}
	stack.Push(5)
	stack.Push(7)

	fmt.Println(stack)

	res, exists := stack.Pop()
}
