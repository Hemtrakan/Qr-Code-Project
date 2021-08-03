package present

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"qrcode/access/constant"
	"qrcode/control"
	"qrcode/present/structure"
	"qrcode/utility"
)

func uploadFile(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	// Source
	if form, err := context.MultipartForm(); err == nil {
		files := form.File["file"]
		fileCategory := form.Value["type"][0]

		for _, file := range files {
			var fileType constant.CategoryFile
			switch fileCategory {
			case string(constant.QRCode):
				fileType = constant.QRCode
				break
			case string(constant.TeamPage):
				fileType = constant.TeamPage
				break
			default:
				log.Fatal("Unimplemented")
			}
			fileName, err := api.UploadFile(*file, fileType)
			if err != nil {
				return utility.FiberError(context,http.StatusBadRequest,err.Error())
			}
			return context.JSON(structure.FileResponse{Data: *fileName})
		}
	}

	return  utility.FiberError(context,http.StatusBadRequest,"upload error")

}

func getUrlFile(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	file := new(structure.FileRequest)
	if err := context.QueryParser(file); err != nil {
		return utility.FiberError(context,http.StatusBadRequest,err.Error())
	}

	var fileType constant.CategoryFile
	switch file.Data {
	case string(constant.QRCode):
		fileType = constant.QRCode
		break
	case string(constant.TeamPage):
		fileType = constant.TeamPage
		break
	default:
		log.Fatal("Unimplemented")
	}

	fileURL, err := api.GetUrlFile(file.File, fileType)
	if err != nil {
		return err
	}
	return context.Status(http.StatusOK).SendString(fileURL.String())
}
