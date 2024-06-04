package generics

import "fmt"

// generic เข้ามาช่วยเรื่องการที่ฟังชั่นก์สามารถรับหลาย ๆ data type เข้าไปประมวลผลได้
// เช่น ฟังชั่นบวกเลขทั้งหมดใน array โดยที่สามารถเป็น int array หรือ float array ก็ได้
// โดย T = template type เพื่อแปลงฟังชั้นก์ให้ตรงกับ data ที่ผ่านเข้ามาใน parameter

// แบบปกติ
// func sumArray[T int | float32](num []T) T {
// 	var sum T
// 	for i := range num {
// 		sum += num[i]
// 	}
// 	return sum
// }

// best practise ให้เขียนเป็น interface แยก type
type (
	Number interface {
		int | float32
	}
	// สามารถนำมาประยุกต์ใช้กับ struct และ interface ได้ด้วย
	// เช่น damage ของ playerเป็นได้ทั้ง int และ float
	Player[T Number] struct {
		Name         string
		Level        int
		AttackDamage T
	}
	Database[T Number] interface {
	}
)

func sumArray[T Number](num []T) T {
	var sum T
	for _, v := range num {
		sum += v
	}
	return sum
}

func UseGeneric() {
	intNumArray := []int{1, 2, 3}
	floatNumArray := []float32{1.5, 2.6, 3.8}

	sumNumArrayResult := sumArray(intNumArray)
	sumFloatArrayResult := sumArray(floatNumArray)

	fmt.Println(sumNumArrayResult)
	fmt.Println(sumFloatArrayResult)
}
