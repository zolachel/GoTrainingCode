package main

import "fmt"

func main() {
	var v interface{}
	v = 1
	fmt.Printf("%T %#v\n", v, v)
	v = "1"
	fmt.Printf("%T %#v\n", v, v)
	v = []int{1, 2, 3}
	fmt.Printf("%T %#v\n", v, v)
	v = map[string]int{"Beer": 1}
	fmt.Printf("%T %#v\n", v, v)
}
