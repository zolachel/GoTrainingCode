package main

import (
	"fmt"
	"sync"
)

func say(name string, wg *sync.WaitGroup) {
	fmt.Println("hello", name)
	defer wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go say("beer", &wg)

	fmt.Println("before wait")
	wg.Wait()
	fmt.Println("after wait")
}
