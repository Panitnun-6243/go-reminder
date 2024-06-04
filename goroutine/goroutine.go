package goroutine

import (
	"fmt"
	"time"
)

// goroutine -> concurrency programming เป็นการเขียนโปรแกรมในเชิง multithreading
// ไม่ใช่ synchronous แต่เป็น asynchronous (ไม่ต่อเนื่อง)
// เมื่อมีการประกาศ goroutine -> go hello() -> มันจะแยกการทำงานไปีอก thread หนึ่งละค่อยรันฟังก์ชั่นด้านล่างของ
// มันต่อ ๆ ไปแต่ถ้าการทำงานใน main จบเมื่อไหร่มันจะ force quit ออกจากโปรแกรมโดยไม่สนว่าใน ฟังชั่นก์ goroutine
// จะยังทำงานอยู่มั้ย เพราะฉะนั้นจึงมี sync.WaitGroup มาช่วยในการให้ฟังชั่นก์ที่เป็น goroutine ทำงานให้เสร็จก่อน
// จึงค่อย quit ออกจากโปรแกรม, ตัว keyword defer คือเป็นการช่วยบอกว่าก่อนที่ function จะ return ให้มาทำงานที่ defer ก่อนนะ
// func main() {
// 	thread := 3
// 	var wg sync.WaitGroup
// 	wg.Add(thread) // เป็นการกำหนดว่าจะมีการรัน goroutine และมีกี่ thread ที่ต้องรอ
// 	for i := range thread {
// 		go func(index int, wg *sync.WaitGroup) {
// 			defer wg.Done() // ใช้เพื่อบอก wg.Wait ว่ามีกี่ thread ที่ืทำงานเสร็จแล้ว
// 			fmt.Println(index)
// 		}(i, &wg) // ทำการ loop ทีละ thread แต่พอเป็น goroutine function จึงเหมือนแยกร่าง ไม่ต่อกัน
// เช่น ถ้า loop ปกติจะเป็น index 0,1,2 แต่เมื่อแยกร่างอาจจะเป็น 0,2,1 ก็ได้
// 	}
// 	wg.Wait() // รอจนว่า wg.Add จะเป็น 0 จึงจะหยุดทำงาน แล้วก็ return parent function ออกไปได้อย่างสมบูรณ์
// }

// ลองใช้ดู
func UseGouroutine() {
	fmt.Println("Print on Thread 1")
	go helloGoroutine()
	fmt.Println("Print on Thread 2")
	// หน่วงเวลาเพื่อให้ได้เห็น goroutine ทำงานชัด ๆ เพราะเดี๋ยว "Print on Thread 2" มันทำงานแปปเดียวละจบเลย
	time.Sleep(1 * time.Second)
}

func helloGoroutine() {
	for i := range 10 {
		fmt.Println("Goroutine loop number", i)
		fmt.Println("Hello from another thread that are not thread 1 and 2")
	}

}
