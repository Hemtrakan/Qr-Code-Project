package rdbms

import (
	rdbmsstructure "qrcode/access/rdbms/structure"
	"qrcode/environment"
	"sync"
)

var (
	factory FactoryInterface
	once    sync.Once
)

type FactoryInterface interface {
	// Customer
	GetAllAccountOwner() (response []rdbmsstructure.Account,Error error)
	GetAllAccountOperator() (response []rdbmsstructure.Account,Error error)
	UpdateProfile(Account rdbmsstructure.Account) (Error error)
	DeleteAccount(id int) (Error error)

	//  QRCode
	CreateQrCode(QrCode rdbmsstructure.QrCode) (Error error)
	GetQrCode(OwnerId uint,templateName string) (response []rdbmsstructure.QrCode,Error error)
	GetQrCodeById(OwnerId int) (response []rdbmsstructure.QrCode,Error error)
	UpdateQrCode(QrCode rdbmsstructure.QrCode) (Error error)
	UpdateQrCodeById(QrCode rdbmsstructure.QrCode) (Error error)

	// Acconut
	Register(Account rdbmsstructure.Account) (Error error)
	Login(login rdbmsstructure.Account) (response rdbmsstructure.Account,Error error)
	GetAccount(id int) (response rdbmsstructure.Account,Error error)

}

func Create(env *environment.Properties) FactoryInterface {
	once.Do(func() {
		factory = gormInstance(env)
	})
	return factory
}