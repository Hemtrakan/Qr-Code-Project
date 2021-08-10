package structure

type FileZipByTemplateName struct {
	OwnerId      uint   `json:"owner_id"`
	TemplateName string `json:"template_name"`
	FileZip      string `json:"file_zip"`
}

type FileZip struct {
	OwnerId  int      `json:"owner_id"`
	FileZip  string   `json:"file_zip"`
	FileName []FileNames `json:"file_name"`
}

type FileNames struct {
	QrCodeId string   `json:"qr_code_id"`
	Filename string `json:"filename"`
}

type GetQrCode struct {
	OwnerId               uint   `json:"owner_id"`
	TemplateName          string `json:"template_name"`
	QrCodeId              string `json:"qr_code_id"`
	ThisQrCodeIsGenerated bool   `json:"this_qr_code_is_generated"`
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
