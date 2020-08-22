package main

import "fmt"

//Samsung is a type
type Samsung struct {
	Version string
}

//Info is a method
func (s Samsung) Info() string {
	return fmt.Sprintf("info: %s", s.Version)
}

//OnePlus is a type
type OnePlus struct {
	Samsung //Embed -> ไม่ใช่ inherit เพราะ Sumsung จะเก็บ instance ของ OnePlus ไม่ได้ แต่จะมี methods & properties of Samsung
	Version string
}

func main() {
	s := Samsung{Version: "s10+"}

	fmt.Println(s.Info())

	//o := OnePlus{Version: "7 Pro"}
	o := OnePlus{}
	o.Samsung.Version = "s11+"
	o.Version = "7 Pro"
	fmt.Printf("o: %#v\n", o) //o: main.OnePlus{Samsung:main.Samsung{Version:"s11+"}, Version:"7 Pro"}

	fmt.Println(o.Info()) //info: s11+
}
