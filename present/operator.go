package present

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"net/http"
	"qrcode/access/constant"
	"qrcode/control"
	"qrcode/present/structure"
	"qrcode/utility"
)

func getOperatorId(context *fiber.Ctx) (id uint, Error error) {
	user := context.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	var UserId = claims["id"].(float64)
	id = uint(UserId)
	return
}

func getAccountByLineId(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	lineId := context.Params("id")
	responses, err := api.GetAccountByLineId(lineId)
	if err != nil {
		return context.Status(http.StatusOK).JSON("")
	}
	return context.JSON(responses)
}

func ChangePasswordOperator(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	OwnerId, err := getOwnerId(context)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	changePasswordOwner := new(structure.ChangePasswordOwnerAndOperator)
	err = context.BodyParser(changePasswordOwner)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	err = ValidateStruct(*changePasswordOwner)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	err = api.ChangePasswordOwnerAndOperator(OwnerId, *changePasswordOwner)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return utility.FiberSuccess(context, http.StatusOK, "เปลี่ยนรหัสผ่านสำเร็จ")
}

func LoginOperator(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	Login := new(structure.LoginOperator)
	err := context.BodyParser(Login)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	err = ValidateStruct(*Login)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	Token, err := api.LoginOperator(Login)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return utility.FiberError(context, http.StatusOK, Token)
}

func insertDataQrCodeOps(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	var data = new(structure.InsertDataQrCode)
	if err := context.BodyParser(data); err != nil {
		return utility.FiberError(context, http.StatusBadRequest, "ส่งชนิดของข้อมูลมาผิด")
	}
	err := ValidateStruct(*data)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	//OperatorId , err := getOperatorId(context)
	//
	//if err != nil {
	//	return utility.FiberError(context, http.StatusBadRequest, err.Error())
	//}
	//err = api.CheckAccountOperator(OperatorId,data.OwnerId)
	//if err != nil {
	//	return utility.FiberError(context, http.StatusBadRequest, err.Error())
	//}
	err = api.InsertDataQrCode(data)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return utility.FiberSuccess(context, http.StatusOK, "บันทึกข้อมูลสำเร็จ")
}

func updateDataQrCodeOps(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	var data = new(structure.UpdateDataQrCode)
	if err := context.BodyParser(data); err != nil {
		return utility.FiberError(context, http.StatusBadRequest, "ส่งชนิดของข้อมูลมาผิด")
	}
	err := ValidateStruct(*data)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	//OperatorId , err := getOperatorId(context)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	//err = api.CheckAccountOperator(OperatorId,data.OwnerId)
	//if err != nil {
	//	return utility.FiberError(context, http.StatusBadRequest, err.Error())
	//}
	err = api.UpdateDataQrCode(data)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return utility.FiberSuccess(context, http.StatusOK, "บันทึกข้อมูลสำเร็จ")
}
