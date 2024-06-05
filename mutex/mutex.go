package mutex

import (
	"fmt"
	"sync"
	"time"
)

// mutex นำมาแก้ปัญหา racing codition ในระดับการ deployment ซึ่งไม่เคยเกิดใน local
// racing condition คือปรากฎการณ์ที่มี request เข้ามาเยอะ ๆ พร้อมแล้ว server แบ่งการ
// ทำงานเป็น multithread พร้อมกันทำให้ไม่เป็นลำดับเช่น transaction ในกระเป๋าเงินออนไลน์
// server1 -> balance - 100 amount, 2 -> balance - 50 amount, 3 -> balance - 20 amount
// แต่การเป็นว่า server 2,3 ทำงานไวกว่า 1 ทำให้เกิดข้อมูลทับซ้อนทั้ง ๆ ที่บันทึกข้อมูลใน database ไปแล้ว
// แต่ process ของ server 1 ทำงานได้ช้ากว่า เช่น มีเงิน 1000,
// 2 -> 1000 - 50 = 950
// 3 -> 950 - 20 = 930
// 1 -> 1000 - 100 = 900
// กลายเป็นว่าข้อมูลสุดท้ายเก็บลงใน db จะเท่ากับ 900 แทนที่จะเป็น 830 ร้ายแรงอยู่นะ
// จึงต้องมีตัว mutex มาเป็นกลไกการล็อคเพื่อให้แน่ใจว่าเมื่อมี Goroutine หลายๆตัวเข้าถึงข้อมูลพร้อมๆกัน
// จะมีเพียงแค่ Goroutine เดียวเท่านั้นที่เข้าถึงข้อมูล ณ ช่วงเวลานั้น
var (
	mu      sync.Mutex
	balance int = 1000
)

// มี 3 request/data (transaction) เข้ามาพร้อมกัน
func UseMutex() {
	doneChannel := make(chan bool, 3)
	go updateBalace(doneChannel, 100) // ตัวแทน server1
	go updateBalace(doneChannel, 50)  // ตัวแทน server2
	go updateBalace(doneChannel, 20)  // ตัวแทน server3

	<-doneChannel
	<-doneChannel
	<-doneChannel

	fmt.Println(balance)
}

func updateBalace(doneCh chan bool, amount int) {
	mu.Lock() // function บน thread ไหนใช้งานก่อนไม่รู้แต่ล้อคทรัพยากรเลย พอเสร็จแล้ว
	// ค่อย unlock แล้วให้ function บน thread ตัวต่อไปใช้ต่อเป็นลำดับขั้น (synchronous)

	fmt.Println("Updating balance")
	time.Sleep(1 * time.Second)
	fmt.Println("Now, the current transaction is", balance, "-", amount)
	balance -= amount

	doneCh <- true

	mu.Unlock()
	fmt.Println("Balance is updated")

}
