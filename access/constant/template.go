package constant

import (
	"errors"
)

type Templates string

const (
	serviceWashingMachine Templates = "ServiceWashingMachine"
	computer Templates = "Computer"
)

var Template = []Templates{
	serviceWashingMachine,
	computer,
}

func (Templates Templates) Templates() (result *string, Errors error) {
	switch Templates {
	case serviceWashingMachine:
		TemplatesName := "ServiceWashingMachine"
		result = &TemplatesName
	case computer:
		TemplatesName := "Computer"
		result = &TemplatesName
	default:
		Errors = errors.New("unimplemented")
	}
	return
}
