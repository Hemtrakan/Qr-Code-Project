package present

import (
	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt"
	"net/http"
	"qrcode/access/constant"
	"qrcode/control"
	"qrcode/present/structure"
	"qrcode/utility"
	"strconv"
)


func registerOwner(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	owner := new(structure.RegisterOwners)
	err := context.BodyParser(owner)
	if err != nil{
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}

	err = validateStruct(*owner)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}

	err = api.RegisterOwner(owner)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return utility.FiberError(context, http.StatusOK, "สมัครสมาชิกสำเร็จ")
}

func registerOperator(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	operator := new(structure.RegisterOperator)
	err := context.BodyParser(operator)
	if err != nil{
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	err = validateStruct(*operator)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	err = api.RegisterOperator(operator)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return utility.FiberError(context, http.StatusOK, "สมัครสมาชิกสำเร็จ")
}

func login(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	Login := new(structure.Login)
	err := context.BodyParser(Login)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	err = validateStruct(*Login)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	Token, err := api.Login(Login)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return utility.FiberError(context, http.StatusOK, Token)
}


func LoginAdmin(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	Login := new(structure.Login)
	err := context.BodyParser(Login)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	err = validateStruct(*Login)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	Token, err := api.LoginAdmin(Login)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return utility.FiberError(context, http.StatusOK, Token)
}


func getAccount(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	user := context.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	var UserId = claims["id"].(float64)
	var id = int(UserId)
	responses, err := api.GetAccount(id)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return context.JSON(responses)
}

func getAllAccountOwner(context *fiber.Ctx)error  {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)

	//searchAccountOwner := new(structure.SearchAccountOwner)
	//if err := context.QueryParser(searchAccountOwner); err != nil {
	//	return utility.FiberError(context, http.StatusBadRequest, err.Error())
	//}
	responses, err := api.GetAllAccountOwner()
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return context.JSON(responses)
}

func getSubOwner(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	id := context.Params("id")
	OwnerId, err := strconv.Atoi(id)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest,"กรอกได้แต่ตัวเลขเท่านั้น")
	}
	response , err := api.GetSubOwner(OwnerId)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest,err.Error())
	}
	return context.JSON(response)
}

func getAllAccountOperator(context *fiber.Ctx)error  {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	//searchAccountOperator := new(structure.SearchAccountOperator)
	//if err := context.QueryParser(searchAccountOperator); err != nil {
	//	return utility.FiberError(context, http.StatusBadRequest, err.Error())
	//}
	responses, err := api.GetAllAccountOperator()
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return context.JSON(responses)
}

func getOwnerByIdOps(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	id := context.Params("id")
	OperatorId, err := strconv.Atoi(id)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest,"กรอกได้แต่ตัวเลขเท่านั้น")
	}
	response , err := api.GetOwnerByIdOps(OperatorId)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest,err.Error())
	}
	return context.JSON(response)
}

func getAccountById(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	UserId := context.Params("id")
	id, err := strconv.Atoi(UserId)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest,"กรอกได้แต่ตัวเลขเท่านั้น")
	}
	responses, err := api.GetAccount(id)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return context.JSON(responses)
}

func updateProfile(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	Account := new(structure.UpdateProFile)
	err := context.BodyParser(Account)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	err = validateStruct(*Account)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	UserId := context.Params("id")
	id, err := strconv.Atoi(UserId)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest,"กรอกได้แต่ตัวเลขเท่านั้น")
	}
	err = api.UpdateProfile(uint(id),Account)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return utility.FiberError(context, http.StatusOK, "แก้ไขข้อมูลผู้ใช้งานสำเร็จ")
}

func changePassword(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	ChangePassword := new(structure.ChangePassword)
	err := context.BodyParser(ChangePassword)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	err = validateStruct(*ChangePassword)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	UserId := context.Params("id")
	id, err := strconv.Atoi(UserId)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest,"กรอกได้แต่ตัวเลขเท่านั้น")
	}
	err = api.ChangePassword(uint(id),ChangePassword)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return utility.FiberError(context, http.StatusOK, "เปลี่ยนรหัสผ่านสำเร็จ")
}

func deleteAccount(context *fiber.Ctx) error{
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	UserId := context.Params("id")
	id, err := strconv.Atoi(UserId)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest,"กรอกได้แต่ตัวเลขเท่านั้น")
	}
	err = api.DeleteAccount(uint(id))
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return utility.FiberError(context, http.StatusOK, "ลบข้อมูลผู้ใช้สำเร็จ")
}


func admin(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	err := api.RegisterAdmin()
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, "สมัครไปแล้ว")
	}
	return utility.FiberError(context, http.StatusOK, "สำหรับ UserAdmin")
}

// AuthError Auth
func AuthError(c *fiber.Ctx, e error) error {
	c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message": "Unauthorized",
	})
	return nil
}
