package main

import (
	"fmt"
	"math"
)

func compute(param func(float64, float64) float64) float64 {
	return param(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x + y)
	}

	fmt.Println(hypot(10, 6))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
