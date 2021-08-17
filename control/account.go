package control

import (
	"errors"
	"gorm.io/gorm"
	"qrcode/access/constant"
	rdbmsstructure "qrcode/access/rdbms/structure"
	"qrcode/present/structure"
	"qrcode/utility"
	"strconv"
	"time"
)

func (ctrl *APIControl) RegisterOwner(reqOwner *structure.RegisterOwners) (Error error) {

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
		FirstName:   "FirstName",
		LastName:    "LastName",
		PhoneNumber: "PhoneNumber",
		LineId:      "LineId",
		Role:        string(constant.Admin),
	}
	err = ctrl.insert(admin)
	if err != nil {
		Error = err
		return
	}
	return
}

func (ctrl *APIControl) RegisterOperator(reqOperator *structure.RegisterOperator) (Error error) {
	OwnerId := int(reqOperator.SubOwnerId)
	data, err := ctrl.access.RDBMS.GetAccount(OwnerId)
	if data.ID == 0 {
		Error = errors.New("there is no owner of this id in the system.")
		return
	}
	if data.Role != string(constant.Owner) {
		Error = errors.New("invalid user rights")
		return
	}

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
		Error = errors.New("incorrect password.")
		return
	}
	Token, err = utility.AuthenticationLogin(data.ID, data.Role)
	if err != nil {
		Error = err
		return
	}
	return Token, nil
}

func (ctrl *APIControl) LoginAdmin(reqLogin *structure.Login) (Token string, Error error) {
	login := rdbmsstructure.Account{
		Username: reqLogin.Username,
		Password: reqLogin.Password,
	}
	data, err := ctrl.access.RDBMS.Login(login)
	if err != nil {
		Error = err
		return
	}
	if data.Role == string(constant.Admin) {
		err = utility.VerifyPassword(data.Password, login.Password)
		if err != nil {
			Error = errors.New("incorrect password.")
			return
		}
		Token, err = utility.AuthenticationLogin(data.ID, data.Role)
		if err != nil {
			Error = err
			return
		}
	} else {
		Error = errors.New("Your user rights are not reached.")
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
		SubOwnerId:  int(data.SubOwnerId),
	}
	return
}

func (ctrl *APIControl) GetAllAccountOwner() (response []structure.UserAccountOwner, Error error) {
	var DataArray []structure.UserAccountOwner
	res, err := ctrl.access.RDBMS.GetAllAccountOwner()
	if err != nil {
		Error = err
		return
	}
	if len(res) == 0 {
		Error = errors.New("record not found")
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
		DataArray = append(DataArray, UserAccountStructure)
	}
	response = DataArray
	return
}

func (ctrl *APIControl) GetSubOwner(OwnerId int) (response structure.GetSubOwner, Error error) {
	var DataArray []structure.UserAccountOperator
	ops, err := ctrl.access.RDBMS.GetSubOwner(OwnerId)
	if err != nil {
		Error = err
		return
	}
	owner, err := ctrl.access.RDBMS.GetAccount(OwnerId)
	if err != nil {
		Error = err
		return
	}
	ownerId := int(owner.ID)
	for _, data := range ops {
		id := int(data.ID)
		UserAccountOperator := structure.UserAccountOperator{
			OperatorId:          id,
			OperatorFirstName:   data.FirstName,
			OperatorLastName:    data.LastName,
			OperatorPhoneNumber: data.PhoneNumber,
			OperatorLineId:      data.LineId,
		}
		DataArray = append(DataArray, UserAccountOperator)
	}
	var UserAccountStructure = structure.GetSubOwner{
		OwnerId:             ownerId,
		OwnerFirstName:      owner.FirstName,
		OwnerLastName:       owner.LastName,
		OwnerPhoneNumber:    owner.PhoneNumber,
		OwnerLineId:         owner.LineId,
		UserAccountOperator: DataArray,
	}
	response = UserAccountStructure
	return
}

