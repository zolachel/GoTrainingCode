package main

import "fmt"

func main() {
	var i interface{} = "hello"

	if s, ok := i.(string); ok {
		fmt.Println("s =", s)
	}

	if t, ok := i.(int); ok {
		fmt.Println("t =", t)
	} else {
		fmt.Println("can't assert to t")
	}
}
