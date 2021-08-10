package present

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"qrcode/access/constant"
	"qrcode/control"
	"qrcode/present/structure"
	"qrcode/utility"
	"strconv"
)

func GetQrCodeById(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	id := context.Params("id")
	ownerId , err := strconv.Atoi(id)
	res ,err := api.GetQrCodeById(ownerId)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return context.JSON(res)
}

func createQrCode(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	var files = new(structure.GenQrCode)
	if err := context.BodyParser(files); err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	err := api.CreateQrCode(*files)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return utility.FiberError(context, http.StatusOK, "succeed")
}

func genQrCodeToFileZipByQrCodeId(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	var data = new(structure.FileZip)
	if err := context.BodyParser(data); err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	fileZip , err := api.AddFileZipById(*data)
	if err != nil {
		return utility.FiberError(context,http.StatusBadRequest,err.Error())
	}
	return context.Download(fileZip)
}

func genQrCodeToFileZipByTemplateName(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	var data = new(structure.FileZipByTemplateName)
	if err := context.BodyParser(data); err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	fileZip , err := api.AddFileZipByTemplateName(*data)
	if err != nil {
		return utility.FiberError(context,http.StatusBadRequest,err.Error())
	}
	return context.Download(fileZip)
}



func genQrCodeByName(context *fiber.Ctx) error {
	name := context.Params("name")
	fileImage := string(constant.SaveFileLocationQrCode) + "/" + name + ".PNG"
	return context.Download(fileImage)
}
