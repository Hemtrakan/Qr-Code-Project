package present

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"net/http"
	"qrcode/access/constant"
	"qrcode/control"
	"qrcode/environment"
	"qrcode/present/structure"
	"qrcode/utility"
	"strconv"
)

type GORMFactory struct {
	env *environment.Properties
}

func getTemplateList(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	res := api.GetTemplateList()
	return context.Status(http.StatusOK).JSON(res)
}

func getTemplate(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	Template := context.Params("id")
	res ,err:= api.GetTemplate(Template)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return context.Status(http.StatusOK).JSON(res)
}

func getQrCodeById(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	id := context.Params("id")
	ownerId, err := strconv.Atoi(id)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, "กรอกได้แต่ตัวเลขเท่านั้น")
	}
	res, err := api.GetQrCodeById(ownerId)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return context.JSON(res)
}

func getAllQrCode(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	res, err := api.GetAllQrCode()
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
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
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return utility.FiberSuccess(context, http.StatusOK, "บันทึกข้อมูลสำเร็จ")
}

func updateDataQrCode(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	var data = new(structure.UpdateDataQrCode)
	if err := context.BodyParser(data); err != nil {
		return utility.FiberError(context, http.StatusBadRequest, "ส่งชนิดของข้อมูลมาผิด")
	}
	err := ValidateStruct(*data)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	err = api.UpdateDataQrCode(data)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return utility.FiberSuccess(context, http.StatusOK, "บันทึกข้อมูลสำเร็จ")
}

func updateHistoryInfoDataQrCode(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	var data = new(structure.UpdateHistoryInfoDataQrCode)
	if err := context.BodyParser(data); err != nil {
		return utility.FiberError(context, http.StatusBadRequest, "ส่งชนิดของข้อมูลมาผิด")
	}
	err := ValidateStruct(*data)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	err = api.UpdateHistoryInfoDataQrCode(data)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return utility.FiberSuccess(context, http.StatusOK, "บันทึกข้อมูลสำเร็จ")
}

func updateOpsDataQrCode(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	var data = new(structure.UpdateOpsDataQrCode)
	if err := context.BodyParser(data); err != nil {
		return utility.FiberError(context, http.StatusBadRequest, "ส่งชนิดของข้อมูลมาผิด")
	}
	err := ValidateStruct(*data)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}

	err = api.UpdateOpsDataQrCode(data)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return utility.FiberSuccess(context, http.StatusOK, "บันทึกข้อมูลสำเร็จ")
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
	// สำหรับมี Template แล้ว
	contentType := context.Get("Content-Type")
	if contentType == "" {
		url := string(environment.URLFront) + context.Params("*")
		if err := proxy.Do(context, url); err != nil {
			return err
		}
		// Remove Server header from response
		context.Response().Header.Del(fiber.HeaderServer)
		return nil
	}
	return utility.FiberError(context,http.StatusBadRequest,"ไม่สามารถแสดงหน้านี้ได้")
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
		return utility.FiberError(context, http.StatusBadRequest, "ส่งชนิดของข้อมูลมาผิด")
	}
	err := ValidateStruct(*data)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	err = utility.CreatesNewDirectory()
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	fileZip, err := api.AddFileZipById(*data)
	if err != nil {
		utility.RemoveAllFileLocation()
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	defer utility.RemoveAllFileLocation()
	return context.Download(fileZip)
}

func genQrCodeToFileZipByOwner(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	var data = new(structure.FileZipByOwner)
	if err := context.BodyParser(data); err != nil {
		return utility.FiberError(context, http.StatusBadRequest, "ส่งชนิดของข้อมูลมาผิด")
	}
	err := ValidateStruct(*data)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	err = utility.CreatesNewDirectory()
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	fileZip, err := api.AddFileZipByOwner(*data)
	if err != nil {
		utility.RemoveAllFileLocation()
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	defer utility.RemoveAllFileLocation()
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
	err = utility.CreatesNewDirectory()
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	fileZip, err := api.AddFileZipByTemplateName(*data)
	if err != nil {
		utility.RemoveAllFileLocation()
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	defer utility.RemoveAllFileLocation()
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
	err = api.UpdateStatusQrCode(QrCodeId, *QrCode)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return utility.FiberSuccess(context, http.StatusOK, "เปลี่ยนสถานะ QrCode สำเร็จ")
}

// todo test
//func TestQrCode(context *fiber.Ctx) error {
//	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
//
//	res, err := api.TestQrCode()
//	if err != nil {
//		return utility.FiberError(context, http.StatusBadRequest, err.Error())
//	}
//	return context.Status(http.StatusOK).JSON(res)
//}

//func genQrCodeByName(context *fiber.Ctx) error {
//	name := context.Params("name")
//	fileImage := string(constant.SaveFileLocationQrCode) + "/" + name + ".PNG"
//	return context.Download(fileImage)
//}
