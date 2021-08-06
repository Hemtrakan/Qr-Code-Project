package structure

type GenerateQrCode struct {
	QrCode     string   `json:"qr_code"`
	OwnersId   uint     `json:"owners_id" validate:"required"`
	TeamPageId uint     `json:"team_page_id" validate:"required"`
	Location   Location `json:"location"`
}

type GenQr struct {
	FileZip string `json:"file_zip"`
	Filename []string `json:"filename"`
}

type GetQrCodeImage struct {
	FileName string `json:"file_name"`
}

type Location struct {
	ID          uint   `json:"id"`
	Country     string `json:"country"`
	Address     string `json:"address"`
	SubDistrict string `json:"sub_district"`
	District    string `json:"district"`
	Province    string `json:"province"`
	Zipcode     string `json:"zipcode"`
	XCoordinate string `json:"x_coordinate"`
	YCoordinate string `json:"y_coordinate"`
}
