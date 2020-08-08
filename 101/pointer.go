package main

import "fmt"

func main() {
	i, j, k := 55, 66, 77

	fmt.Println("address i: ", &i)
	fmt.Println("address j: ", &j)
	fmt.Println("address k: ", &k)

	var p *int
	fmt.Println("p value: ", p)

	//*p = 44 -> invalid memory address or nil pointer dereference

	p = &i
	fmt.Println("p value: ", p)
	fmt.Println("p address: ", &p)
	fmt.Println("p value inside address: ", *p)

	*p = 88

	fmt.Println("p value: ", p)
	fmt.Println("p value inside address: ", *p)
	fmt.Println("i value: ", i)

	//x := "string"
	//p = &x -> "cannot use &x (type *string) as type *int in assignment"

	i, j = 42, 2701

	p = &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}
