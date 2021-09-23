package control

import (
	"errors"
	"gorm.io/gorm"
	"qrcode/access/constant"
	rdbmsstructure "qrcode/access/rdbms/structure"
	"qrcode/present/structure"
	"qrcode/utility"
	"regexp"
	"strings"
	"time"
)

func (ctrl *APIControl) RegisterOperatorOwner(reqOperator *structure.RegisterOperator, OwnerId *uint) (Error error) {
	reqOperator.Username = strings.ToLower(reqOperator.Username)
	reqOperator.Password = strings.Trim(reqOperator.Password, "\t \n")
	reqOperator.Firstname = strings.Trim(reqOperator.Firstname, "\t \n")
	reqOperator.Lastname = strings.Trim(reqOperator.Lastname, "\t \n")
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
	_, err = ctrl.access.RDBMS.CheckUserRegister(reqOperator.Username, reqOperator.Phonenumber, 0)
	if err != nil {
		Error = err
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
		Role:        string(constant.Operator),
		SubOwnerId:  OwnerId,
	}
	err = ctrl.insert(Operator)
	if err != nil {
		Error = err
		return
	}
	return
}

func (ctrl *APIControl) GetOperatorById(OperatorId int, OwnerId uint) (response structure.UserAccount, Error error) {
	_, err := ctrl.access.RDBMS.CheckAccountId(uint(OperatorId))
	if err != nil {
		Error = errors.New("ไม่มีช่างซ่อมคนนี้")
		return
	}
	data, err := ctrl.access.RDBMS.GetOperatorById(OperatorId, OwnerId)
	if err != nil {
		Error = errors.New("ไม่มีช่างคนนี้")
		return
	}

	id := data.ID
	response = structure.UserAccount{
		Id:          int(id),
		FirstName:   data.FirstName,
		LastName:    data.LastName,
		PhoneNumber: data.PhoneNumber,
		Role:        data.Role,
		SubOwnerId:  data.SubOwnerId,
	}
	return
}

func (ctrl *APIControl) GetOperatorLine(OwnerId uint) (response []structure.OperatorsLine, Error error) {
	var DataArray []structure.OperatorsLine
	res, err := ctrl.access.RDBMS.GetAllAccountOperatorByOwnerID(OwnerId)
	if err != nil {
		Error = err
		return
	}
	for _, data := range res {
		if data.LineUserId != nil {
			UserAccountStructure := structure.OperatorsLine{
				OperatorId:          data.ID,
				OperatorUserName:    data.Username,
				OperatorFirstName:   data.FirstName,
				OperatorLastName:    data.LastName,
				OperatorLineId:      data.LineUserId,
			}
			DataArray = append(DataArray, UserAccountStructure)
		}
	}

	response = DataArray
	return
}

func (ctrl *APIControl) GetOperator(OwnerId uint) (response []structure.Operators, Error error) {
	var DataArray []structure.Operators
	res, err := ctrl.access.RDBMS.GetAllAccountOperatorByOwnerID(OwnerId)
	if err != nil {
		Error = err
		return
	}
	for _, data := range res {
		UserAccountStructure := structure.Operators{
			OperatorId:          data.ID,
			OperatorUserName:    data.Username,
			OperatorFirstName:   data.FirstName,
			OperatorLastName:    data.LastName,
			OperatorPhoneNumber: data.PhoneNumber,
			OperatorLineId:      data.LineUserId,
			CreatedAt:           data.CreatedAt,
			UpdatedAt:           data.UpdatedAt,
		}
		DataArray = append(DataArray, UserAccountStructure)
	}
	response = DataArray
	return
}

func (ctrl *APIControl) DeleteAccountOperator(OwnerId uint, OperatorId int) (Error error) {
	err := ctrl.access.RDBMS.DeleteAccountByOwner(OwnerId, OperatorId)
	if err != nil {
		Error = err
		return
	}
	return
}

