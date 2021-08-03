package control

import (
	"mime/multipart"
	"net/url"
	"qrcode/access/constant"
)

func (ctrl *APIControl) UploadFile(src multipart.FileHeader, fileType constant.CategoryFile) (*string, error) {
	return ctrl.access.FILE.UploadFile(fileType, src)
}

func (ctrl *APIControl) GetUrlFile(srcName string, fileType constant.CategoryFile) (*url.URL, error) {
	return ctrl.access.FILE.GetFileWithURL(fileType, srcName)
}


