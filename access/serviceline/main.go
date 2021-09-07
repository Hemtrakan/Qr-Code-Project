package serviceline

import "qrcode/environment"

type FactoryServiceLine interface {
	LinkRichMenuToUser(lineId string) (Error error)
}

func ServiceLine(prop *environment.Properties) FactoryServiceLine {
	return APIServiceLine(*prop)
}