package polymorphism

import "fmt"

// ต่อเนื่องมาจาก inheritance เมื่อเราสามารถถ่่ายทอดคุณสมบัติได้แล้ว แต่บาง method ก็สามารถพ้องรูป
// หรือทำงานได้หลายมุมมอง หลายแบบ จึงต้องมีการ overidding/ overloading เพื่อเปลี่ยนแปลงพฤติกรรม
// ภายใน method เหล่านั้น เช่น dog -> เห่า("woof") แต่พอสืบทอดมาที่ golden -> เห่า("โฮ่ง")

type (
	PlayerClass interface {
		Attack()
	}

	Warrior struct {
		Weapon string
	}

	Mage struct {
		Spell string
	}
)

func (w Warrior) Attack() {
	fmt.Println("Warrior is attacking with", w.Weapon)
}

func (m Mage) Attack() {
	fmt.Println("Mage is attacking with", m.Spell)
}

func UsePolymorphism() {
	warrior := Warrior{
		Weapon: "Sword",
	}

	mage := Mage{
		Spell: "Fireball",
	}

	warrior.Attack()
	mage.Attack()
}
