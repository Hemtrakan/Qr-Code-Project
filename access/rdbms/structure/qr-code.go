package structure

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type TeamPage struct {
	gorm.Model
	TeamPageName string
	TeamPageFile string
	UUID         uuid.UUID `gorm:"uniqueIndex"`
	OwnersId     uint      `gorm:"foreignKey:ID"`
}

type LogTeamPage struct {
	gorm.Model
	LogTeamPageName string
	LogTeamPageFile string
	TeamPageId      uint `gorm:"foreignKey:ID"`
	OwnersId        uint `gorm:"foreignKey:ID"`
}

type Qrcode struct {
	gorm.Model
	QrCode     string
	OwnersId   uint     `gorm:"foreignKey:ID"`
	Location   Location `gorm:"foreignKey:ID"`
	TeamPageId uint     `gorm:"foreignKey:ID"`
}

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
