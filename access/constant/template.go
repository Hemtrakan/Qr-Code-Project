package constant

import (
	"errors"
)

type Templates string

const (
	ServiceWashingMachine Templates = "serviceWashingMachine"
	Computer              Templates = "computer"
)

var Template = []Templates{
	ServiceWashingMachine,
	Computer,
}

func (Templates Templates) Templates() (result *string, Errors error) {
	switch Templates {
	case ServiceWashingMachine:
		TemplatesName := "serviceWashingMachine"
		result = &TemplatesName
	case Computer:
		TemplatesName := "computer"
		result = &TemplatesName
	default:
		Errors = errors.New("unimplemented")
	}
	return
}


