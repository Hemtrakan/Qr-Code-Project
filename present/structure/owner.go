package structure

import "gorm.io/gorm"

type Owners struct {
	gorm.Model
	Username    string `json:"username"`
	Password    string `json:"password"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Phonenumber string `json:"phonenumber"`
	Lineid      string `json:"lineid"`
	Role        string  `json:"role"`
}
