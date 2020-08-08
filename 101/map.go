package main

import "fmt"

func main() {
	var m map[string]string
	m = make(map[string]string)
	m["nong"] = "AdisornO"
	fmt.Println(m["nong"])
}
