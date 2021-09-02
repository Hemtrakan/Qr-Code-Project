package utility

import (
	"errors"
	"qrcode/access/constant"
	"qrcode/present/structure/templates/computer"
	"qrcode/present/structure/templates/officeequipment"
)

func CheckTemplate(Templates string) (result interface{}, Error error) {
	if Templates == ""{
		result = ""
		return
	}
	if Templates == string(constant.Computer) {
		result = computer.Info{}
		return
	}
	if Templates == string(constant.OfficeEquipment) {
		result = officeequipment.Info{}
		return
	}
	Error = errors.New("ไม่มี template นี้อยู่ในระบบ")
	return
}

