package structure

import (
	"github.com/gofrs/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

//type Template struct {
//	gorm.Model
//	TeamPageName string
//	TeamPageFile string
//}

type QrCode struct {
	gorm.Model
	OwnerId      uint
	TemplateName string
	Info         datatypes.JSON
	Ops          datatypes.JSON
	HistoryInfo  datatypes.JSON
	QrCodeUUID   uuid.UUID `gorm:"uniqueIndex"`
	Code         string    `gorm:"uniqueIndex"`
}
//type LogTeamPage struct {
//	gorm.Model
//	LogTeamPageName string
//	LogTeamPageFile string
//	TeamPageId      uint `gorm:"foreignKey:ID"`
//	OwnersId        uint `gorm:"foreignKey:ID"`
//}

//type Qrcode struct {
//	gorm.Model
//	QrCode     string
//	OwnersId   uint     `gorm:"foreignKey:ID"`
//	Location   Location `gorm:"foreignKey:ID"`
//	TeamPageId uint     `gorm:"foreignKey:ID"`
//}

type Location struct {
	gorm.Model
	Country     string
	Address     string
	SubDistrict string
	District    string
	Province    string
	Zipcode     string
	XCoordinate string
	YCoordinate string
}
