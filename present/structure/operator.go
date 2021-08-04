package structure

import "gorm.io/gorm"

type RegisterOperator struct {
	gorm.Model
	Username    string `json:"username"`
	Password    string `json:"password"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Phonenumber string `json:"phonenumber"`
	Lineid      string `json:"lineid"`
	Role        string `json:"role"`
	SubOwnerId  uint   `json:"sub_owner_id"`
}
