package present

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"net/http"
	"qrcode/access/constant"
	"qrcode/control"
	"qrcode/present/structure"
	"qrcode/utility"
	"strconv"
)

func getOwnerId(context *fiber.Ctx) (id uint, Error error) {
	user := context.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	var UserId = claims["id"].(float64)
	id = uint(UserId)
	return
}

func LoginOwner(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	Login := new(structure.LoginOwner)
	err := context.BodyParser(Login)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	err = ValidateStruct(*Login)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	Token, err := api.LoginOwner(Login)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return utility.FiberError(context, http.StatusOK, Token)
}

func registerOperatorOwner(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	operator := new(structure.RegisterOperator)
	err := context.BodyParser(operator)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	err = ValidateStruct(*operator)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	id, err := getOwnerId(context)
	err = api.RegisterOperatorOwner(operator, &id)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return utility.FiberSuccess(context, http.StatusOK, "สมัครสมาชิกสำเร็จ")
}

func getOperator(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	id, err := getOwnerId(context)
	if err != nil {
		return utility.FiberError(context,http.StatusBadRequest, err.Error())
	}
	res ,err := api.GetOperator(id)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return context.Status(http.StatusOK).JSON(res)
}

func getOperatorLine(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	id, err := getOwnerId(context)
	if err != nil {
		return utility.FiberError(context,http.StatusBadRequest, err.Error())
	}
	res ,err := api.GetOperatorLine(id)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return context.Status(http.StatusOK).JSON(res)
}

func getOperatorById(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	OwnerId, err := getOwnerId(context)
	if err != nil {
		return utility.FiberError(context,http.StatusBadRequest, err.Error())
	}
	UserId := context.Params("id")
	OperatorId, err := strconv.Atoi(UserId)
	res ,err := api.GetOperatorById(OperatorId,OwnerId)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return context.Status(http.StatusOK).JSON(res)
}



func deleteAccountOperator(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	OwnerId, err := getOwnerId(context)
	UserId := context.Params("id")
	OperatorId, err := strconv.Atoi(UserId)
	err = api.DeleteAccountOperator(OwnerId, OperatorId)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return utility.FiberSuccess(context, http.StatusOK, "ลบข้อมูลของช่างซ่อมสำเร็จ")
}

func ChangePasswordOperatorByOwner(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	OwnerId, err := getOwnerId(context)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	changePasswordOwner := new(structure.ChangePasswordOperator)
	err = context.BodyParser(changePasswordOwner)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	err = ValidateStruct(*changePasswordOwner)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	err = api.ChangePasswordOperator(OwnerId, *changePasswordOwner)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return utility.FiberSuccess(context, http.StatusOK, "เปลี่ยนรหัสผ่านสำเร็จ")
}


func ChangePasswordOwner(context *fiber.Ctx) error {
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

func updateStatusQrCodeOwner(context *fiber.Ctx) error {
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
	ownerId, err := getOwnerId(context)
	err = api.UpdateStatusQrCodeOwner(ownerId,QrCodeId, *QrCode)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return utility.FiberSuccess(context, http.StatusOK, "เปลี่ยนสถานะ QrCode สำเร็จ")
}

func getQrCodeOwnerById(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	ownerId, err := getOwnerId(context)
	res, err := api.GetQrCodeById(int(ownerId))
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return context.JSON(res)
}

