package abstraction

import "encoding/json"

// ต่อเนื่องมาจาก encapsulation หรือก็คือการให้เข้าถึง/รู้แค่ข้อมูลที่จำเป็น เช่น เราไปเดทกับสาว แต่เราก็แค่บอกว่าใส่ชุดอะไร
// เสื้อสีอะไร กางเกงอะไรเพื่อที่จะได้หากันเจอง่าย ๆ ที่จุดนัดพบ แต่จริง ๆ แล้วเรายังมีข้อมูลหรือพฤติกรรมอีกมากมายที่เขายังไม่รู้
// และเรายังไม่จำเป็นต้องบอก ซึ่งเราและสาวก็คือ object 2 ตัวที่ติดต่อสื่อสารกัน (object นึงเรียกใช้งานอีก object นึง)

// implement abstraction ใน Go นั้นจะใช้ interface ในการทำ โดย interface จะเป็นตัวกลางในการเชื่อมต่อระหว่าง object 2 ตัว
// โดย object ที่ต้องการเชื่อมต่อกันจะต้องมี method ที่ตรงกับ interface นั้น ๆ ทั้งหมด ถึงจะสามารถเชื่อมต่อกันได้

// สร้าง interface ที่มี method ที่ต้องมีใน object ที่ต้องการเชื่อมต่อกัน
type User interface {
	Name() string
}

// สร้าง struct ที่มี attribute และ method ที่ต้องมีใน interface
type user struct {
	Username string `json:"name"`
}

// implement method ที่ต้องมีใน interface
func (u *user) Name() string {
	return "Mr. " + u.Username
}

// สร้าง function ที่ใช้ในการสร้าง object และ return ออกมาเป็น interface
func ParseUserData(data []byte) (User, error) {
	user := &user{}
	if err := json.Unmarshal(data, user); err != nil {
		return nil, err
	}
	return user, nil
}

// สร้าง function ที่ใช้ในการเรียกใช้งาน object ที่สร้างจาก interface
func UseAbstraction() {
	data := []byte(`{"name": "Tanny"}`)
	user, err := ParseUserData(data)
	if err != nil {
		panic(err)
	}
	// สามารถเรียกใช้ method จาก interface ได้เลย
	println(user.Name())
}
