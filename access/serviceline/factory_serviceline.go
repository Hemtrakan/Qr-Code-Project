package serviceline

import (
	"github.com/go-resty/resty/v2"
	"qrcode/environment"
)

type Properties struct {
	prop environment.Properties
}

func (r Properties) LinkRichMenuToUser(lineId string) (Error error) {
	client := resty.New()
	URL := "https://api.line.me/v2/bot/user/" + lineId + "/richmenu/richmenu-251552ebdf8a56ace4cbd9999f5d4fd0"
	_, err := client.R().
		SetHeader("Authorization", r.prop.Authorization).
		Post(URL)
	if err != nil {
		Error = err
		return
	}
	return
}

func APIServiceLine(prop environment.Properties) Properties {
	return Properties{
		prop: prop,
	}
}
