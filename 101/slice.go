package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	s := primes[1:4]           //[]int{3, 5, 7}
	t := primes[2:4]           //[]int{5, 7}
	u := primes[2:len(primes)] //[]int{5, 7, 11, 13}
	v := primes[:]             //[]int{2, 3, 5, 7, 11, 13}
	w := primes[1:]            //[]int{3, 5, 7, 11, 13}
	x := primes[:4]            //[]int{2, 3, 5, 7}

	fmt.Printf("%#v\n", primes) //[6]int{2, 3, 5, 7, 11, 13}
	fmt.Printf("%#v\n", s)
	fmt.Printf("%#v\n", t)
	fmt.Printf("%#v\n", u)
	fmt.Printf("%#v\n", v)
	fmt.Printf("%#v\n", w)
	fmt.Printf("%#v\n", x)

	s[0] = 6
	fmt.Printf("%#v\n", primes) //[6]int{2, 6, 5, 7, 11, 13} เปลี่ยนค่าจาก 3 เป็น 6 ด้วย!
	fmt.Printf("%#v\n", s)      //[]int{6, 5, 7}
}
