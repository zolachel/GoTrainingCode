package main

import "fmt"

type student struct {
	name string
	id   int
}

func (s student) String() string {
	return fmt.Sprintf("Hello %d : %s\n", s.id, s.name)
}

func main() {
	s := student{id: 35, name: "Beer"}
	fmt.Println(s) // use the String method
}
