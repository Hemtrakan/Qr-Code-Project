package serviceline

import "qrcode/environment"


type serviceLine struct {
	prop environment.Properties
}

func (s serviceLine) ServiceLine() (Error error) {
	panic("implement me")
}

func APIServiceLine(prop environment.Properties) serviceLine {
	return serviceLine{
		prop: prop,
	}
}
