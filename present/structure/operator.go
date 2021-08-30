package structure

import "gorm.io/gorm"

type RegisterOperator struct {
	gorm.Model
	Username    string `json:"username" validate:"required"`
	Password    string `json:"password" validate:"required"`
	Firstname   string `json:"firstname" validate:"required"`
	Lastname    string `json:"lastname" validate:"required"`
	Phonenumber string `json:"phonenumber" validate:"required"`
	Lineid      string `json:"lineid" validate:"required"`
	Role        string `json:"role"`
	SubOwnerId  *uint   `json:"sub_owner_id"`
}