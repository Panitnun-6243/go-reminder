package pointer

import (
	"fmt"
)

func UsePointer(p *int) {
	// pointer ก็แค่ตัวแปรหนึ่งที่ใช้เก็บ address ของตัวแปรตัวอื่นเป็นชั้น ๆ ไปเรื่อย ๆ เช่น a = 1, *b เก็บ address a, **c เก็บ address *b, ***d เก็บ address **c
	// ไม่สามารถข้ามไปเก็บแบบ **c เก็บ address a
	// ประกาศได้ 2 แบบ -> ศัพท์เทคนิคจะเรียกว่าการ allowcate หรือก็คือการจอง memory
	// แบบแรก
	var a *string
	a1 := "Hello Pointer"
	a = &a1
	fmt.Println("This is the real value of a1 ->", a1)
	fmt.Println("This is the address of a1 ->", &a1)
	fmt.Println("This is the real value of a which is the address of a1 ->", a)
	fmt.Println("This is the real address of a which is not the same address of a1 ->", &a)

	// แบบสอง
	b := new(*string) // double pointer
	b = &a
	fmt.Println("This is the real value of b which is the address of a ->", b)
	fmt.Println("This is the first value that have been sequently pointed from a1 -> a -> b which **b and *a are", **b, "and", *a)
	// pass by pointer
	*p *= 2
}
