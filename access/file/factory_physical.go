package file

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"qrcode/access/constant"
	"qrcode/environment"
	"strings"
	"time"
)

type PhysicalFactory struct {
	env *environment.Properties
}

func physicalInstance(env *environment.Properties) PhysicalFactory {
	return PhysicalFactory{env: env}
}

func (instance PhysicalFactory) UploadFile(category constant.CategoryFile, file multipart.FileHeader) (*string, error) {

	loc, _ := time.LoadLocation("Asia/Bangkok")
	fileNamePrefix := time.Now().In(loc).Format("20060102_150405")
	fileNameExt := strings.Split(file.Header.Get("Content-Type"), "/")[len(strings.Split(file.Header.Get("Content-Type"), "/"))-1]
	fileName := fmt.Sprintf("%s.%s", fileNamePrefix, fileNameExt)

	src, err := file.Open()
	if err != nil {
		return nil, err
	}

	dst, err := os.Create(fmt.Sprintf("assets/%s/%s", category, fileName))
	if err != nil {
		return nil, err
	}

	// Copy
	if _, err = io.Copy(dst, src); err == nil {
		if err := src.Close(); err != nil {
			return nil, err
		}

		if err := dst.Close(); err != nil {
			return nil, err
		}
		return &fileName, nil
	} else {
		return nil, err
	}
}

func (instance PhysicalFactory) DownloadFile(category constant.CategoryFile, fileName string) (*os.File, error) {
	file, err := os.Open(fmt.Sprintf("assets/%s/%s", category, fileName))
	if err != nil {
		return nil, err
	}
	return file, nil
}

