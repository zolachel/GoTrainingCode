package main

import "fmt"

type Vertext struct {
	X int
	Y int
	x int
	y int
}

func main() {
	v := Vertext{1, 2, 3, 4}

	p := &v

	fmt.Println(v)
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println((*p).X)
	fmt.Println((*p).x)
	fmt.Println(p.x) //same as (*p).x
}
