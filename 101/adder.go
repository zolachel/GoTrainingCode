package main

import "fmt"

func adder() (func() int, func() int) {
	sum := 0

	return func() int {
			sum++
			return sum
		}, func() int {
			return sum
		}
}

func main() {
	fmt.Println("#")
	inc, curr := adder()
	fmt.Println(curr()) //0
	fmt.Println(inc())  //1
	fmt.Println(inc())  //2
	fmt.Println(curr()) //2

	fmt.Println("##")
	inc2, curr2 := adder()
	fmt.Println(curr2()) //0
	fmt.Println(inc2())  //1
	fmt.Println(inc2())  //2
	fmt.Println(inc2())  //3
	fmt.Println(curr2()) //3

	fmt.Println("###")
	fmt.Println(inc()) //3
	fmt.Println(inc()) //4
}
