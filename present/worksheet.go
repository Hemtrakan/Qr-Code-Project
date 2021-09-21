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


// Owner

func UpdateOption(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	update := new(structure.UpdateOption)
	err := context.BodyParser(update)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	err = api.UpdateOption(*update)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return utility.FiberSuccess(context, http.StatusOK, "เปลี่ยนสำเร็จ")
}

func ownerGetWorksheet(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	OwnerID , err := getOwnerId(context)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	res, err := api.OwnerGetWorksheet(OwnerID)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return context.Status(http.StatusOK).JSON(res)
}

func ownerGetWorksheetById(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	id := context.Params("id")
	ReportID, err := strconv.Atoi(id)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	res, err := api.OwnerGetWorksheetById(uint(ReportID))
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return context.Status(http.StatusOK).JSON(res)
}

func ownerWorksheet(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	report := context.Params("id")
	reportId, err := strconv.Atoi(report)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	workSheet := new(structure.ReportID)
	err = context.BodyParser(workSheet)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	err = ValidateStruct(*workSheet)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	err = api.OwnerWorksheet(uint(reportId), *workSheet)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return utility.FiberSuccess(context, http.StatusOK, "รับงานสำเร็จ")
}

func ownerGetUpdateWorksheet(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	QrCodeId := context.Params("id")
	res, err := api.OwnerGetUpdateWorksheet(QrCodeId)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return context.Status(http.StatusOK).JSON(res)
}

func ownerUpdateWorksheet(context *fiber.Ctx) error {
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
	err = api.OwnerUpdateWorksheet(uint(reportId), *update)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return utility.FiberSuccess(context, http.StatusOK, "ดำเนิดการแก้ปัญหาสำเร็จ")
}

func ownerDeleteWorksheet(context *fiber.Ctx) error {
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
	err = api.OwnerDeleteWorksheet(uint(reportId), *update)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return utility.FiberSuccess(context, http.StatusOK, "ยกเลิกรายการสำเร็จ")
}

// Ops
func getTypeWorksheet(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	res := api.GetTypeWorksheet()
	return context.Status(http.StatusOK).JSON(res)
}

func getWorksheet(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	lineId := context.Params("id")
	res, err := api.GetWorksheet(lineId)
	if err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	return context.Status(http.StatusOK).JSON(res)
}

func getWorksheetById(context *fiber.Ctx) error {
	api := context.Locals(constant.LocalsKeyControl).(*control.APIControl)
	workSheet := new(structure.ReportID)
	if err := context.QueryParser(workSheet); err != nil {
		return utility.FiberError(context, http.StatusBadRequest, err.Error())
	}
	res, err := api.GetWorksheetById(workSheet)
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
	workSheet := new(structure.ReportID)
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