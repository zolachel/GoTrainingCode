package main

import "fmt"

type student struct {
	name, class string
}

func (stu student) info() {
	fmt.Printf("name=%v class=%v \n", stu.name, stu.class)
}

func main() {
	s := map[int]student{999555: {"Chertam", "5/5"},
		999666: {"Crist", "5/4"}}

	s[966777] = student{"Aya", "1/1"} // add
	s[966888] = student{"Ayu", "1/2"} // add

	for _, stud := range s {
		stud.info()
	}
}

// create method info of student --> print name and class
// for range loop, invoke info method
