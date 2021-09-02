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
	GetAllAccountOperatorByOwnerID(OwnerId uint) (response []rdbmsstructure.Account,Error error)
	GetSubOwner(OwnerId int) (response []rdbmsstructure.Account,Error error)
	GetOwnerByIdOps(OperatorId int) (response rdbmsstructure.Account,Error error)
	UpdateProfile(Account rdbmsstructure.Account) (Error error)
	DeleteAccount(id uint) (Error error)
	DeleteAccountByOwner(OwnerId uint,OperatorId int) (Error error)
	GetOperatorById(OperatorId int ,OwnerId uint) (response rdbmsstructure.Account,Error error)

	//  QRCode
	GetDataQrCode(QrCodeUUID string) (response []rdbmsstructure.QrCode,Error error)
	CreateQrCode(QrCode rdbmsstructure.QrCode) (Error error)
	GetQrCode(OwnerId uint,templateName string) (response []rdbmsstructure.QrCode,Error error)
	CheckCode(OwnerId uint,templateName ,Code string) (response rdbmsstructure.QrCode,Error error)
	CountCode(OwnerId uint,templateName ,Code string) (response []rdbmsstructure.QrCode,Error error)
	GetQrCodeByOwnerId(OwnerId int) (response []rdbmsstructure.QrCode,Error error)
	GetQrCodeByQrCodeId(OwnerId int,QrCodeId string) (response rdbmsstructure.QrCode,Error error)
	GetAllQrCode() (response []rdbmsstructure.QrCode,Error error)
	InsertDataQrCodeById(QrCode rdbmsstructure.QrCode) (Error error)
	UpdateHistoryInfoQrCodeById(QrCode rdbmsstructure.HistoryInfo) (Error error)
	UpdateOpsQrCodeById(QrCode rdbmsstructure.Ops) (Error error)
	UpdateQrCodeActive(QrCode rdbmsstructure.QrCode) (Error error)
	DeleteQrCode(QrCodeUUID string) (Error error)


	// Template
	InsertQrCode(QrCodeUUID string,QrCode rdbmsstructure.QrCode) (Error error)

	// Acconut
	Register(Account rdbmsstructure.Account) (Error error)
	Login(login rdbmsstructure.Account) (response rdbmsstructure.Account,Error error)
	GetAccount(id int) (response rdbmsstructure.Account,Error error)


	CheckAccountId(id uint) (response *rdbmsstructure.Account,Error error)
	CheckUserRegister(Username , PhoneNumber ,LineId  string,UserId uint) (response *rdbmsstructure.Account,Error error)


	//TestGetQrData() (response []rdbmsstructure.TestQrCode,Error error)
	//TestInsertQR(Qr rdbmsstructure.TestQrCode) (Error error)
	//TestInsertHistory(History rdbmsstructure.TestHistory) (Error error)
	//TestInsertOps(Ops rdbmsstructure.TestOps) (Error error)


}

func Create(env *environment.Properties) FactoryInterface {
	once.Do(func() {
		factory = gormInstance(env)
	})
	return factory
}