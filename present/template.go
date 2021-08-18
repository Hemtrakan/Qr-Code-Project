package present

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"qrcode/access/constant"
	"qrcode/control"
)

func getTemplate(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	res := api.GetTemplate()
	return context.Status(http.StatusOK).JSON(res)
}

//func getByIdTeamPage(context *fiber.Ctx) error {
//	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
//	contentType := context.Get("Content-Type")
//	id := context.Params("id")
//	res, err := api.GetByIdTeamPage(structure.GetByIdTeamPage{TeamPageId: id})
//	if err != nil {
//		return utility.FiberError(context, http.StatusBadRequest, err.Error())
//	}
//	if contentType == "" {
//		//return context.Render("index", fiberRoute.Map{"res": res.Data}, "layouts/main")
//		url, err := api.GetHtml(id)
//		if err != nil {
//			fmt.Println(url)
//		}
//		context.Status(http.StatusOK).SendString(url)
//	}
//	if contentType == "application/json" {
//		return context.Status(http.StatusOK).JSON(res)
//	}
//	//return context.Status(http.StatusOK).JSON(res)
//	return utility.FiberError(context, http.StatusBadRequest, "record not found")
//}
//
//func getAllTeamPage(context *fiber.Ctx) error {
//	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
//	user := context.Locals("user").(*jwt.Token)
//	claims := user.Claims.(jwt.MapClaims)
//	var UserId = claims["id"].(float64)
//	var id = int(UserId)
//	res, err := api.GetAllTeamPage(id)
//	if err != nil {
//		return utility.FiberError(context, http.StatusBadRequest, err.Error())
//	}
//	return context.Status(http.StatusOK).JSON(res)
//}
//
//
//func getAllTeamPageById(context *fiber.Ctx) error {
//	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
//	userid := context.Params("id")
//	id, err := strconv.Atoi(userid)
//	if err != nil {
//		return utility.FiberError(context, http.StatusBadRequest, err.Error())
//	}
//	res, err := api.GetAllTeamPage(id)
//	if err != nil {
//		return utility.FiberError(context, http.StatusBadRequest, err.Error())
//	}
//	return context.Status(http.StatusOK).JSON(res)
//}
//
//func insertTeamPage(context *fiber.Ctx) error {
//	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
//	TeamPage := new([]structure.Template)
//	err := context.BodyParser(TeamPage)
//	if err != nil {
//		return utility.FiberError(context, http.StatusBadRequest, err.Error())
//	}
//	err = api.InsertTeamPage(TeamPage)
//	if err != nil {
//		return utility.FiberError(context, http.StatusBadRequest, err.Error())
//	}
//	return utility.FiberError(context, http.StatusOK, "succeed")
//}
//
//func updateTeamPage(context *fiber.Ctx) error {
//	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
//	TeamPage := new(structure.Template)
//	err := context.BodyParser(TeamPage)
//	if err != nil {
//		return utility.FiberError(context, http.StatusBadRequest, err.Error())
//	}
//	id := context.Params("id")
//	err = api.UpdateTeamPage(structure.GetByIdTeamPage{TeamPageId: id}, TeamPage)
//	if err != nil {
//		return utility.FiberError(context, http.StatusBadRequest, err.Error())
//	}
//	return utility.FiberError(context, http.StatusOK, "succeed")
//}
//
//func deleteTeamPage(context *fiber.Ctx) error {
//	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
//	id := context.Params("id")
//	err := api.DeleteTeamPage(structure.GetByIdTeamPage{TeamPageId: id})
//	if err != nil {
//		return utility.FiberError(context, http.StatusBadRequest, err.Error())
//	}
//	return utility.FiberError(context, http.StatusOK, "succeed")
//}

//func getAllLogTeamPage(context *fiber.Ctx) error {
//	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
//	id := context.Params("id")
//	teamPageId, err := strconv.Atoi(id)
//	if err != nil {
//		return utility.FiberError(context, http.StatusBadRequest, err.Error())
//	}
//	res, err := api.GetAllLogTeamPage(structure.GetAllLogTeamPage{ID: uint(teamPageId)})
//	if err != nil {
//		return utility.FiberError(context, http.StatusBadRequest, err.Error())
//	}
//	return context.Status(http.StatusOK).JSON(res)
//}
