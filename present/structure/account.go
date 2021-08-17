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


type GetSubOwner struct {
	OwnerId          int    `json:"owner_id"`
	OwnerFirstName   string `json:"owner_first_name"`
	OwnerLastName    string `json:"owner_last_name"`
	OwnerPhoneNumber string `json:"owner_phone_number"`
	OwnerLineId      string `json:"owner_line_id"`
	UserAccountOperator []UserAccountOperator `json:"user_account_operator"`
}

type UserAccountOperator struct {
	OperatorId          int    `json:"operator_id"`
	OperatorFirstName   string `json:"operator_first_name"`
	OperatorLastName    string `json:"operator_last_name"`
	OperatorPhoneNumber string `json:"operator_phone_number"`
	OperatorLineId      string `json:"operator_line_id"`
}
