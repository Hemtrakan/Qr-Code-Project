package file

import (
	"mime/multipart"
	"net/url"
	"os"
	"qrcode/access/constant"
	"qrcode/environment"
	"sync"
)


var (
	factory FactoryInterface
	once    sync.Once
)


type FactoryInterface interface {
	UploadFile(category constant.CategoryFile, fileHeader multipart.FileHeader) (*string, error)
	DownloadFile(category constant.CategoryFile, fileName string) (*os.File, error)
	GetFileWithURL(category constant.CategoryFile, fileName string) (*url.URL, error)
}

func Create(env *environment.Properties) FactoryInterface {
	once.Do(func() {
		factory = MinioInstance(env)
	})
	return factory
}
