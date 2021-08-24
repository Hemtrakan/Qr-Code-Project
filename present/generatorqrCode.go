package present

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/png"
	"net/http"
	"os"
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
	fileZip,err := api.CreateQrCode(*files)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return context.Download(fileZip)
	//return utility.FiberError(context, http.StatusOK, "สร้าง QrCode สำเร็จ")
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

func Test(context *fiber.Ctx) error {
	path := string(constant.SaveFileLocationQrCode) + "/" + "computer-1.PNG"
	f, err := os.Create("path.PNG")
	if err != nil {
		// Handle error
	}
	defer f.Close()
	img := image.NewRGBA(image.Rect(0, 0, 200, 200))
	col := color.RGBA{68, 255, 236, 255}
	point := fixed.Point26_6{fixed.Int26_6(200 * 64), fixed.Int26_6(200 * 64)}
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString("computer-1")
	if err := png.Encode(f, img); err != nil {
		fmt.Println("1")
		panic(err)
	}
	return context.Download(path)
}

//func genQrCodeByName(context *fiber.Ctx) error {
//	name := context.Params("name")
//	fileImage := string(constant.SaveFileLocationQrCode) + "/" + name + ".PNG"
//	return context.Download(fileImage)
//}