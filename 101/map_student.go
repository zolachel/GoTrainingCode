package main

import "fmt"

type Student struct {
	Name, Class string
}

var s = map[int]Student{
	965444: Student{
		"Adisorn", "6/1",
	},
	965333: Student{
		"Wiyada", "6/2",
	},
	999555: {"Chertam", "5/5"},
	999666: {"Crist", "5/4"},
}

func main() {
	s[966777] = Student{"Aya", "1/1"}              // add
	s[966888] = Student{Name: "Ayu", Class: "1/2"} // add

	s[966999] = Student{Class: "1/3", Name: "Ayo"}    // add
	s[966999] = Student{Class: "1/3", Name: "Ayoooo"} // update
	fmt.Println(s)

	s1 := s[965444]
	fmt.Println(s1) //{Adisorn 6/1}

	s1 = s[000000]
	fmt.Println(s1) // { }

	delete(s, 965444)
	fmt.Println(s)

	s2, ok := s[000000]
	fmt.Println(s2) // { }
	fmt.Println(ok) // false
}
