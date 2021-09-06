package constant

import (
	"errors"
)

type Templates string

const (
	Computer              Templates = "computer"
	OfficeEquipment       Templates = "officeEquipment"
)

var Template = []Templates{
	Computer,
	OfficeEquipment,
}

func (Templates Templates) Templates() (result *string, Errors error) {
	switch Templates {
	//case ServiceWashingMachine:
	//	TemplatesName := "serviceWashingMachine"
	//	result = &TemplatesName
	case Computer:
		TemplatesName := "computer"
		result = &TemplatesName
	case OfficeEquipment:
		TemplatesName := "officeEquipment"
		result = &TemplatesName
	default:
		Errors = errors.New("unimplemented")
	}
	return
}
