package control

import (
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/datatypes"
	rdbmsstructure "qrcode/access/rdbms/structure"
	"qrcode/present/structure/templates/computer"
	"time"
)

func (ctrl *APIControl) InsertComputer(QrCodeId string, req computer.Info) (Error error) {
	res, err := ctrl.access.RDBMS.GetDataQrCode(QrCodeId)
	if err != nil {
		Error = errors.New("ไม่มี QrCode นี้อยู่ในระบบ")
		return
	}

	if res.First == true {
		Error = errors.New("QrCode นี้ ถูกเพิ่มข้อมูลไปแล้ว")
		return
	}

	json, err := json.Marshal(req)
	if err != nil {
		Error = errors.New("ไม่สามารถแปลงเป็น Json ได้")
		return
	}
	QrCodeData := rdbmsstructure.QrCode{
		Info:  datatypes.JSON(json),
		First: true,
	}
	err = ctrl.access.RDBMS.InsertQrCode(QrCodeId, QrCodeData)
	if err != nil {
		Error = errors.New("ไม่สามารถบันทึกข้อมูลได้")
		return
	}
	return
}

func (ctrl *APIControl) UpdateComputer(QrCodeId string, req computer.Info) (Error error) {
	res, err := ctrl.access.RDBMS.GetDataQrCode(QrCodeId)
	if err != nil {
		Error = errors.New("ไม่มี QrCode นี้อยู่ในระบบ")
		return
	}

	jsonReq, err := json.Marshal(req)
	if err != nil {
		Error = errors.New("ไม่สามารถแปลงเป็น Json ได้")
		return
	}


	QrCodeData := rdbmsstructure.QrCode{
		Info:        datatypes.JSON(jsonReq),
	}
	err = ctrl.access.RDBMS.InsertQrCode(QrCodeId, QrCodeData)
	if err != nil {
		Error = errors.New("ไม่สามารถแก้ไขข้อมูลได้")
		return
	}

	History := rdbmsstructure.History{
		QrCodeUUID:  res.QrCodeUUID,
		HistoryInfo: res.Info,
		UserId:      1,
		UpdatedAt:   time.Now(),
	}

	err = ctrl.access.RDBMS.InsertHistory(History)
	if err != nil {
		fmt.Println("")
		Error = errors.New("ไม่สามารถแก้ไขข้อมูลได้")
		return
	}
	return
}
