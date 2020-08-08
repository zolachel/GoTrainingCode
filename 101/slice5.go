package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	fmt.Println(s[2])
	s[2] = 123
	fmt.Println(s[2])

	s = append(s, 4, 2, 3)
	fmt.Println(s)

	data := []int{1, 2, 3, 4}

	result := append(s, data...)

	fmt.Println(result)

	s = []int{}

	if s == nil {
		fmt.Println("s is nil")
	} else {
		fmt.Println("s is ", s)
	}

	s = make([]int, 5)
	fmt.Println(s) //[0 0 0 0 0]
	printSlice(s)  //len=5 cap=5 [0 0 0 0 0]

	s = make([]int, 5, 11)
	fmt.Println(s) //[0 0 0 0 0]
	printSlice(s)  //len=5 cap=11 [0 0 0 0 0]

	s = append(s, 1, 2, 3, 4, 5, 6, 7)
	printSlice(s) //len=12 cap=22 [0 0 0 0 0 1 2 3 4 5 6 7]

	//s = make([]int, 5, 4) len larger than cap in make([]int)
}
