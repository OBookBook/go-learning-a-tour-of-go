package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	image := make([][]uint8, dy)

	for y := 0; y < dy; y++ {
		image[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			image[y][x] = uint8((x + y) / 2)
		}
	}

	return image
}

func main() {
	pic.Show(Pic)
}
