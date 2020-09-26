package main

import (
	"fmt"
	"time"
)

type Ball struct {
	hits int
}

func player(name string, table chan *Ball) {
	for {
		ball := <-table
		ball.hits++

		fmt.Println(name, ball.hits)
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}

func main() {
	table := make(chan *Ball)
	go player("Beer", table)
	go player("Yui", table)

	ball := new(Ball)
	table <- ball

	time.Sleep(2 * time.Second)
	<-table
	fmt.Println("done")
}