func (ctrl *APIControl) GetAllAccountOperator() (response []structure.UserAccountOperator, Error error) {
	var DataArray []structure.UserAccountOperator
	res, err := ctrl.access.RDBMS.GetAllAccountOperator()
	if err != nil {
		Error = err
		return
	}
	if len(res) == 0 {
		Error = errors.New("record not found")
		return
	}
	for _, data := range res {
		id := int(data.ID)
		UserAccountStructure := structure.UserAccountOperator{
			OperatorId:          id,
			OperatorFirstName:   data.FirstName,
			OperatorLastName:    data.LastName,
			OperatorPhoneNumber: data.PhoneNumber,
			OperatorLineId:      data.LineId,
		}
		DataArray = append(DataArray, UserAccountStructure)
	}
	response = DataArray
	return
}

func (ctrl *APIControl) GetOwnerByIdOps(OperatorId int) (response structure.GetOwnerByOperator, Error error) {
	ops, err := ctrl.access.RDBMS.GetOwnerByIdOps(OperatorId)
	opsId := int(ops.SubOwnerId)
	owner, err := ctrl.access.RDBMS.GetAccount(opsId)
	if err != nil {
		Error = err
		return
	}
	OwnerId := int(owner.ID)
	response = structure.GetOwnerByOperator{
		Operator: structure.Operator{
			FirstName:   ops.FirstName,
			LastName:    ops.LastName,
			PhoneNumber: ops.PhoneNumber,
			LineId:      ops.LineId,
			Owner: structure.Owner{
				OwnerId:     OwnerId,
				FirstName:   owner.FirstName,
				LastName:    owner.LastName,
				PhoneNumber: owner.PhoneNumber,
				LineId:      owner.LineId,
			},
		},
	}
	return
}

func (ctrl *APIControl) UpdateProfile(id uint, Account *structure.UpdateProFile) (Error error) {
	data := rdbmsstructure.Account{
		Model: gorm.Model{
			ID:        id,
			UpdatedAt: time.Now(),
		},
		FirstName:   Account.FirstName,
		LastName:    Account.LastName,
		PhoneNumber: Account.PhoneNumber,
		LineId:      Account.LineId,
	}

	sid := strconv.FormatUint(uint64(id), 16)
	userId, err := strconv.Atoi(sid)
	res, err := ctrl.access.RDBMS.GetAccount(userId)
	if res.ID == 0 {
		Error = errors.New("record not found")
		return
	}
	err = ctrl.access.RDBMS.UpdateProfile(data)
	if err != nil {
		Error = err
		return
	}
	return
}

func (ctrl *APIControl) ChangePassword(id uint, password *structure.ChangePassword) (Error error) {
	hashPassword, err := utility.Hash(password.Password)
	if err != nil {
		return err
	}
	data := rdbmsstructure.Account{
		Model: gorm.Model{
			ID:        id,
			UpdatedAt: time.Now(),
		},
		Password: string(hashPassword),
	}

	//sid := strconv.FormatUint(uint64(id), 16)
	//userId, err := strconv.Atoi(sid)
	//res, err := ctrl.access.RDBMS.GetAccount(userId)
	//if res.ID == 0 {
	//	Error = errors.New("record not found")
	//	return
	//}
	err = ctrl.access.RDBMS.UpdateProfile(data)
	if err != nil {
		Error = err
		return
	}
	return
}

func (ctrl *APIControl) DeleteAccount(id int) (Error error) {
	sid := strconv.FormatUint(uint64(id), 16)
	userId, err := strconv.Atoi(sid)
	res, err := ctrl.access.RDBMS.GetAccount(userId)
	if res.ID == 0 {
		Error = errors.New("record not found")
		return
	}
	err = ctrl.access.RDBMS.DeleteAccount(id)
	if err != nil {
		Error = err
		return
	}
	return
}

func (ctrl *APIControl) insert(Account rdbmsstructure.Account) error {
	err := ctrl.access.RDBMS.Register(Account)
	if err != nil {
		return err
	}
	return nil
}
