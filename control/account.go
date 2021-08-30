package control

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"qrcode/access/constant"
	rdbmsstructure "qrcode/access/rdbms/structure"
	"qrcode/present/structure"
	"qrcode/utility"
	regexp "regexp"
	"strings"
	"time"
	"unicode"
)

func validPassword(s string) error {
next:
	for name, classes := range map[string][]*unicode.RangeTable{
		"upper case": {unicode.Upper, unicode.Title},
		"lower case": {unicode.Lower},
		"numeric":    {unicode.Number, unicode.Digit},
		//"special":    {unicode.Space, unicode.Symbol, unicode.Punct, unicode.Mark},
	} {
		for _, r := range s {
			if unicode.IsOneOf(classes, r) {
				continue next
			}
		}
		return fmt.Errorf("password must have at least one %s character", name)
	}
	return nil
}

func (ctrl *APIControl) RegisterOwner(reqOwner *structure.RegisterOwners) (Error error) {
	reqOwner.Username = strings.ToLower(reqOwner.Username)
	reqOwner.Password = strings.Trim(reqOwner.Password, "\t \n")
	reqOwner.Firstname = strings.Trim(reqOwner.Firstname, "\t \n")
	reqOwner.Lastname = strings.Trim(reqOwner.Lastname, "\t \n")
	reqOwner.Lineid = strings.Trim(reqOwner.Lineid, "\t \n")
	reqOwner.Phonenumber = strings.Trim(reqOwner.Phonenumber, "\t \n")
	user, err := regexp.MatchString("^[a-z0-9_-]{6,20}$", reqOwner.Username)
	if !user {
		return errors.New("username ต้องไม่ต่ำกว่า 6 ตัว และ ไม่เกิน 20 ตัว และมีอักษรพิเศษได้แค่ _- เท่านั้น")
	}
	pass, err := regexp.MatchString("^[a-zA-Z0-9_!-]{6,20}$", reqOwner.Password)
	if !pass {
		return errors.New("password ต้องไม่ต่ำกว่า 6 ตัว และ ไม่เกิน 20 ตัว ต้องมีตัวพิมพ์ใหญ่และตัวพิมพ์เล็กตัวเลขและมีอักษรพิเศษได้แค่ !_- เท่านั้น")
	}
	err = validPassword(reqOwner.Password)
	if err != nil {
		return errors.New("password ต้องไม่ต่ำกว่า 6 ตัว และ ไม่เกิน 20 ตัว ต้องมีตัวพิมพ์ใหญ่และตัวพิมพ์เล็กตัวเลขและมีอักษรพิเศษได้แค่ !_- เท่านั้น")
	}
	if !(len(reqOwner.Firstname) <= 30) {
		return errors.New("firstname ต้องไม่เกิน 30 ตัว")
	}
	if !(len(reqOwner.Lastname) <= 30) {
		return errors.New("lastname ต้องไม่เกิน 30 ตัว")
	}
	if reqOwner.Firstname == "" {
		return errors.New("firstname ต้องไม่ว่าง")
	}
	if reqOwner.Lastname == "" {
		return errors.New("lastname ต้องไม่ว่าง")
	}
	Phonenumber, err := regexp.MatchString("^[0-9]{9,10}$", reqOwner.Phonenumber)
	if !Phonenumber {
		return errors.New("phonenumber ต้องไม่ต่ำกว่า 9 ตัว และ ไม่เกิน 10 ตัว ต้องมีแต่ตัวเลขเท่านั้น")
	}
	LineId, err := regexp.MatchString("^[a-z0-9._-]{1,20}$", reqOwner.Lineid)
	if !LineId {
		return errors.New("lineid ต้องไม่ต่ำกว่า 1 ตัว และ ไม่เกิน 20 ตัว ไม่สามารถใส่ตัวพิมพ์ใหญ่ได้และมีอักษรพิเศษได้แค่ ._- เท่านั้น")
	}
	_, err = ctrl.access.RDBMS.CheckUserRegister(reqOwner.Username, reqOwner.Phonenumber, reqOwner.Lineid, 0)
	if err != nil {
		Error = err
		return
	}
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

	res , err := ctrl.access.RDBMS.GetAccount(1)
	if res.Username == "admin"{
		Error = errors.New("สมัครไปแล้ว")
		return
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
	reqOperator.Username = strings.ToLower(reqOperator.Username)
	reqOperator.Password = strings.Trim(reqOperator.Password, "\t \n")
	reqOperator.Firstname = strings.Trim(reqOperator.Firstname, "\t \n")
	reqOperator.Lastname = strings.Trim(reqOperator.Lastname, "\t \n")
	reqOperator.Lineid = strings.Trim(reqOperator.Lineid, "\t \n")
	reqOperator.Phonenumber = strings.Trim(reqOperator.Phonenumber, "\t \n")
	user, err := regexp.MatchString("^[a-z0-9_-]{6,20}$", reqOperator.Username)
	if !user {
		return errors.New("username ต้องไม่ต่ำกว่า 6 ตัว และ ไม่เกิน 20 ตัว และมีอักษรพิเศษได้แค่ _- เท่านั้น")
	}
	pass, err := regexp.MatchString("^[a-zA-Z0-9_!-]{6,20}$", reqOperator.Password)
	if !pass {
		return errors.New("password ต้องไม่ต่ำกว่า 6 ตัว และ ไม่เกิน 20 ตัว ต้องมีตัวพิมพ์ใหญ่และตัวพิมพ์เล็กตัวเลขและมีอักษรพิเศษได้แค่ !_- เท่านั้น")
	}
	err = validPassword(reqOperator.Password)
	if err != nil {
		return errors.New("password ต้องไม่ต่ำกว่า 6 ตัว และ ไม่เกิน 20 ตัว ต้องมีตัวพิมพ์ใหญ่และตัวพิมพ์เล็กตัวเลขและมีอักษรพิเศษได้แค่ !_- เท่านั้น")
	}
	if !(len(reqOperator.Firstname) <= 30) {
		return errors.New("firstname ต้องไม่เกิน 30 ตัว")
	}
	if !(len(reqOperator.Lastname) <= 30) {
		return errors.New("lastname ต้องไม่เกิน 30 ตัว")
	}
	if reqOperator.Firstname == "" {
		return errors.New("firstname ต้องไม่ว่าง")
	}
	if reqOperator.Lastname == "" {
		return errors.New("lastname ต้องไม่ว่าง")
	}
	Phonenumber, err := regexp.MatchString("^[0-9]{9,10}$", reqOperator.Phonenumber)
	if !Phonenumber {
		return errors.New("phonenumber ต้องไม่ต่ำกว่า 9 ตัว และ ไม่เกิน 10 ตัว ต้องมีแต่ตัวเลขเท่านั้น")
	}
	LineId, err := regexp.MatchString("^[a-z0-9._-]{1,20}$", reqOperator.Lineid)
	if !LineId {
		return errors.New("lineid ต้องไม่ต่ำกว่า 1 ตัว และ ไม่เกิน 20 ตัว ไม่สามารถใส่ตัวพิมพ์ใหญ่ได้และมีอักษรพิเศษได้แค่ ._- เท่านั้น")
	}
	_, err = ctrl.access.RDBMS.CheckUserRegister(reqOperator.Username, reqOperator.Phonenumber, reqOperator.Lineid, 0)
	if err != nil {
		Error = err
		return
	}
	OwnerId := int(*reqOperator.SubOwnerId)
	data, err := ctrl.access.RDBMS.GetAccount(OwnerId)
	if err != nil {
		Error = errors.New("ไม่มีเจ้าของคนนี้ในระบบ")
		return
	}
	if data.ID == 0 {
		Error = errors.New("there is no owner of this id in the system")
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
		Error = errors.New("incorrect password")
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
			Error = errors.New("incorrect password")
			return
		}
		Token, err = utility.AuthenticationLogin(data.ID, data.Role)
		if err != nil {
			Error = err
			return
		}
	} else {
		Error = errors.New("ไม่มีสิทธิ์ในการเข้าสู้ระบบ")
		return
	}

	return Token, nil
}

