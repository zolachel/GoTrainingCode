package main

import "fmt"

type square struct {
	width float64
}

func (s square) area() float64 {
	return s.width * s.width
}

type rectangle struct {
	width float64
	high  float64
}

func (r rectangle) area() float64 {
	return r.width * r.high
}

func showArea(s areaer) {
	fmt.Println(s.area())
}

type areaer interface {
	area() float64
}

func main() {
	s := square{width: 4}
	showArea(s)
	r := rectangle{width: 4, high: 5}
	showArea(r)
}
