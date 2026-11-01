package main

import "fmt"

type Vertex struct {
	X, Y int
}

func scaleByPointer(v *Vertex, f int) {
	v.X *= f
	v.Y *= f
}

func scaleByValue(v Vertex, f int) {
	v.X *= f
	v.Y *= f
}

func scaled(v Vertex, f int) Vertex {
	return Vertex{v.X * f, v.Y * f}
}

func main() {
	a := Vertex{10, 20}
	scaleByPointer(&a, 10)
	fmt.Println(a) // {100 200}

	// by passing a copy no one ever gets the result of the function
	b := Vertex{1, 2}
	scaleByValue(b, 10)
	fmt.Println(b) // {1 2}

	// to illustrate the lesson, you avoid pointers by returning a value instead
	b = scaled(b, 10) // reassign the result
	fmt.Println(b)    // {10 20}

}
