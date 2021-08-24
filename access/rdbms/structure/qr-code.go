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
	First        bool
}

type History struct {
	QrCodeUUID  uuid.UUID
	HistoryInfo datatypes.JSON
	UserId      uint
	UpdatedAt   time.Time
}
