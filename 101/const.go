package main

import "fmt"

const Pi = 3.14

const (
	monday = iota
	tuesday
	wednesday
)

const (
	jan = iota * 50
	feb
	mar
	apr
)

func main() {
	const world = "World"
	fmt.Println("Hello", world)
	fmt.Println("Happy", Pi, "Day")

	const truth = true
	fmt.Println("Go rules?", truth)

	fmt.Println(monday, tuesday, wednesday) //0 1 2

	fmt.Println(jan, feb, mar, apr) // 1 2 3
}
