package structure

import (
	"github.com/gofrs/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type QrCode struct {
	gorm.Model
	// OwnerId ไอดีของเจ้าของ QrCode
	OwnerId uint
	// TemplateName ระบุแม่แบบของ QrCode
	TemplateName string
	// Info ข้อมูลของ QrCode นั้นๆ
	Info datatypes.JSON
	//  QrCodeUUID QrCodeId ที่เป็น uuid
	QrCodeUUID uuid.UUID `gorm:"uniqueIndex"`
	//Code ตั้งชื่อ QrCodeName
	Code string
	//Count จำนวณของ TemplateName ที่ถูกสร้างขึ้นเจ้า Owner นั้นๆ
	Count string
	//First   สถานะ การสร้างโดยยังไม่เพิ่มข้อมูล ยังไม่เพิ่มเป็น false เพิ่มข้อมูลแล้ว เป็น true
	First bool
	//Active  สถานะ การเปิดปิดการใช้งาน เปิด คือ true ปิด คือ false
	Active      bool
	DataHistory []HistoryInfo `gorm:"foreignKey:QrCodeRefer"`
	DataOps     []Ops         `gorm:"foreignKey:QrCodeRefer"`
}

type HistoryInfo struct {
	gorm.Model
	// QrCodeID QrCodeId ที่เป็น uuid
	QrCodeID uuid.UUID
	// HistoryInfo ข้อมูลในการอัพเดท
	HistoryInfo datatypes.JSON
	//UserId ใครเป็นคนอัพเดท
	UserId      uint
	QrCodeRefer uint
}

type Ops struct {
	gorm.Model
	// QrCodeID QrCodeId ที่เป็น uuid
	QrCodeID uuid.UUID
	// Operator การกระทำต่อสิ่งนั้นๆ
	Operator datatypes.JSON
	//UserId ใครเป็นคนอัพเดท
	UserId      uint
	QrCodeRefer uint
}

//type TestQrCode struct {
//	gorm.Model
//	QrCodeName  string
//	QrCodeId    uint
//	DataHistory []TestHistory `gorm:"foreignKey:QrCodeRefer"`
//	DataOps     []TestOps     `gorm:"foreignKey:QrCodeRefer"`
//}
//
//type TestHistory struct {
//	gorm.Model
//	HistoryData string
//	QrCodeRefer uint
//}
//
//type TestOps struct {
//	gorm.Model
//	OpsData     string
//	QrCodeRefer uint
//}
