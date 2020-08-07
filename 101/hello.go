package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("Hello, Beer", rand.Intn(100))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(9))
}
