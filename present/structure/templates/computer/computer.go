package computer

import (
	"gorm.io/datatypes"
	"time"
)

type Computer struct {
	Info    Info    `json:"info"`
	History History `json:"history"`
	Ops     Ops     `json:"ops"`
}

type History struct {
	History   datatypes.JSON `json:"History"`
	UserId    uint           `json:"user_id"`
	UpdatedAt time.Time      `json:"updated_at"`
}
