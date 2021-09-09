package present

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"qrcode/access/constant"
	"qrcode/control"
	"qrcode/present/structure"
	"qrcode/utility"
	"strconv"
)

func getTypeWorksheet(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	res := api.GetTypeWorksheet()
	return context.Status(http.StatusOK).JSON(res)
}

func getWorksheet(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	workSheet := new(structure.LineUserId)
	err := context.BodyParser(workSheet)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	err = ValidateStruct(*workSheet)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	res, err := api.GetWorksheet(workSheet)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return context.Status(http.StatusOK).JSON(res)
}

func getWorksheetById(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	report := context.Params("id")
	reportId, err := strconv.Atoi(report)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	workSheet := new(structure.LineUserId)
	err = context.BodyParser(workSheet)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	err = ValidateStruct(*workSheet)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	res, err := api.GetWorksheetById(uint(reportId), workSheet)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return context.Status(http.StatusOK).JSON(res)
}

func worksheet(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	report := context.Params("id")
	reportId, err := strconv.Atoi(report)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	workSheet := new(structure.LineUserId)
	err = context.BodyParser(workSheet)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	err = ValidateStruct(*workSheet)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	err = api.Worksheet(uint(reportId), *workSheet)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return utility.FiberSuccess(context, http.StatusOK, "รับงานสำเร็จ")
}

func insertWorksheet(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	report := new(structure.InsertWorksheet)
	err := context.BodyParser(report)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	err = ValidateStruct(*report)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	err = api.InsertWorksheet(report)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return utility.FiberSuccess(context, http.StatusOK, "รายงานปัญหาสำเร็จ")
}

func getUpdateWorksheet(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	QrCodeId := context.Params("id")
	res, err := api.GetUpdateWorksheet(QrCodeId)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return context.Status(http.StatusOK).JSON(res)
}

func updateWorksheet(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	report := context.Params("id")
	reportId, err := strconv.Atoi(report)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	update := new(structure.UpdateWorksheet)
	err = context.BodyParser(update)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	err = ValidateStruct(*update)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	err = api.UpdateWorksheet(uint(reportId), *update)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return utility.FiberSuccess(context, http.StatusOK, "ดำเนิดการแก้ปัญหาสำเร็จ")
}

func deleteWorksheet(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	report := context.Params("id")
	reportId, err := strconv.Atoi(report)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	update := new(structure.UpdateWorksheet)
	err = context.BodyParser(update)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	err = ValidateStruct(*update)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	err = api.DeleteWorksheet(uint(reportId), *update)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return utility.FiberSuccess(context, http.StatusOK, "ยกเลิกรายการสำเร็จ")
}
