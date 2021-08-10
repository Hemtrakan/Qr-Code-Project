package structure
//
//import (
//	"github.com/gofrs/uuid"
//	"time"
//)
//
//type Template struct {
//	OwnerId      int         `json:"owner_id"`
//	TemplateName string      `json:"template_name"`
//	QrCodeTypeId string      `json:"qr_code_type_id"`
//	info         interface{} `json:"info"`
//	ops          interface{} `json:"ops"`
//}
//
//type computerInfo struct {
//	history  []computer `json:"history"`
//	computer `json:"computer"`
//}
//
//type computerOps struct {
//	OpsComputer []string `json:"ops_computer"`
//}
//
//type computer struct {
//	Case        string `json:"case"`
//	PowerSupply string `json:"power_supply"`
//	MainBoar    string `json:"main_boar"`
//	CPU         string `json:"cpu"`
//	RAM         string `json:"ram"`
//	GraphicCard string `json:"graphic_card"`
//	HardDisk    string `json:"hard_disk"`
//}
//
//type GetAllTeamPage struct {
//	Id           uint      `json:"id" query:"id"`
//	TeamPageName string    `json:"team_page_name" query:"team_page_name"`
//	TeamPageFile string    `json:"team_page_file" query:"team_page_file"`
//	TeamPageId   uuid.UUID `json:"team_page_id" query:"team_page_id"`
//	QrCodeType   string    `json:"qr_code_type"`
//}
//
//type GetByIdTeamPage struct {
//	Id         uint        `json:"id" query:"id"`
//	TeamPageId string      `json:"team_page_id" query:"team_page_id"`
//	Data       interface{} `json:"data"`
//}
//
//type ResGetByIdTeamPage struct {
//	TeamPageId uint        `json:"team_page_id" query:"team_page_id"`
//	Data       interface{} `json:"data"`
//}
//
//type GetAllLogTeamPage struct {
//	ID        uint        `json:"id"`
//	UpdatedAt time.Time   `json:"updated_at"`
//	LogData   interface{} `json:"log_data"`
//}
