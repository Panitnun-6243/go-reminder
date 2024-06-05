package oop

import "fmt"

// oop คือ programming paradigm ประเภทนึงที่นำ concept ของ class และ object มาใช้งาน
// class -> แบบแปลน, object -> instance ที่สร้างตามแปลนของ class

// เครื่องบิน -> spitfire, boeing
type Airplane struct {
	// attribute
	Name string
}

// implement behaviour (method) ในคลาส Airplane
func (a *Airplane) fly() {
	fmt.Println(a.Name, "is flying...")
}

func (a *Airplane) fire() {
	fmt.Printf("%s is firing missile!!!", a.Name)
}

func UseOop() {
	boeingAirplane := Airplane{
		Name: "Boeing Airplane",
	}
	f16 := Airplane{
		Name: "F16",
	}
	fmt.Println(boeingAirplane.Name)
	boeingAirplane.fly()
	f16.fire()

}
