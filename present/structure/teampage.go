package structure

import (
	"github.com/gofrs/uuid"
	"time"
)

type TeamPage struct {
	OwnerId      int         `json:"owner_id"`
	TeamPageName string      `json:"team_page_name"`
	QrCodeType   string      `json:"qr_code_type"`
	Info         interface{} `json:"info"`
	Ops          interface{} `json:"ops"`
}

type GetAllTeamPage struct {
	Id           uint      `json:"id" query:"id"`
	TeamPageName string    `json:"team_page_name" query:"team_page_name"`
	TeamPageFile string    `json:"team_page_file" query:"team_page_file"`
	TeamPageId   uuid.UUID `json:"team_page_id" query:"team_page_id"`
	QrCodeType   string    `json:"qr_code_type"`
}

type GetByIdTeamPage struct {
	Id         uint        `json:"id" query:"id"`
	TeamPageId string      `json:"team_page_id" query:"team_page_id"`
	Data       interface{} `json:"data"`
}

type ResGetByIdTeamPage struct {
	TeamPageId uint        `json:"team_page_id" query:"team_page_id"`
	Data       interface{} `json:"data"`
}

type GetAllLogTeamPage struct {
	ID        uint        `json:"id"`
	UpdatedAt time.Time   `json:"updated_at"`
	LogData   interface{} `json:"log_data"`
}
