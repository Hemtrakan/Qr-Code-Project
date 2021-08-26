package present

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	gim "github.com/ozankasikci/go-image-merge"
	"github.com/pbnjay/pixfont"
	"image"
	"image/color"
	"image/jpeg"
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
	//contentType := context.Get("Content-Type")
	//if contentType == "" {
	//	url := "http://192.168.1.104:12000/viewdata/" + id
	//	if err := proxy.Do(context, url); err != nil {
	//		return err
	//	}
	//	// Remove Server header from response
	//	context.Response().Header.Del(fiber.HeaderServer)
	//	return nil
	//} else if contentType == "application/json;charset=UTF-8" {
	//	res, err := api.GetDataQrCode(id)
	//	if err != nil {
	//		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	//	}
	//	return context.Status(http.StatusOK).JSON(res)
	//} else {
		res, err := api.GetDataQrCode(id)
		if err != nil {
			return utility.FiberError(context, http.StatusBadRequest, err.Error())
		}
		return context.Status(http.StatusOK).JSON(res)
	//}
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
	fileZip, err := api.AddFileZipByOwner(*data)
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

func GetFileContentType(out *os.File) (string, error) {

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}

func test(context *fiber.Ctx) error {
	//imgFile1, err := os.Open("fileqrcode/bell-1.jpeg")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//imgFile2, err := os.Open("filetext/bell-1.jpeg")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//img1, _, err := image.Decode(imgFile1)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//img2, _, err := image.Decode(imgFile2)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//sp2 := image.Point{img1.Bounds().Dx(), 0}
	//r2 := image.Rectangle{sp2, sp2.Add(img2.Bounds().Size())}
	//r := image.Rectangle{image.Point{0, 0}, r2.Max}
	//rgba := image.NewRGBA(r)
	//draw.Draw(rgba, img1.Bounds(), img1, image.Point{0, 0}, draw.Src)
	//draw.Draw(rgba, r2, img2, image.Point{100, 0}, draw.Src)
	//out, err := os.Create("qr/output.jpg")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//var opt jpeg.Options
	//opt.Quality = 80
	//jpeg.Encode(out, rgba, &opt)


	//f1, err := os.Open("qr/bell-1.jpeg")
	//if err != nil {
	//	panic(err)
	//}
	//defer f1.Close()
	//
	//// Get the content
	//contentType, err := GetFileContentType(f1)
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println("Content Type: " + contentType)


	//
	img := image.NewRGBA(image.Rect(0, 0, 100, 20))
	pixfont.DrawString(img, 10, 10, "bell-1", color.White)
	f, _ := os.OpenFile("qr/bell-5.PNG", os.O_CREATE|os.O_RDWR, 0644)
	png.Encode(f, img)
	file, err := os.Create("qr/computer-1.png")
	if err != nil {
		fmt.Println("2 : ",err.Error())
	}


	grids := []*gim.Grid{
		{ImageFilePath: "fileqrcode/bell-1.jpg"},
		{ImageFilePath: "qr/bell-5.png"},
	}
	rgba, err := gim.New(grids, 1, 2).Merge()
	if err != nil {
		fmt.Println("1 : ",err.Error())
	}

	err = jpeg.Encode(file, rgba, &jpeg.Options{Quality: 80})
	if err != nil {
		fmt.Println("3 : ",err.Error())

	}
	return utility.FiberSuccess(context, http.StatusOK, "ทดสอบ")
}



//func genQrCodeByName(context *fiber.Ctx) error {
//	name := context.Params("name")
//	fileImage := string(constant.SaveFileLocationQrCode) + "/" + name + ".PNG"
//	return context.Download(fileImage)
//}