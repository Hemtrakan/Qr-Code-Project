package constant

import (
	"errors"
)

type Templates string

const (
	serviceWashingMachine Templates = "ServiceWashingMachine"
)

var Template = []Templates{
	serviceWashingMachine,
}

func (userRole Templates) Templates() (result *string, Errors error) {
	switch userRole {
	case serviceWashingMachine:
		TemplatesName := "ServiceWashingMachine"
		result = &TemplatesName
	default:
		Errors = errors.New("unimplemented")
	}
	return
}
