package main

import "fmt"

//Person is a representation of a person
type Person struct {
	name   string
	friend map[string]int
	shape  []int
}

//Walk is a method
func (p Person) Walk() {
	fmt.Println(p.name + " walking")
}

//Eat is a method
func (p Person) Eat() {
	fmt.Println(p.name + " eating")
}

//Geeting is a method
func (p Person) Geeting() {
	fmt.Println("hello " + p.name)
}

//Name is a method
func (p Person) Name() string {
	return p.name
}

//SetName is a method
func (p *Person) SetName(n string) {
	p.name = n
}

//SetFriend is a method
func (p Person) SetFriend(friendName string, friendValue int) {
	p.friend[friendName] = friendValue
}

//SetShape is a method
func (p *Person) SetShape(sh []int) {
	p.shape = sh
}

//SetWaist is a method
func (p Person) SetWaist(waist int) {
	p.shape[1] = waist
}

func main() {
	p := Person{name: "Beer", friend: map[string]int{"Bank": 1}, shape: []int{34, 24, 36}}
	p.Walk()
	p.Eat()
	p.Geeting()

	fmt.Printf("Before %#v\n", p)
	p.SetName("Aya")
	p.SetFriend("Bank", 2)
	p.SetShape([]int{77, 77, 77})
	p.SetWaist(99)
	fmt.Printf("After %#v\n", p)
}
