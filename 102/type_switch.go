package main

import "fmt"

//Apple is a type
type Apple struct {
	Version string
}

func main() {
	var i interface{} = "hello"
	switch v := i.(type) {
	case int:
		fmt.Printf("%T %d\n", v, v)
	case string:
		fmt.Printf("%T %s\n", v, v)
	default:
		fmt.Println("undefined type")
	}

	var j interface{} = Apple{Version: "iphone X"}
	//var j interface{} = &Apple{Version: "iphone X"}
	switch w := j.(type) {
	case Apple:
		fmt.Printf("%T %v not pointer\n", w, w.Version)
	case *Apple:
		fmt.Printf("%T %v pointer\n", w, w.Version)
	default:
		fmt.Println("undefined phone")
	}
}
