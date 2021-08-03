package environment

import (
	"github.com/Netflix/go-env"
	"log"
)

type Flavor string

const (
	Develop    Flavor = "DEVELOP"
	Devspace   Flavor = "DEVSPACE"
	Production Flavor = "PRODUCTION"
)

type Properties struct {
	// -- core
	Flavor Flavor `env:"FLAVOR,default=DEVELOP"`
	// --

	// -- Gorm
	//GormHost string `env:"GORM_HOST,default=rdbms"`
	//GormHost string `env:"GORM_HOST,default=localhost"`
	GormHost string `env:"GORM_HOST,default=qrcode-rdbms"`
	GormPort string `env:"GORM_PORT,default=5432"`
	GormName string `env:"GORM_NAME,default=qr_code"`
	GormUser string `env:"GORM_USER,default=postgres"`
	GormPass string `env:"GORM_PASS,default=pgpassword"`
	// --

	// -- Minio
	MinioBucketName string `env:"MINIO_BUCKET_NAME,default=tepa-implement"`
	MinioEndpoint   string `env:"MINIO_ENDPOINT,default=sgp1.digitaloceanspaces.com"`
	MinioKeyID      string `env:"MINIO_KEY_ID,default=BHARQKX7CH35NGKFRJ34"`
	MinioSecretKey  string `env:"MINIO_SECRET_KEY,default=2wGM61EEeo3sjZyL16esk/YXvV5xrWNTe5t+1X/Hp28"`
	// --

	//// -- authentication
	//SecurityKey string `env:"SECURITY_KEY,default=t-T-DEV Co., Ltd."`
	//// --

	//// --
	URLFront string `env:"URL_FRONT,default=http://localhost:12000/viewdata/"`

	//GRPC
	//ServerGRPC string `env:"SERVER_GRPC,default=inventory-api:9090"`
	// --

}

func Build() *Properties {
	var prop Properties
	if _, err := env.UnmarshalFromEnviron(&prop); err != nil {
		log.Panic(err)
	}
	return &prop
}
