package main

import "golang.org/x/tour/pic"

func renderPic(dx, dy int) [][]uint8 {
	picture := [][]uint8{}

	for x := 0; x < dx; x++ {
		row := []uint8{}

		for y := 0; y < dy; y++ {
			if x%3 == 0 {
				row = append(row, uint8((x+y)/2))
			} else if x%3 == 1 {
				row = append(row, uint8(y^x/3))
			} else {
				row = append(row, uint8(x))
			}
		}

		picture = append(picture, row)
	}

	return picture
}

func main() {
	pic.Show(renderPic)
}
