package inheritance

import "fmt"

// ต่อเนื่องจากที่เราสามารถควบคุมการเข้าถึงข้อมูลระหว่าง object ได้แล้ว เรายังต้องแก้ปัญหา
// เรื่องที่หลาย ๆ class มี atrribute หรือ method คล้ายกันทำงานอยู่ ซึ่งก็ไม่ได้เหมือนทั้งหมด
// จึงไม่สามารถนำมารวมกันได้ เลยแก้ปัญหาเป็นการสืบทอดคุณสมบัติแทน เช่น golden extend dog
// ก็จะสามารถเดินและเห่า ได้เหมือน dog

type (
	Human struct{}

	Devil struct {
		Human
	}

	Elf struct {
		Human
	}
)

func (h Human) Walk() {
	fmt.Println("Walking")
}

func (h Human) Eat() {
	fmt.Println("Eating")
}

func (h Human) Breath() {
	fmt.Println("Breathing")
}

func (d Devil) Mutate() {
	fmt.Println("Mutating")
}

func (d Devil) Fly() {
	fmt.Println("Flying")
}

func (e Elf) MagicSpell() {
	fmt.Println("Casting a magic spell")
}

func UseInheritance() {
	human := Human{}
	devil := Devil{}
	elf := Elf{}

	fmt.Println("Human: ")
	human.Walk()
	human.Eat()
	human.Breath()
	fmt.Println()

	fmt.Println("Devil: ")
	devil.Walk()
	devil.Eat()
	devil.Breath()
	devil.Mutate()
	devil.Fly()
	fmt.Println()

	fmt.Println("Elf: ")
	elf.Walk()
	elf.Eat()
	elf.Breath()
	elf.MagicSpell()
}
