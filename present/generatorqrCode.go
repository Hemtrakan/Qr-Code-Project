package present

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"qrcode/access/constant"
	"qrcode/control"
	"qrcode/present/structure"
	"qrcode/utility"
)

func genQrCode(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	var files = new(structure.GenQr)
	if err := context.BodyParser(files); err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	file, err := api.GenQrCode(*files)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return context.Download(file)
}

func genQrCodeByName(context *fiber.Ctx) error {
	name := context.Params("name")
	fileimage := string(constant.SaveFileLocationQrCode) + "/" + name + ".PNG"
	return context.Download(fileimage)
}
