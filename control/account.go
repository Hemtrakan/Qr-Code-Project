package control

import (
	"qrcode/access/constant"
	rdbmsstructure "qrcode/access/rdbms/structure"
	"qrcode/present/structure"
	"qrcode/utility"
)

func (ctrl *APIControl) RegisterOwner(reqOwner *structure.Owners) (Error error) {

	hashPassword, err := utility.Hash(reqOwner.Password)
	if err != nil {
		return err
	}

	Owner := rdbmsstructure.Account{
		Username:    reqOwner.Username,
		Password:    string(hashPassword),
		FirstName:   reqOwner.Firstname,
		LastName:    reqOwner.Lastname,
		PhoneNumber: reqOwner.Phonenumber,
		LineId:      reqOwner.Lineid,
		Role:        string(constant.Owner),
	}
	err = ctrl.insert(Owner)
	if err != nil {
		Error = err
		return
	}
	return
}

func (ctrl *APIControl) RegisterAdmin() (Error error) {
	hashPassword, err := utility.Hash("1234")
	if err != nil {
		return err
	}

	admin := rdbmsstructure.Account{
		Username:    "admin",
		Password:    string(hashPassword),
		FirstName:   "admin",
		LastName:    "T-dev",
		PhoneNumber: "-",
		LineId:      "-",
		Role:        string(constant.Admin),
	}
	err = ctrl.insert(admin)
	if err != nil {
		Error = err
		return
	}
	return
}


func (ctrl *APIControl) RegisterOperator(reqOperator *structure.Operator) (Error error) {
	hashPassword, err := utility.Hash(reqOperator.Password)
	if err != nil {
		return err
	}

	Operator := rdbmsstructure.Account{
		Username:    reqOperator.Username,
		Password:    string(hashPassword),
		FirstName:   reqOperator.Firstname,
		LastName:    reqOperator.Lastname,
		PhoneNumber: reqOperator.Phonenumber,
		LineId:      reqOperator.Lineid,
		Role:        string(constant.Operator),
		SubOwnerId:  reqOperator.SubOwnerId,
	}
	err = ctrl.insert(Operator)
	if err != nil {
		Error = err
		return
	}
	return
}

func (ctrl *APIControl) Login(reqLogin *structure.Login) (Token string, Error error) {
	login := rdbmsstructure.Account{
		Username: reqLogin.Username,
		Password: reqLogin.Password,
	}
	data, err := ctrl.access.RDBMS.Login(login)
	if err != nil {
		Error = err
		return
	}
	err = utility.VerifyPassword(data.Password, login.Password)
	if err != nil {
		Error = err
		return
	}
	Token, err = utility.AuthenticationLogin(data.ID, data.Role)
	if err != nil {
		Error = err
		return
	}
	return Token, nil
}

func (ctrl *APIControl) GetAccount(id int) (response structure.UserAccount, Error error) {
	data, err := ctrl.access.RDBMS.GetAccount(id)
	if err != nil {
		Error = err
		return
	}
	id = int(data.ID)
	response = structure.UserAccount{
		Id:          id,
		FirstName:   data.FirstName,
		LastName:    data.LastName,
		PhoneNumber: data.PhoneNumber,
		LineId:      data.LineId,
		Role:        data.Role,
		SubOwnerId: int(data.SubOwnerId),
	}
	return
}

func (ctrl *APIControl) GetAllAccountOwner() (response []structure.UserAccountOwner ,Error error) {
	var DataArray []structure.UserAccountOwner
	res , err := ctrl.access.RDBMS.GetAllAccountOwner()
	if err != nil{
		Error = err
		return
	}
	for _, data := range res {
		id := int(data.ID)
		UserAccountStructure := structure.UserAccountOwner{
			Id:          id,
			FirstName:   data.FirstName,
			LastName:    data.LastName,
			PhoneNumber: data.PhoneNumber,
			LineId:      data.LineId,
			Role:        data.Role,
		}
		DataArray = append(DataArray,UserAccountStructure)
	}
	response = DataArray
	return
}

func (ctrl *APIControl) GetAllAccountOperator() (response []structure.UserAccountOperator ,Error error) {
	var DataArray []structure.UserAccountOperator
	res , err := ctrl.access.RDBMS.GetAllAccountOperator()
	if err != nil{
		Error = err
		return
	}
	for _, data := range res {
		id := int(data.ID)
		UserAccountStructure := structure.UserAccountOperator{
			Id:          id,
			FirstName:   data.FirstName,
			LastName:    data.LastName,
			PhoneNumber: data.PhoneNumber,
			LineId:      data.LineId,
			Role:        data.Role,
			SubOwnerId: int(data.SubOwnerId),
		}
		DataArray = append(DataArray,UserAccountStructure)
	}
	response = DataArray
	return
}

func (ctrl *APIControl) insert(Account rdbmsstructure.Account) error {
	err := ctrl.access.RDBMS.Register(Account)
	if err != nil {
		return err
	}
	return nil
}
