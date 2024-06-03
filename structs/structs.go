package structs

import (
	"encoding/json"
	"fmt"
)

// struct -> แบบแปลนในการสร้าง custom type
// ประกาศ struct ของ player -> username, level, status, class
type player struct {
	Username string `json:"username"`
	Level    uint   `json:"level"`
	Status   string `json:"status"`
	Class    string `json:"class"`
}

func UseStruct() {
	player1 := player{
		Username: "Tanny",
		Level:    100,
		Status:   "Active",
		Class:    "Warrior",
	}
	fmt.Println(player1)
	// ใช้ marshal เพื่อบีบอัด struct (object) ให้เป็น json (byte) แต่ต้อง pass by interence เพราะมัน copy ค่าไป (รับเป็น interface any)
	jsonPlayer, _ := json.MarshalIndent(&player1, "", "\t")
	// อย่าลืมแปลง byte เป้น string
	fmt.Println(string(jsonPlayer))
	// ใช้ unmarshal แปลงกลับเป็น instance (obeject)
	var objectPlayer player
	json.Unmarshal(jsonPlayer, &objectPlayer)
	fmt.Println(objectPlayer)
	// ใช้ method จาก struct
	fmt.Println(player1.getPlayerUsername())
	fmt.Println("Before level up", player1.Level)
	player1.upLevel()
	fmt.Println("After level up", player1.Level)
}

// การเอา struct ไปใช้อารมณ์เหมือนสร้าง method ใน class
func (p player) getPlayerUsername() string {
	return p.Username
}

// method uplevel
func (p *player) upLevel() {
	p.Level += 1
}