func (ctrl *APIControl) GetAccount(id int) (response structure.UserAccount, Error error) {
	data, err := ctrl.access.RDBMS.CheckAccountId(uint(id))
	if err != nil {
		Error = errors.New("ไม่พบผู้ใช้งาน")
		return
	}
	response = structure.UserAccount{
		Id:          int(data.ID),
		FirstName:   data.FirstName,
		LastName:    data.LastName,
		PhoneNumber: data.PhoneNumber,
		LineId:      data.LineId,
		Role:        data.Role,
		SubOwnerId:  data.SubOwnerId,
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
	for _, data := range res {
		id := int(data.ID)
		UserAccountStructure := structure.UserAccountOwner{
			Id:          id,
			UserName:    data.Username,
			FirstName:   data.FirstName,
			LastName:    data.LastName,
			PhoneNumber: data.PhoneNumber,
			LineId:      data.LineId,
			Role:        data.Role,
			CreatedAt:   data.CreatedAt,
			UpdatedAt:   data.UpdatedAt,
		}
		DataArray = append(DataArray, UserAccountStructure)
	}
	//response.Paginator = &pagination
	response = DataArray
	return
}

func (ctrl *APIControl) GetSubOwner(OwnerId int) (response structure.GetSubOwner, Error error) {
	res, err := ctrl.access.RDBMS.CheckAccountId(uint(OwnerId))
	if err != nil {
		Error = errors.New("record not found")
		return
	}
	if res.Role != string(constant.Owner) {
		Error = errors.New("สิทธิ์ของคุณไม่ถูกต้อง")
		return
	}
	ops, err := ctrl.access.RDBMS.GetSubOwner(OwnerId)
	if err != nil {
		Error = err
		return
	}
	UserAccountStructure := structure.GetSubOwner{}
	var UserAccountOperatorArray []structure.Operators
	{
	}
	for _, data := range ops {
		for _, dataOps := range data.OpsAccount {
			UserAccountOperator := structure.Operators{
				OperatorId:          dataOps.ID,
				OperatorUserName:    dataOps.Username,
				OperatorFirstName:   dataOps.FirstName,
				OperatorLastName:    dataOps.LastName,
				OperatorPhoneNumber: dataOps.PhoneNumber,
				OperatorLineId:      dataOps.LineId,
				CreatedAt:           dataOps.CreatedAt,
				UpdatedAt:           dataOps.UpdatedAt,
			}
			UserAccountOperatorArray = append(UserAccountOperatorArray, UserAccountOperator)
		}
		UserAccountStructure = structure.GetSubOwner{
			OwnerId:             data.ID,
			OwnerUserName:       data.Username,
			OwnerFirstName:      data.FirstName,
			OwnerLastName:       data.LastName,
			OwnerPhoneNumber:    data.PhoneNumber,
			OwnerLineId:         data.LineId,
			UserAccountOperator: UserAccountOperatorArray,
		}
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
	for _, data := range res {
		id := int(data.ID)

		owner, err := ctrl.access.RDBMS.GetAccount(int(*data.SubOwnerId))
		if err != nil {
			Error = err
			return
		}

		UserAccountStructure := structure.UserAccountOperator{
			OperatorId:          id,
			OperatorUserName:    data.Username,
			OperatorFirstName:   data.FirstName,
			OperatorLastName:    data.LastName,
			OperatorPhoneNumber: data.PhoneNumber,
			OperatorLineId:      data.LineId,
			OwnerId:             *data.SubOwnerId,
			OwnerName:           owner.FirstName + " " + owner.LastName,
			CreatedAt:           data.CreatedAt,
			UpdatedAt:           data.UpdatedAt,
		}
		DataArray = append(DataArray, UserAccountStructure)
	}
	response = DataArray
	return
}

func (ctrl *APIControl) GetOwnerByIdOps(OperatorId int) (response structure.GetOwnerByOperator, Error error) {

	check, err := ctrl.access.RDBMS.CheckAccountId(uint(OperatorId))
	if err != nil {
		Error = errors.New("record not found")
		return
	}
	if check.Role != string(constant.Operator) {
		Error = errors.New("สิทธิ์ของคุณไม่ถูกต้อง")
		return
	}
	ops, err := ctrl.access.RDBMS.GetOwnerByIdOps(OperatorId)
	if err != nil {
		Error = err
		return
	}

	opsId := int(*ops.SubOwnerId)
	owner, err := ctrl.access.RDBMS.GetAccount(opsId)
	if err != nil {
		Error = err
		return
	}
	response = structure.GetOwnerByOperator{
		Operator: structure.Operator{
			Id:          ops.ID,
			UserName:    ops.Username,
			FirstName:   ops.FirstName,
			LastName:    ops.LastName,
			PhoneNumber: ops.PhoneNumber,
			LineId:      ops.LineId,
			CreatedAt:   ops.CreatedAt,
			UpdatedAt:   ops.UpdatedAt,
			Owner: structure.Owner{
				OwnerId:     owner.ID,
				FirstName:   owner.FirstName,
				LastName:    owner.LastName,
				PhoneNumber: owner.PhoneNumber,
				LineId:      owner.LineId,
				CreatedAt:   owner.CreatedAt,
				UpdatedAt:   owner.UpdatedAt,
			},
		},
	}
	return
}

func (ctrl *APIControl) UpdateProfile(id uint, Account *structure.UpdateProFile) (Error error) {
	Account.FirstName = strings.Trim(Account.FirstName, "\t \n")
	Account.LastName = strings.Trim(Account.LastName, "\t \n")
	Account.LineId = strings.Trim(Account.LineId, "\t \n")
	Account.PhoneNumber = strings.Trim(Account.PhoneNumber, "\t \n")

	res, err := ctrl.access.RDBMS.CheckAccountId(id)
	if err != nil {
		Error = errors.New("record not found")
		return
	}
	if !(res.PhoneNumber == Account.PhoneNumber) {
		_, err = ctrl.access.RDBMS.CheckUserRegister("", Account.PhoneNumber, "", 0)
		if err != nil {
			Error = err
			return
		}
	}
	if !(res.LineId == Account.LineId) {
		_, err = ctrl.access.RDBMS.CheckUserRegister("", "", Account.LineId, 0)
		if err != nil {
			Error = err
			return
		}
	}

	if !(len(Account.FirstName) <= 30) {
		return errors.New("firstname ต้องไม่เกิน 30 ตัว")
	}
	if !(len(Account.LastName) <= 30) {
		return errors.New("lastname ต้องไม่เกิน 30 ตัว")
	}
	if Account.FirstName == "" {
		return errors.New("firstname ต้องไม่ว่าง")
	}
	if Account.LastName == "" {
		return errors.New("lastname ต้องไม่ว่าง")
	}

	Phonenumber, err := regexp.MatchString("^[0-9]{9,10}$", Account.PhoneNumber)
	if !Phonenumber {
		return errors.New("phonenumber ต้องไม่ต่ำกว่า 9 ตัว และ ไม่เกิน 10 ตัว ต้องมีแต่ตัวเลขเท่านั้น")
	}
	LineId, err := regexp.MatchString("^[a-z0-9._-]{1,20}$", Account.LineId)
	if !LineId {
		return errors.New("lineid ต้องไม่ต่ำกว่า 1 ตัว และ ไม่เกิน 20 ตัว ไม่สามารถใส่ตัวพิมพ์ใหญ่ได้และมีอักษรพิเศษได้แค่ ._- เท่านั้น")
	}

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
	err = ctrl.access.RDBMS.UpdateProfile(data)
	if err != nil {
		Error = err
		return
	}
	return
}

func (ctrl *APIControl) ChangePassword(id uint, password *structure.ChangePassword) (Error error) {
	_, err := ctrl.access.RDBMS.CheckAccountId(id)
	if err != nil {
		Error = errors.New("record not found")
		return
	}
	pass, err := regexp.MatchString("^[a-zA-Z0-9_!-]{6,20}$", password.Password)
	if !pass {
		return errors.New("password ต้องไม่ต่ำกว่า 6 ตัว และ ไม่เกิน 20 ตัว ต้องมีตัวพิมพ์ใหญ่และตัวพิมพ์เล็กตัวเลขและมีอักษรพิเศษได้แค่ !_- เท่านั้น")
	}
	err = validPassword(password.Password)
	if err != nil {
		return errors.New("password ต้องไม่ต่ำกว่า 6 ตัว และ ไม่เกิน 20 ตัว ต้องมีตัวพิมพ์ใหญ่และตัวพิมพ์เล็กตัวเลขและมีอักษรพิเศษได้แค่ !_- เท่านั้น")
	}
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
	err = ctrl.access.RDBMS.UpdateProfile(data)
	if err != nil {
		Error = err
		return
	}
	return
}

func (ctrl *APIControl) DeleteAccount(id uint) (Error error) {
	res, err := ctrl.access.RDBMS.CheckAccountId(id)
	if err != nil {
		Error = errors.New("record not found")
		return
	}
	if res.Role == string(constant.Owner) {
		ops, err := ctrl.access.RDBMS.GetSubOwner(int(id))
		if err != nil {
			Error = errors.New("record not found")
			return
		}
		err = ctrl.access.RDBMS.DeleteAccount(id)
		if err != nil {
			Error = err
			return
		}
		for _, data := range ops {
			for _, delOps := range data.OpsAccount{
				err = ctrl.access.RDBMS.DeleteAccount(delOps.ID)
				if err != nil {
					Error = err
					return
				}
			}
		}
	} else {
		err = ctrl.access.RDBMS.DeleteAccount(id)
		if err != nil {
			Error = err
			return
		}
	}
	return
}

//func (ctrl *APIControl) DeleteAccountOwner(OwnerId uint) (Error error) {
//	res, err := ctrl.access.RDBMS.CheckAccountId(OwnerId)
//	if err != nil {
//		Error = errors.New("record not found")
//		return
//	}
//	if res.Role == string(constant.Owner){
//		ops, err := ctrl.access.RDBMS.GetSubOwner(int(OwnerId))
//		if err != nil {
//			Error = errors.New("record not found")
//			return
//		}
//		for _, data := range ops {
//			err = ctrl.access.RDBMS.DeleteAccount(data.ID)
//			if err != nil {
//				Error = err
//				return
//			}
//		}
//	}else {
//		err = ctrl.access.RDBMS.DeleteAccount(id)
//		if err != nil {
//			Error = err
//			return
//		}
//	}
//
//	return
//}

func (ctrl *APIControl) insert(Account rdbmsstructure.Account) error {
	err := ctrl.access.RDBMS.Register(Account)
	if err != nil {
		return err
	}
	return nil
}
