package interfaces

import "fmt"

// ประกาศ struct พร้อม interface ที่เกี่ยวข้องกับ db และ test db
// interface สำคัญมากสำหรับ unit testing เพราะถ้าเราเปลี่ยน db เพื่อใช้เทสใน function นั้น ๆ แล้ว function ก็ควรทำงานได้ปกติ
type (
	Database interface {
		Insert() error
		Update() error
	}

	PostgresDB struct{}
	MockDB     struct{}
)

// เช็คว่าเรา implement method ครบยัง เพราะถ้า return เป็น interface แต่ถ้า struct ยังไม่สามารถใช้งาน method ได้ มันจะ error
func newPostgresDB() Database {
	return &PostgresDB{}
}

func newMockDB() Database {
	return &MockDB{}
}

// implement interface
// postgres
func (p *PostgresDB) Insert() error {
	fmt.Println("Actual insert player coin")
	return nil
}

func (p *PostgresDB) Update() error {
	fmt.Println("Actual update player name")
	return nil
}

// mock test
func (m *MockDB) Insert() error {
	fmt.Println("Moct test insert player coin")
	return nil
}

func (m *MockDB) Update() error {
	fmt.Println("Mock test update player name")
	return nil
}

// using the function that pass db as an argument
func insertPlayerCoin(db Database) error {
	return db.Insert()
}
func updatePlayerName(db Database) error {
	return db.Update()
}

func UseInterface() {
	// เราสามาถใช้ฟังชั่นก์เดียวกันได้แต่พอ pass db ที่เป็น struct ที่แตกต่างกัน (แต่ใช้ interface ร่วมกัน)
	// มันก็จะ return หรือทำงานแตกต่างกันตามที่เรา implement แยกในแต่ละ interface method ของ db นั้น ๆ
	// *when we want to implement feature
	insertPlayerCoin(newPostgresDB())
	updatePlayerName(newPostgresDB())
	// *when we want to test feature
	insertPlayerCoin(newMockDB())
	updatePlayerName(newMockDB())
}
