package structure

import (
	"github.com/gofrs/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type QrCode struct {
	gorm.Model
	OwnerId      uint
	TemplateName string
	Info         datatypes.JSON
	Ops          datatypes.JSON
	QrCodeUUID   uuid.UUID `gorm:"uniqueIndex"`
	Code         string
	Count        string
	First        bool // todo สถานะ การสร้างโดยยังไม่เพิ่มข้อมูล ยังไม่เพิ่มเป็น false เพิ่มข้อมูลแล้ว เป็น true
	Active       bool // todo สถานะ การเปิดปิดการใช้งาน เปิด คือ true ปิด คือ false
}

type History struct {
	QrCodeUUID  uuid.UUID
	HistoryInfo datatypes.JSON
	UserId      uint
	UpdatedAt   time.Time
}
