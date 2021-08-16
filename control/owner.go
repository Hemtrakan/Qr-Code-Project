package control

import (
	"qrcode/access/constant"
	rdbmsstructure "qrcode/access/rdbms/structure"
	"qrcode/present/structure"
	"qrcode/utility"
)

func (ctrl *APIControl) RegisterOperatorOwner(reqOperator *structure.RegisterOperator,OwnerId uint) (Error error) {
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
		SubOwnerId:  OwnerId,
	}
	err = ctrl.insert(Operator)
	if err != nil {
		Error = err
		return
	}
	return
}

func (ctrl *APIControl) GetOperator(OwnerId uint) (response []structure.UserAccountOperator,Error error) {
	var DataArray []structure.UserAccountOperator
	res, err := ctrl.access.RDBMS.GetAllAccountOperatorByOwnerID(OwnerId)
	if err != nil {
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
			SubOwnerId:  int(data.SubOwnerId),
		}
		DataArray = append(DataArray, UserAccountStructure)
	}
	response = DataArray
	return
}

func (ctrl *APIControl) DeleteAccountOperator(OwnerId uint,OperatorId int) (Error error){
	err := ctrl.access.RDBMS.DeleteAccountByOwner(OwnerId,OperatorId)
	if err != nil {
		Error = err
		return
	}
	return
}
