package utility

import (
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/datatypes"
	"qrcode/access/constant"
	"qrcode/present/structure/templates/computer"
	"qrcode/present/structure/templates/knex"
	"reflect"
	"strings"
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
	Error = errors.New("ไม่มี template นี้อยู่ในระบบ")
	return
}

func CheckStructureTemplate(Templates string, obj interface{}) (result interface{}, Error error) {
	if Templates == string(constant.Computer) {
		byteInfoObj, err := json.Marshal(obj)
		if err != nil {
			Error = err
			return
		}

		byteInfo, err := json.Marshal(computer.Info{})
		if err != nil {
			Error = err
			return
		}
		fmt.Println("json : ", datatypes.JSON(byteInfoObj))
		fmt.Println("jsonByteInfo : ", datatypes.JSON(byteInfo))
		//infoObj := datatypes.JSON(byteInfoObj)
		//infos := datatypes.JSON(byteInfo)
		//if infoObj != infos {
		//
		//}



		// ดึง key ของ struct info ออกมา
		info := &computer.Info{}
		//long and bored code
		t := reflect.TypeOf(*info)
		if t.Kind() == reflect.Struct {
			for i := 0; i < t.NumField(); i++ {
				fmt.Println(strings.ToLower(t.Field(i).Name))
			}
		} else {
			Error = errors.New("not a struct")
			return
		}

		return
	}
	if Templates == string(constant.ServiceWashingMachine) {
		result = knex.ServiceWashingMachine{}
		return
	}
	Error = errors.New("StructureTemplate ไม่ถูกต้อง")
	return
}
