package structure

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ChangePassword struct {
	Password string `json:"password"`
}

type UpdateProFile struct {
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	PhoneNumber string `json:"phonenumber"`
	LineId      string `json:"lineid"`
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

type GetOwnerByOperator struct {
	Operator Operator `json:"operator"`
}

type Operator struct {
	FirstName   string `json:"operator_firstname"`
	LastName    string `json:"operator_lastname"`
	PhoneNumber string `json:"operator_phonenumber"`
	LineId      string `json:"operator_lineid"`
	Owner       Owner  `json:"owner"`
}

type Owner struct {
	OwnerId     int    `json:"owner_id"`
	FirstName   string `json:"owner_firstname"`
	LastName    string `json:"owner_lastname"`
	PhoneNumber string `json:"owner_phonenumber"`
	LineId      string `json:"owner_lineid"`
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
