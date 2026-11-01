package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {

	// 1. make the outer slice with length dy
	//    (use make([][]uint8, dy))
	image := make([][]uint8, dy)

	// 2. loop y from 0 to dy-1:
	//    make the row: make([]uint8, dx), assign it to image[y]
	for y := 0; y < dy; y++ {
		image[y] = make([]uint8, dx)

		// 3. inside that loop, a second loop x from 0 to dx-1:
		//    set image[y][x] to your formula, wrapped in uint8(...)
		for x := 0; x < dx; x++ {
			// image[y][x] = uint8((x + y) / 2)
			image[y][x] = uint8(x * y * (x * y))
			// image[y][x] = uint8(x^y)
		}
	}

	// 4. return the outer slice
	return image
}

func main() {
	pic.Show(Pic)
}
