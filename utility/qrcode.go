package utility

import (
	"errors"
	"log"
	"os"
	"qrcode/access/constant"
)

func CreatesNewDirectory() (Error error) {
	err := os.Mkdir(string(constant.SaveFileLocationQrCode), 0755)
	if err != nil {
		Error = errors.New("ไม่สามารถสร้าง SaveFileLocationQrCode ได้")
		return
	}
	err = os.Mkdir(string(constant.SaveFileLocationZipFile), 0755)
	if err != nil {
		Error = errors.New("ไม่สามารถสร้าง SaveFileLocationZipFile ได้")
		return
	}
	err = os.Mkdir(string(constant.SaveFileLocationExposed), 0755)
	if err != nil {
		Error = errors.New("ไม่สามารถสร้าง SaveFileLocationExposed ได้")
		return
	}
	return
}

func RemoveAllFileLocation() {
	err :=os.RemoveAll(string(constant.SaveFileLocationZipFile))
	err = os.RemoveAll(string(constant.SaveFileLocationQrCode))
	err = os.RemoveAll(string(constant.SaveFileLocationExposed))
	if err != nil {
		log.Println("err",err.Error())
	}
	return
}