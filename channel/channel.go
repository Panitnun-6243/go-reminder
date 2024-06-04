package channel

import "fmt"

// channel เหมือนเป็น pipe หรือท่อที่เอาไว้เก็บและขนส่ง state (ข้อมูล) ระหว่าง goroutine หลาย ๆ
// ตัว แล้วเอาข้อมูลนั้นไปเก็บไว้ในตัวแปรซักอัน เพื่อนำไปใช้งานต่อ
// *อย่าลืม close(channel) เพื่อให้มันไม่สามารถรับค่าได้และป้องกันการเกิด deadlock
// ลูกศรชี้ออกจาก channel func double(channelVariableName <-chan int) หมายความว่า channel สามารถส่งได้อย่างเดียวในขณะนั้น
// ลูกศรชี้เข้า channel func double(channelVariableName chan<- int) หมายความว่า channel สามารถรับได้อย่างเดียวในขณะนั้น
// *buffer size ใน channel ถ้าใช้รับหรือส่งไม่ครบจะเกิดการ wait forever ต้องเช็คดี ๆ
func basicPipe() {
	fmt.Println("I'm thread 1")
	// สร้าง channel , 10 คือ buffer size หรือก็คือ channel กี่ตัว
	messagePipe := make(chan string, 10)
	// run goroutine เพื่อที่จะส่งข้อความ
	go func() {
		// ส่งข้อมูล string เข้าไปใน channel จาก goroutine ที่ 2 (โปรแกรมปกติของเราตอนนี้อยู่ thread ที่ 1)
		messagePipe <- "Hello, I'm from another thread send through channel"
	}()
	// thread ที่ 1 รับข้อความจาก goroutine thread ที่ 2 ผ่าน channel
	msg := <-messagePipe
	fmt.Println(msg)
}

// ฟังชั่นก์ที่ใช้ job channel ส่ง 1-10 ไปคูณ 2 แล้วส่งไปให้ result channel รับ
func doubleAndPassToResultChannel(jobCh <-chan int, resultCh chan<- int) {
	for eachJob := range jobCh {
		resultCh <- eachJob * 2
	}
}

func UseChannel() {
	basicPipe()

	// สร้างท่อส่่งข้อมูล 10 ท่อ
	jobChannel := make(chan int, 10)
	// สร้างท่อรับข้อมูฃ 10 ท่อ
	resultChannel := make(chan int, 10)

	// ส่งค่า 1-10 เข้าไปเก็บไว้ใน job channel
	for i := range 10 {
		jobChannel <- i + 1 // จะได้เริ่มที่ 1 ไม่ใช่เริ่ม 0

	}
	close(jobChannel)
	go doubleAndPassToResultChannel(jobChannel, resultChannel)

	for range 10 {
		result := <-resultChannel
		fmt.Println(result)
	}
}
