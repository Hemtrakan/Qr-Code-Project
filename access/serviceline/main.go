package serviceline

import "qrcode/environment"

type FactoryServiceLine interface {
	ServiceLine() (Error error)
}

func ServiceLine(prop environment.Properties) FactoryServiceLine {
	return APIServiceLine(prop)
}