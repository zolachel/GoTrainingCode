package main

import "fmt"

type rectangle struct {
	width  int
	length int
}

/*func area(rec rectangle) int {
	return rec.width * rec.length
}

func main() {
	rec := rectangle{3, 4}
	a := area(rec)
	fmt.Println(a)
}*/

func (rect rectangle) area() int { //สลับ rec ไปอยู่ข้างหน้าเพื่อทำให้ area เป็น method ของ rectangle, (rect rectangle) เรียกว่า receiver
	return rect.width * rect.length
}

func main() {
	rec := rectangle{3, 4}
	a := rec.area() //เปลี่ยนวิธีเรียก method
	fmt.Println(a)
}
