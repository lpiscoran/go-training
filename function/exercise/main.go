package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
// Implement a fibonacci function that returns a function (a closure) that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...)
func fibonacci() func() int {
	current, next := 0, 1

	return func() int {
		defer func() {
			current, next = next, current+next
		}()
		return current
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
