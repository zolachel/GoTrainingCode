package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	/*value := math.Pow(x, n)

	if value < lim {
		return value
	}*/

	if value := math.Pow(x, n); value < lim {
		return value
	} else {
		fmt.Println("else value:", value)
	}

	// cann't refer to value here

	return lim
}

func main() {
	fmt.Println(pow(5, 2, 24))
}
