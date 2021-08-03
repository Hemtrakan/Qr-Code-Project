package file

import (
	"fmt"
	"github.com/minio/minio-go"
	"io"
	"log"
	"mime/multipart"
	"net/url"
	"os"
	"qrcode/access/constant"
	"qrcode/environment"
	"qrcode/utility"
	"strings"
	"time"
)

type MinioFactory struct {
	env    *environment.Properties
	client *minio.Client
}

func MinioInstance(env *environment.Properties) MinioFactory {
	client, err := minio.NewV4(env.MinioEndpoint, env.MinioKeyID, env.MinioSecretKey, true)
	if err != nil {
		log.Panic(err)
	}
	return MinioFactory{client: client, env: env}
}

func (instance MinioFactory) UploadFileWithOSFile(category constant.CategoryFile, file *os.File) (*string, error) {
	loc, _ := time.LoadLocation("Asia/Bangkok")
	fileNamePrefix := time.Now().In(loc).Format("20060102_150405")
	fileNameExt := strings.Split(file.Name(), ".")[len(strings.Split(file.Name(), "."))-1]
	fileName := fmt.Sprintf("%s.%s", fileNamePrefix, fileNameExt)
	fmt.Println(fileName)
	path, ext := utility.MinioCreateObjectName("ORGANIZATION", "ORGS_ID", fmt.Sprintf("%s/%s/%s", "INVENTORY", category, fileName)) // todo edit orgs_id เป็น id ของ แต่ละบริษัท

	stat, err := file.Stat()
	if err != nil {
		log.Fatalln(err)
	}

	_, err = instance.client.PutObject(instance.env.MinioBucketName, path, file, stat.Size(), minio.PutObjectOptions{ContentType: fmt.Sprintf("image/%s", ext)})
	if err != nil {
		log.Fatalln(err)
	}
	return &fileName, nil
}

func (instance MinioFactory) UploadFile(category constant.CategoryFile, file multipart.FileHeader) (*string, error) {
	loc, _ := time.LoadLocation("Asia/Bangkok")
	fileNamePrefix := time.Now().In(loc).Format("20060102_150405")
	fileNameExt := strings.Split(file.Header.Get("Content-Type"), "/")[len(strings.Split(file.Header.Get("Content-Type"), "/"))-1]
	fileName := fmt.Sprintf("%s.%s", fileNamePrefix, fileNameExt)

	path, ext := utility.MinioCreateObjectName("ORGANIZATION", "ORGS_ID", fmt.Sprintf("%s/%s/%s", "INVENTORY", category, fileName))

	src, err := file.Open()
	if err != nil {
		return nil, err
	}

	_, err = instance.client.PutObject(instance.env.MinioBucketName, path, src, file.Size, minio.PutObjectOptions{ContentType: fmt.Sprintf("image/%s", ext)})
	if err != nil {
		fmt.Println(err)

		log.Fatalln(err)
	}

	return &fileName, nil
}

func (instance MinioFactory) GetFileWithURL(category constant.CategoryFile, fileName string) (*url.URL, error) {
	path, _ := utility.MinioCreateObjectName("ORGANIZATION", "ORGS_ID", fmt.Sprintf("%s/%s/%s", "INVENTORY", category, fileName))
	if _, err := instance.client.GetObject(instance.env.MinioBucketName, path, minio.GetObjectOptions{}); err != nil {
		return nil, err
	} else {
		return instance.client.PresignedGetObject(instance.env.MinioBucketName, path, 5*time.Minute, nil)
	}
}

func (instance MinioFactory) DownloadFile(category constant.CategoryFile, fileName string) (*os.File, error) {

	path, _ := utility.MinioCreateObjectName("ORGANIZATION", "ORGS_ID", fmt.Sprintf("%s/%s/%s", "INVENTORY", category, fileName))
	object, err := instance.client.GetObject(instance.env.MinioBucketName, path, minio.GetObjectOptions{})
	if err != nil {
		log.Fatalln(err)
	}

	localFile, err := os.Create(fmt.Sprintf("tmp/%s", fileName))
	if err != nil {
		log.Fatalln(err)
	}

	if _, err = io.Copy(localFile, object); err != nil {
		log.Fatalln(err)
	}

	localFile, err = os.Open(fmt.Sprintf("/tmp/%s", fileName))
	if err != nil {
		log.Fatalln(err)
	}

	return localFile, nil
}
