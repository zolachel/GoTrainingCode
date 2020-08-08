package main

import "fmt"

type Vertext struct {
	x int
	y int
}

func main() {
	//v := Vertext{1, 2}
	v := Vertext{x: 1, y: 2}
	fmt.Println(v)
	fmt.Printf("%+v\n", v)
	fmt.Printf("%#v\n", v)
	fmt.Printf("%t\n", v)
	fmt.Printf("%T\n", v)
	fmt.Println(v.x)
	fmt.Println(&v)
	fmt.Println(&v.y)

	v.x = 15
	v.y = 20
	fmt.Printf("%+v\n", v)

	p := &v
}
