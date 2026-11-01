package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)
	fmt.Println(len(q)) //6
	
	//q[6] = 88  // this fails runtime error: index out of range [6] with length 6
	q[5] = 88 
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
	
	shift := s[1:4]  
	shift[0].i = 23
	shift[0].b = false
	fmt.Println(s)

}  