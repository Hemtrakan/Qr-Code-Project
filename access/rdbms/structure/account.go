package structure

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Username    string
	Password    string
	FirstName   string
	LastName    string
	PhoneNumber string
	LineId      string
	LineUserId  *string
	Role        string
	SubOwnerId  *uint
	OpsAccount  []Account `gorm:"foreignKey:SubOwnerId"`
}
//
//type DataInAccount struct {
//	gorm.Model
//	Username    string
//	Password    string
//	FirstName   string
//	LastName    string
//	PhoneNumber string
//	LineId      string
//	LineUserId  *string
//	Role        string
//	SubOwnerId  *uint
//}
//
//type Account struct {
//	DataInAccount
//	OpsAccount []DataInAccount `gorm:"foreignKey:SubOwnerId"`
//}
