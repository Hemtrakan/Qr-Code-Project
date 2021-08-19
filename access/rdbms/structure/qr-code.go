package structure

import (
	"github.com/gofrs/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type QrCode struct {
	gorm.Model
	OwnerId      uint
	TemplateName string
	Info         datatypes.JSON
	Ops          datatypes.JSON
	HistoryInfo  datatypes.JSON
	QrCodeUUID   uuid.UUID `gorm:"uniqueIndex"`
	Code         string
	Count        string
}
