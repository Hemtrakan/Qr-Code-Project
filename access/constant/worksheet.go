package constant

import (
	"errors"
)

type Worksheets string

const (
	WorksheetsStatus1 string = "รอดำเนิดการ"
	WorksheetsStatus2 string = "กำลังดำเนินการ"
	WorksheetsStatus3 string = "ดำเนินการเสร็จสิ้น"
	WorksheetsStatus4 string = "รายการถูกยกเลิก"
)

const (
	type1 Worksheets = "แจ้งซ่อม"
	type2 Worksheets = "แจ้งปัญหา"
)

var Worksheet = []Worksheets{
	type1,
	type2,
}

func (Worksheet Worksheets) TypeWorksheets() (result *string, Errors error) {
	switch Worksheet {
	case type1:
		TemplatesName := "แจ้งซ่อม"
		result = &TemplatesName
	case type2:
		TemplatesName := "แจ้งปัญหา"
		result = &TemplatesName
	default:
		Errors = errors.New("unimplemented")
	}
	return
}
