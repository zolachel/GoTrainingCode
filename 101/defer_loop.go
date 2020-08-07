package main

import "fmt"

func main() {
	fmt.Println("counting")

	/*for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}*/

	/*for i := 0; i < 10; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}*/

	for i := 0; i < 10; i++ {
		defer func(x int) {
			fmt.Println(x)
		}(i)
	}

	fmt.Println("done")
}
