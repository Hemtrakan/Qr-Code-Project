package utility

import (
	"errors"
	"qrcode/access/constant"
	"qrcode/present/structure/templates/computer"
	"qrcode/present/structure/templates/knex"
)

func CheckTemplate(Templates string) (result interface{}, Error error) {
	if Templates == string(constant.Computer) {
		result = computer.Info{}
		return
	}
	if Templates == string(constant.ServiceWashingMachine) {
		result = knex.WashingMachineInfo{}
		return
	}
	Error = errors.New("this template does not exist in the system")
	return
}
