package structure

import "gorm.io/gorm"

type RegisterOwners struct {
	gorm.Model
	Username    string `json:"username" validate:"required"`
	Password    string `json:"password" validate:"required"`
	Firstname   string `json:"firstname" validate:"required"`
	Lastname    string `json:"lastname" validate:"required"`
	Phonenumber string `json:"phonenumber" validate:"required"`
	Lineid      string `json:"lineid" validate:"required"`
	Role        string `json:"role"`
}

type ChangePasswordOwner struct {
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required"`
}

type ChangePasswordOperator struct {
	OperatorId  int    `json:"operator_id" validate:"required"`
	NewPassword string `json:"new_password" validate:"required"`
}
