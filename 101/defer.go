package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("const.go")

	if err != nil {
		log.Fatal(err)
	}

	defer f.close()

	defer fmt.Println("world") //ทำตอนจะจบ function

	fmt.Println("Hello")
}
