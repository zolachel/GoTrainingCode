package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(s)                                 //[2 3 5 7 11 13]
	fmt.Printf("len=%d cap=%d \n", len(s), cap(s)) //len=6 cap=6

	s = s[1:4]
	fmt.Println(s)                                 //[3 5 7]
	fmt.Printf("len=%d cap=%d \n", len(s), cap(s)) //len=3 cap=5
	//cap is capacity
	//ตอน slice ตัดแต่ส่วนหัวไม่ตัดท้าย ดังนั้น cap จะมีส่วนท้ายเหลืออยู่ กรณีนี้ก็ตัด 2 ออกยังเหลือ 3 5 7 11 13
	s = s[:2]
	fmt.Println(s)                                 //[3 5]
	fmt.Printf("len=%d cap=%d \n", len(s), cap(s)) //len=2 cap=5

	s = s[1:]
	fmt.Println(s)                                 //[5]
	fmt.Printf("len=%d cap=%d \n", len(s), cap(s)) //len=1 cap=4

	s = append(s, 99)
	fmt.Println(s)                                 //[5 99]
	fmt.Printf("len=%d cap=%d \n", len(s), cap(s)) //len=2 cap=4

	s = append(s, 88, 77)
	fmt.Println(s)                                 //[5 99 88 77]
	fmt.Printf("len=%d cap=%d \n", len(s), cap(s)) //len=4 cap=4

	s = append(s, 66)
	fmt.Println(s)                                 //[5 99 88 77 66]
	fmt.Printf("len=%d cap=%d \n", len(s), cap(s)) //len=5 cap=8 จะเห็นว่า cap เพิ่มเป็น double
}
