package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func nakedReturn(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	a, b := swap("hi", "beer")
	fmt.Println(a, b)

	fmt.Println(nakedReturn(9))
}
