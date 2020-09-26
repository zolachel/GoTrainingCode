package main

import (
	"fmt"
)

func say(name string, c chan int) {
	fmt.Println("hello", name)
	c <- 1
}

func main() {
	c := make(chan int) // channel ที่ไม่มี buffer
	//c := make(chan int, 3) // channel ที่่มี 3 buffer

	go say("beer", c)

	fmt.Println("wait")
	val := <-c //default of channel is blocking
	fmt.Println("done!", val)
}
