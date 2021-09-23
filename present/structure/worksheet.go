package structure

import (
	"github.com/gofrs/uuid"
	"qrcode/access/constant"
	"time"
)

type TypeWorksheet struct {
	Data []TypeWorksheets `json:"data"`
}
type TypeWorksheets struct {
	TypeWorksheet constant.Worksheets `json:"type_worksheet"`
}

type GetWorksheet struct {
	Info      interface{} `json:"info,omitempty"`
	Worksheet []Worksheet `json:"worksheet,omitempty"`
}

type Worksheet struct {
	ID uint `json:"id,omitempty"`
	// การเพิ่ม option ในส่วน QrCode ตอนที่มีการ assign งาน
	// มาให้สามารถกำหนดได้ว่าเป็นรูปแบบที่ (Owner) สามารถระบุให้ช่างได้เลย
	Option          bool              `json:"option"`
	QrCodeID        uuid.UUID         `json:"qr_code_id,omitempty"`
	Info            interface{}       `json:"info,omitempty"`
	Text            string            `json:"text,omitempty"`
	Type            string            `json:"type,omitempty"`
	Ops             *string           `json:"ops,omitempty"`
	OwnerId         uint              `json:"owner_id,omitempty"`
	StatusWorksheet []StatusWorksheet `json:"status_worksheet"`
}

type UpdateOption struct {
	Option       bool   `json:"option" validate:"required"`
}

type StatusWorksheet struct {
	Status     string      `json:"status,omitempty"`
	UpdateAt   *time.Time  `json:"update_at,omitempty"`
	Text       *string     `json:"text,omitempty"`
	Equipments []Equipment `json:"equipments,omitempty"`
}

type Equipment struct {
	NameEquipment string `json:"name_equipment,omitempty"`
}

type ReportID struct {
	ReportID   uint   `json:"report_id" query:"report_id"`
	LineUserId string `json:"line_user_id" query:"line_user_id"`
}

type InsertWorksheet struct {
	QrCodeID uuid.UUID `json:"qr_code_id,omitempty" validate:"required"`
	Text     string    `json:"text,omitempty" validate:"required"`
	Type     string    `json:"type,omitempty" validate:"required"`
}

type UpdateWorksheet struct {
	LineUserId string      `json:"line_user_id"`
	Text       string      `json:"text,omitempty" validate:"required"`
	Equipments []Equipment `json:"equipments,omitempty"`
}
