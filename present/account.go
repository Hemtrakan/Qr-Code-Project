package present

import (
	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt"
	"net/http"
	"qrcode/access/constant"
	"qrcode/control"
	"qrcode/present/structure"
	"qrcode/utility"
)


func registerOwner(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	owner := new(structure.Owners)
	err := context.BodyParser(owner)
	if err != nil{
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	err = api.RegisterOwner(owner)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, "this user already exists")
	}
	return utility.FiberError(context, http.StatusOK, "success")
}

func registerOperator(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	operator := new(structure.Operator)
	err := context.BodyParser(operator)
	if err != nil{
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	err = api.RegisterOperator(operator)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, "this user already exists")
	}
	return utility.FiberError(context, http.StatusOK, "success")
}

func login(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	Login := new(structure.Login)
	err := context.BodyParser(Login)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	Token, err := api.Login(Login)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, "record not found")
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
	responses, err := api.GetAllAccountOwner()
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return context.JSON(responses)
}

func getAllAccountOperator(context *fiber.Ctx)error  {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	responses, err := api.GetAllAccountOperator()
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return context.JSON(responses)
}

func AuthError(c *fiber.Ctx, e error) error {
	c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message": "Unauthorized",

	})
	return nil
}