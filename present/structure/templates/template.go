package templates

import "qrcode/access/constant"

type Templatesdata struct {
	Data []Template `json:"data"`
}
type Template struct {
	TemplatesName constant.Templates `json:"templates_name"`
}