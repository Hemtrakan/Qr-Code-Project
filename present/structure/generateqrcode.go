package structure

type FileZipByTemplateName struct {
	OwnerId      uint   `json:"owner_id"`
	TemplateName string `json:"template_name"`
	FileZip      string `json:"file_zip"`
}

type FileZip struct {
	OwnerId  int      `json:"owner_id"`
	FileZip  string   `json:"file_zip"`
	QrCodeId []string `json:"qr_code_id"`
}

type GetQrCode struct {
	OwnerId      uint   `json:"owner_id"`
	TemplateName string `json:"template_name"`
	QrCodeId     string `json:"qr_code_id"`
	CodeName     string `json:"code_name"`
}

type GetDataQrCode struct {
	QrCodeId    string `json:"qr_code_id"`
	Info        string `json:"info"`
	Ops         string `json:"ops"`
	HistoryInfo string `json:"history_info"`
	OwnerId     int    `json:"owner_id"`
}

type ArrayFileName struct {
	FileName string `json:"file_name"`
}

type GetQrCodeImage struct {
	FileName string `json:"file_name"`
}

type GenQrCode struct {
	OwnerId      uint   `json:"owner_id"`
	CodeName     string `json:"code_name"`
	TemplateName string `json:"template_name"`
	Amount       int    `json:"amount"`
}

type DelQrCode struct {
	QrCodeId []string `json:"qr_code_id"`
}
