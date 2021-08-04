package structure

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateProFile struct {
	Id          int    `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	PhoneNumber string `json:"phonenumber"`
	LineId      string `json:"lineid"`
	Role        string `json:"role"`
	SubOwnerId  int    `json:"sub_owner_id"`
}

type UserAccount struct {
	Id          int    `json:"id"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	PhoneNumber string `json:"phonenumber"`
	LineId      string `json:"lineid"`
	Role        string `json:"role"`
	SubOwnerId  int    `json:"sub_owner_id"`
}

type UserAccountOwner struct {
	Id          int    `json:"id"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	PhoneNumber string `json:"phonenumber"`
	LineId      string `json:"lineid"`
	Role        string `json:"role"`
}

type UserAccountOperator struct {
	Id          int    `json:"id"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	PhoneNumber string `json:"phonenumber"`
	LineId      string `json:"lineid"`
	Role        string `json:"role"`
	SubOwnerId  int    `json:"sub_owner_id"`
}
