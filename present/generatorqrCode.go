package present

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"net/http"
	"qrcode/access/constant"
	"qrcode/control"
	"qrcode/present/structure"
	"qrcode/utility"
	"strconv"
)


func getTemplate(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	res := api.GetTemplate()
	return context.Status(http.StatusOK).JSON(res)
}


func getQrCodeById(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	id := context.Params("id")
	ownerId, err := strconv.Atoi(id)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest,"กรอกได้แต่ตัวเลขเท่านั้น")
	}
	res, err := api.GetQrCodeById(ownerId)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest,err.Error())
	}
	return context.JSON(res)
}

func getAllQrCode(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	res, err := api.GetAllQrCode()
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest,err.Error())
	}
	return context.JSON(res)
}

func insertDataQrCode(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	var data = new(structure.InsertDataQrCode)
	if err := context.BodyParser(data); err != nil {
		return utility.FiberError(context, http.StatusBadRequest, "ส่งชนิดของข้อมูลมาผิด")
	}
	err := ValidateStruct(*data)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}

	err = api.InsertDataQrCode(data)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest,err.Error())
	}
	return utility.FiberSuccess(context , http.StatusOK,"บันทึกข้อมูลสำเร็จ")
}

func getDataQrCodeJson(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	id := context.Params("id")
	res, err := api.GetDataQrCode(id)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return context.Status(http.StatusOK).JSON(res)
}

func getDataQrCode(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	id := context.Params("id")
	contentType := context.Get("Content-Type")
	if contentType == "" {
		url := "http://192.168.1.104:12000/viewdata/" + id
		if err := proxy.Do(context, url); err != nil {
			return err
		}
		// Remove Server header from response
		context.Response().Header.Del(fiber.HeaderServer)
		return nil
	} else if contentType == "application/json;charset=UTF-8" {
		res, err := api.GetDataQrCode(id)
		if err != nil {
			return utility.FiberError(context, http.StatusBadRequest, err.Error())
		}
		return context.Status(http.StatusOK).JSON(res)
	} else {
		res, err := api.GetDataQrCode(id)
		if err != nil {
			return utility.FiberError(context, http.StatusBadRequest, err.Error())
		}
		return context.Status(http.StatusOK).JSON(res)
	}
}

func createQrCode(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	var files = new(structure.GenQrCode)
	if err := context.BodyParser(files); err != nil {
		return utility.FiberError(context, http.StatusBadRequest, "ส่งชนิดของข้อมูลมาผิด")
	}
	err := ValidateStruct(*files)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	err = api.CreateQrCode(*files)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	//return context.Download(fileZip)
	return utility.FiberError(context, http.StatusOK, "สร้าง QrCode สำเร็จ")
}

func genQrCodeToFileZipByQrCodeId(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	var data = new(structure.FileZip)
	if err := context.BodyParser(data); err != nil {
		return utility.FiberError(context, http.StatusBadRequest,"ส่งชนิดของข้อมูลมาผิด")
	}
	err := ValidateStruct(*data)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	fileZip, err := api.AddFileZipById(*data)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return context.Download(fileZip)
}

func genQrCodeToFileZipByTemplateName(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	var data = new(structure.FileZipByTemplateName)
	if err := context.BodyParser(data); err != nil {
		return utility.FiberError(context, http.StatusBadRequest, "ส่งชนิดของข้อมูลมาผิด")
	}
	err := ValidateStruct(*data)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	fileZip, err := api.AddFileZipByTemplateName(*data)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return context.Download(fileZip)
}

func deleteQrCode(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	var QrCode = new(structure.DelQrCode)
	if err := context.BodyParser(QrCode); err != nil {
		return utility.FiberError(context, http.StatusBadRequest, "ส่งชนิดของข้อมูลมาผิด")
	}
	err := ValidateStruct(*QrCode)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	err = api.DeleteQrCode(*QrCode)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return utility.FiberSuccess(context, http.StatusOK, "ลบ QrCode สำเร็จ")
}

func updateStatusQrCode(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	var QrCode = new(structure.StatusQrCode)
	if err := context.BodyParser(QrCode); err != nil {
		return utility.FiberError(context, http.StatusBadRequest, "ส่งชนิดของข้อมูลมาผิด")
	}
	QrCodeId := context.Params("id")
	err := ValidateStruct(*QrCode)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	err = api.UpdateStatusQrCode(QrCodeId,*QrCode)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return utility.FiberSuccess(context, http.StatusOK, "เปลี่ยนสถานะ QrCode สำเร็จ")
}



//func genQrCodeByName(context *fiber.Ctx) error {
//	name := context.Params("name")
//	fileImage := string(constant.SaveFileLocationQrCode) + "/" + name + ".PNG"
//	return context.Download(fileImage)
//}