package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	
	var pict [][]uint8

	for i := 0; i < dx; i++ {
		pict = append(pict,[]uint8{1})
		for j := 0; j < dy; j++ {
			pict[i] = append(pict[i],uint8(i*j))
		}
		fmt.Println(pict[i])
	}
}

func main() {
	pic.Show(Pic)
}
