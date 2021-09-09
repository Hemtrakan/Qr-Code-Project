package environment

import (
	"github.com/Netflix/go-env"
	"log"
)

type Flavor string
type URL string

const (
	Develop    Flavor = "DEVELOP"
	Devspace   Flavor = "DEVSPACE"
	Production Flavor = "PRODUCTION"
)

const(
	//URLFront URL = "https://9013-1-47-146-211.ngrok.io"
	URLFront URL = "http://192.168.1.104:12000/qr/"
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
	
	//// -- authentication
	//SecurityKey string `env:"SECURITY_KEY,default=t-T-DEV Co., Ltd."`
	//// --

	//// -- URL
	URLFront string `env:"URL_FRONT,default=http://localhost:12000/viewdata/"`
	//URLFront URL `env:"URL_FRONT,default=https://4eb2-1-47-146-211.ngrok.io/viewdata/"`
	//URLQRCode string `env:"URL_QR_CODE,default=https://liff.line.me/1656385614-YE6rXz2M/api/qr/"`
	URLQRCode string `env:"URL_QR_CODE,default=https://192.168.1.104:8080/api/qr/"`

	//// -- ServiceLine
	Authorization string `env:"AUTHORIZATION,default=Bearer gkFHAAtmlfClxm0//s233eQb6eTaksrvKzJ+p171IpINsonSX2JV0LMlnKTxTRbdPAc/1fU27N/77/nv+vjffmBJBCOUKbYME0fZ3HOLlM7rlfnF8uddV4JMjjXAuRVN/9YnU4XjDnp2vgrjVAoTtQdB04t89/1O/w1cDnyilFU="`

}

func Build() *Properties {
	var prop Properties
	if _, err := env.UnmarshalFromEnviron(&prop); err != nil {
		log.Panic(err)
	}
	return &prop
}
