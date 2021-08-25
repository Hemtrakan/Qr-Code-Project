package structure

import (
	"github.com/gofrs/uuid"
	"gorm.io/datatypes"
	"time"
)

type FileZipByTemplateName struct {
	OwnerId      uint   `json:"owner_id" validate:"required"`
	TemplateName string `json:"template_name" validate:"required"`
	FileZip      string `json:"file_zip" validate:"required"`
}

type FileZip struct {
	OwnerId  int      `json:"owner_id" validate:"required"`
	FileZip  string   `json:"file_zip" validate:"required"`
	QrCodeId []string `json:"qr_code_id" validate:"required"`
}

type GetQrCode struct {
	OwnerId       uint      `json:"owner_id"`
	OwnerUsername string    `json:"owner_username"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	TemplateName  string    `json:"template_name"`
	QrCodeId      string    `json:"qr_code_id"`
	CodeName      string    `json:"code_name"`
	URL           string    `json:"url"`
	Status        bool      `json:"status"`
}

type GetDataQrCode struct {
	QrCodeId     string         `json:"qr_code_id"`
	Info         datatypes.JSON `json:"info"`
	Ops          datatypes.JSON `json:"ops"`
	HistoryInfo  []GetHistory   `json:"history_info"`
	OwnerId      int            `json:"owner_id"`
	TemplateName string         `json:"template_name"`
	CodeName     string         `json:"code_name"`
}

type GetHistory struct {
	HistoryInfo datatypes.JSON
	UserId      uint
	UpdatedAt   time.Time
}

type ArrayFileName struct {
	FileName string `json:"file_name"`
}

type GetQrCodeImage struct {
	FileName string `json:"file_name"`
}

type GenQrCode struct {
	OwnerId      uint   `json:"owner_id" validate:"required"`
	CodeName     string `json:"code_name" validate:"required"`
	TemplateName string `json:"template_name"`
	Amount       int    `json:"amount" validate:"required"`
}

type InsertDataQrCode struct {
	OwnerId      uint        `json:"owner_id" validate:"required"`
	QrCodeId     uuid.UUID   `json:"qr_code_id" validate:"required"`
	TemplateName string      `json:"template_name" validate:"required"`
	Info         interface{} `json:"info" validate:"required"`
}
type DelQrCode struct {
	QrCodeId []string `json:"qr_code_id" validate:"required"`
}

type StatusQrCode struct {
	QrCodeId uuid.UUID `json:"qr_code_id"   validate:"required"`
	Status   bool      `json:"status"`
}