func (ctrl *APIControl) ChangePasswordOperator(OwnerId uint, password structure.ChangePasswordOperator) (Error error) {
	pass, err := regexp.MatchString("^[a-zA-Z0-9_!-]{6,20}$", password.NewPassword)
	if !pass {
		return errors.New("password ต้องไม่ต่ำกว่า 6 ตัว และ ไม่เกิน 20 ตัว ต้องมีตัวพิมพ์ใหญ่และตัวพิมพ์เล็กตัวเลขและมีอักษรพิเศษได้แค่ !_- เท่านั้น")
	}
	err = validPassword(password.NewPassword)
	if err != nil {
		return errors.New("password ต้องไม่ต่ำกว่า 6 ตัว และ ไม่เกิน 20 ตัว ต้องมีตัวพิมพ์ใหญ่และตัวพิมพ์เล็กตัวเลขและมีอักษรพิเศษได้แค่ !_- เท่านั้น")
	}
	_, err = ctrl.access.RDBMS.GetOperatorById(password.OperatorId, OwnerId)
	if err != nil {
		Error = errors.New("เจ้าของคนนี้ไม่มีสิทธิ์เปลี่ยนรหัสผ่านช่างซ่อม")
	}
	changePassword, err := utility.Hash(password.NewPassword)
	if err != nil {
		Error = errors.New("ไม่สามารถเปลี่ยนรหัสผ่านได้ ติดต่อผู้ดูแลระบบ")
		return
	}
	change := rdbmsstructure.Account{
		Model: gorm.Model{
			ID:        uint(password.OperatorId),
			UpdatedAt: time.Now(),
		},
		Password: string(changePassword),
	}
	err = ctrl.access.RDBMS.UpdateProfile(change)
	if err != nil {
		Error = errors.New("ไม่สามารถเปลี่ยนรหัสผ่านได้ ติดต่อผู้ดูแลระบบ")
		return
	}
	return
}

func (ctrl *APIControl) ChangePasswordOwnerAndOperator(OwnerId uint, password structure.ChangePasswordOwnerAndOperator) (Error error) {
	_, err := ctrl.access.RDBMS.CheckAccountId(OwnerId)
	if err != nil {
		Error = errors.New("record not found")
		return
	}
	pass, err := regexp.MatchString("^[a-zA-Z0-9_!-]{6,20}$", password.NewPassword)
	if !pass {
		return errors.New("password ต้องไม่ต่ำกว่า 6 ตัว และ ไม่เกิน 20 ตัว ต้องมีตัวพิมพ์ใหญ่และตัวพิมพ์เล็กตัวเลขและมีอักษรพิเศษได้แค่ !_- เท่านั้น")
	}
	err = validPassword(password.NewPassword)
	if err != nil {
		return errors.New("password ต้องไม่ต่ำกว่า 6 ตัว และ ไม่เกิน 20 ตัว ต้องมีตัวพิมพ์ใหญ่และตัวพิมพ์เล็กตัวเลขและมีอักษรพิเศษได้แค่ !_- เท่านั้น")
	}
	data, err := ctrl.access.RDBMS.GetAccount(int(OwnerId))
	err = utility.VerifyPassword(data.Password, password.OldPassword)
	if err != nil {
		return errors.New("รหัสผ่านเก่าไม่ถูกต้อง")
	}
	changePassword, err := utility.Hash(password.NewPassword)
	if err != nil {
		Error = errors.New("ไม่สามารถเปลี่ยนรหัสผ่านได้ ติดต่อผู้ดูแลระบบ")
		return
	}
	change := rdbmsstructure.Account{
		Model: gorm.Model{
			ID:        OwnerId,
			UpdatedAt: time.Now(),
		},
		Password: string(changePassword),
	}
	err = ctrl.access.RDBMS.UpdateProfile(change)
	if err != nil {
		Error = errors.New("ไม่สามารถเปลี่ยนรหัสผ่านได้ ติดต่อผู้ดูแลระบบ")
		return
	}
	return
}

func (ctrl *APIControl) UpdateStatusQrCodeOwner(ownerId uint, QrCodeId string, req structure.StatusQrCode) (Error error) {
	res, err := ctrl.access.RDBMS.GetDataQrCode(QrCodeId)
	if err != nil {
		Error = errors.New("Qr-Code ที่จะเปลี่ยนสถานะไม่มีอยู่ในระบบ")
		return
	}
	for _, data := range res {
		if data.OwnerId != ownerId {
			Error = errors.New("ผู้ใช้งานไม่ถูกต้อง")
			return
		}

		Qr := rdbmsstructure.QrCode{
			QrCodeUUID: data.QrCodeUUID,
			Active:     *req.Active,
		}
		err = ctrl.access.RDBMS.UpdateQrCodeActive(Qr)
		if err != nil {
			Error = err
			return
		}
	}
	return
}
