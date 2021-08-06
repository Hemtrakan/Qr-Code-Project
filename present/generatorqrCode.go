package present

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"qrcode/access/constant"
	"qrcode/control"
	"qrcode/present/structure"
	"qrcode/utility"
)

func genQrCodeBD(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	GenerateQrCode := new(structure.GenerateQrCode)
	if err := context.BodyParser(GenerateQrCode); err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	err := api.GenerateQrCode(GenerateQrCode)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return utility.FiberSuccess(context, http.StatusOK, "success")
}

func genQrCode(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	var reqfiles = new(structure.GenQr)
	if err := context.BodyParser(reqfiles); err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	file, err := api.GenQrCode(*reqfiles)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	//filezip := "./" + file
	return utility.FiberError(context, http.StatusOK, "success")
}
