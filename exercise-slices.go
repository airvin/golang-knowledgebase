package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {

    pict := make([][]uint8,dy)
	
	for x := range pict {
		pict[x] = make([]uint8,dx)
		for y := range pict[x] {
			pict[x][y] = uint8(x^y)
		}
	}

    return pict
}

func main() {
	pic.Show(Pic)
}