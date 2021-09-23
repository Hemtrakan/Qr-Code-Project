package structure

import (
	"qrcode/utility"
	"time"
)

type Login struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginOwner struct {
	Username string  `json:"username" validate:"required"`
	Password string  `json:"password" validate:"required"`
	UID      string `json:"uid"`
}

type LoginOperator struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	UID      *string `json:"uid" validate:"required"`
}

type ChangePassword struct {
	Password string `json:"password" validate:"required"`
}

type UpdateProFile struct {
	FirstName   string `json:"firstname" validate:"required"`
	LastName    string `json:"lastname" validate:"required"`
	PhoneNumber string `json:"phonenumber" validate:"required"`
	LineId      string `json:"lineid" validate:"required"`
}

type UserAccount struct {
	Id          int    `json:"id"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	PhoneNumber string `json:"phonenumber"`
	LineId      string `json:"lineid"`
	Role        string `json:"role"`
	SubOwnerId  *uint  `json:"sub_owner_id"`
}

type GetOwnerByOperator struct {
	Operator Operator `json:"operator"`
}

type Operator struct {
	Id          uint      `json:"operator_id"`
	UserName    string    `json:"operator_username"`
	FirstName   string    `json:"operator_firstname"`
	LastName    string    `json:"operator_lastname"`
	PhoneNumber string    `json:"operator_phonenumber"`
	LineId      string    `json:"operator_lineid"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Owner       Owner     `json:"owner"`
}

type Owner struct {
	OwnerId     uint      `json:"owner_id"`
	FirstName   string    `json:"owner_firstname"`
	LastName    string    `json:"owner_lastname"`
	PhoneNumber string    `json:"owner_phonenumber"`
	LineId      string    `json:"owner_lineid"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserAccountOwnerWithPaginationResponse struct {
	*utility.Paginator
	Detail []UserAccountOwner `json:"detail"`
}

type UserAccountOperatorWithPaginationResponse struct {
	*utility.Paginator
	Detail []UserAccountOperator `json:"detail"`
}

type UserAccountOwner struct {
	Id          int       `json:"id"`
	UserName    string    `json:"user_name"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	PhoneNumber string    `json:"phone_number"`
	LineId      string    `json:"line_id"`
	Role        string    `json:"role"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GetSubOwner struct {
	OwnerId             uint        `json:"owner_id"`
	OwnerUserName       string      `json:"owner_user_name"`
	OwnerFirstName      string      `json:"owner_first_name"`
	OwnerLastName       string      `json:"owner_last_name"`
	OwnerPhoneNumber    string      `json:"owner_phone_number"`
	OwnerLineId         string      `json:"owner_line_id"`
	UserAccountOperator []Operators `json:"user_account_operator"`
}

type Operators struct {
	OperatorId          uint      `json:"operator_id"`
	OperatorUserName    string    `json:"operator_user_name"`
	OperatorFirstName   string    `json:"operator_first_name"`
	OperatorLastName    string    `json:"operator_last_name"`
	OperatorPhoneNumber string    `json:"operator_phone_number"`
	OperatorLineId      *string    `json:"operator_line_id"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

type OperatorsLine struct {
	OperatorId          uint      `json:"operator_id"`
	OperatorUserName    string    `json:"operator_user_name"`
	OperatorFirstName   string    `json:"operator_first_name"`
	OperatorLastName    string    `json:"operator_last_name"`
	OperatorLineId      *string    `json:"operator_line_id"`
}

type UserAccountOperator struct {
	OperatorId          int       `json:"operator_id"`
	OperatorUserName    string    `json:"operator_user_name"`
	OperatorFirstName   string    `json:"operator_first_name"`
	OperatorLastName    string    `json:"operator_last_name"`
	OperatorPhoneNumber string    `json:"operator_phone_number"`
	OperatorLineId      string    `json:"operator_line_id"`
	OwnerId             uint      `json:"owner_id"`
	OwnerName           string    `json:"owner_name"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

type SearchAccountOwner struct {
	Page        *int    `json:"page"   query:"page"`
	Limit       *int    `json:"limit"  query:"limit"`
	Firstname   *string `json:"firstname" query:"firstname"`
	Lastname    *string `json:"lastname" query:"lastname"`
	Phonenumber *string `json:"phonenumber" query:"phonenumber"`
	Lineid      *string `json:"lineid" query:"lineid"`
}

type SearchAccountOperator struct {
	Page        *int    `json:"page"   query:"page"`
	Limit       *int    `json:"limit"  query:"limit"`
	Name        *string `json:"name"   query:"name"`
	Firstname   *string `json:"firstname" query:"firstname"`
	Lastname    *string `json:"lastname" query:"lastname"`
	Phonenumber *string `json:"phonenumber" query:"phonenumber"`
	Lineid      *string `json:"lineid" query:"lineid"`
}
