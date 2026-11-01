package main

import (
	"fmt"
	"slices"
)

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)

	s = append(s, 2, 3, 11, 13)
	printSlice(s)  // unsorted len=6 cap=8 [5 7 2 3 11 13]
	slices.Sort(s)
	printSlice(s) // sorted len=6 cap=8 [2 3 5 7 11 13]

	s = s[0:6]
	printSlice(s)

	println("s[0]:", s[0])
	// println("s[7]:", s[7]) //index out of range [7] with length 6

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
