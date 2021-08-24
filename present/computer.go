package present

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"qrcode/access/constant"
	"qrcode/control"
	"qrcode/present/structure/templates/computer"
	"qrcode/utility"
)

func Insert(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	QrCodeId := context.Params("id")
	computer := new(computer.Info)
	err := context.BodyParser(computer)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	err = ValidateStruct(*computer)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	err = api.InsertComputer(QrCodeId,*computer)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest,err.Error())
	}
	return utility.FiberSuccess(context,http.StatusOK,"บันทึกข้อมูลสำเร็จ")
}

func UpData(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	QrCodeId := context.Params("id")
	computer := new(computer.Info)
	err := context.BodyParser(computer)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	err = ValidateStruct(*computer)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	err = api.UpdateComputer(QrCodeId,*computer)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest,err.Error())
	}
	return utility.FiberSuccess(context,http.StatusOK,"บันทึกข้อมูลสำเร็จ")
}