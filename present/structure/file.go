package structure

type FileRequest struct {
	File     string `json:"file"`
	Data string `json:"data"`
}

type FileResponse struct {
	Data string `json:"data"`
}


type FileReqCreate struct {
	FileName []FileName `json:"file_name"`
}

type FileName struct {
	Name string `json:"name"`
}