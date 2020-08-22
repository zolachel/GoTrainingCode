package main

import "fmt"

type Day int

func (d Day) Today() string {
	return fmt.Sprintf("today : %d", d)
}

func main() {
	//var d Day
	//d = 2
	d := Day(2)
	fmt.Println(d.Today())
	fmt.Println(int(d)) // convert กลับไปเป็น int ได้ด้วย
}
