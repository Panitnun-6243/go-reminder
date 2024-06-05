package encapsulation

import "encoding/json"

// encapsulation คือ การห่อหุ้มข้อมูล/ซ่อนข้อมูล ข้อมูลอะไรล่ะ? ก็คุณสมบัติ (attribute) และพฤติกรรม (method) ของคลาสนั่นเอง
// โดยทำเพื่อป้องกันไม่ให้โปรแกรมเมอที่ทำงานด้วยกันหรือตัวเราเองใช้การเข้าถึงหรือแปลงค่าข้อมูลแบบผิด ๆ เช่น
// มี object อยู่หลายตัวซึ่งแต่ละตัวก็มีการติดต่อรับส่งข้อมูลกัน ถ้าไม่มีการห่อหุ้มเอาไว้ แต่ละ object เข้าถึงและเปลี่ยนแปลง
// คุณสมบัติของ object อื่นได้ทั้งหมด

// ตัวอย่างเช่น เรามีคลาส Airplane ซึ่งมี attribute คือ Name และ method คือ fly() และ fire() ถ้าไม่มีการห่อหุ้มเอาไว้
// แล้วเราสร้าง object 2 ตัว ซึ่ง object นึงเปลี่ยนแปลงค่า attribute ของ object อีกตัว ก็จะทำให้ object อีกตัวเปลี่ยนแปลงด้วย
// ซึ่งไม่ควรให้เกิดขึ้น จึงต้องมีการห่อหุ้มเอาไว้

// implement encapsulation ใน Go นั้นจะใช้ struct ในการสร้างคลาส โดย attribute จะเป็นตัวแปรที่ประกาศไว้ใน struct และ method จะเป็น
// function ที่ประกาศไว้ใน struct ด้วย โดย method จะต้องมี receiver เป็น pointer ของ struct นั้น ๆ ทุกครั้ง

// สร้าง struct ที่มี attribute และ method ที่ต้องการให้เข้าถึงได้เฉพาะ
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// implement method ที่ต้องการให้เข้าถึงได้เฉพาะ
func (u *User) GetUsername() string {
	return u.Username
}

// implement method ที่ต้องการให้เข้าถึงได้เฉพาะ
func (u *User) GetPassword() string {
	return u.Password
}

// implement method ที่ต้องการให้เข้าถึงได้เฉพาะ
func (u *User) SetUsername(username string) {
	u.Username = username
}

// implement method ที่ต้องการให้เข้าถึงได้เฉพาะ
func (u *User) SetPassword(password string) {
	u.Password = password
}

// implement method ที่ต้องการให้เข้าถึงได้เฉพาะ
func (u *User) ToJson() []byte {
	userJson, _ := json.MarshalIndent(u, "", "\t")
	return userJson
}

// implement method ที่ต้องการให้เข้าถึงได้เฉพาะ
func (u *User) FromJson(data []byte) {
	json.Unmarshal(data, u)
}

// สร้าง function ที่ใช้ในการสร้าง object และ return ออกมาเป็น struct
func NewUser(username, password string) *User {
	return &User{
		Username: username,
		Password: password,
	}
}

// สร้าง function ที่ใช้ในการเรียกใช้งาน object ที่สร้างจาก struct
func UseEncapsulation() {
	user := User{"Tanny Panich", "1234"}
	user.SetUsername("Tanny")
	user.SetPassword("1234")
	println(user.GetUsername())
	println(user.GetPassword())
	println(string(user.ToJson()))
	user.FromJson([]byte(`{"username": "Tanny", "password": "1234"}`))
	println(user.GetUsername())
	println(user.GetPassword())
}
