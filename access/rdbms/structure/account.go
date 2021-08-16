package structure

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Username    string `gorm:"uniqueIndex"`
	Password    string
	FirstName   string
	LastName    string
	PhoneNumber string `gorm:"uniqueIndex"`
	LineId      string `gorm:"uniqueIndex"`
	Role        string
	SubOwnerId  uint
}