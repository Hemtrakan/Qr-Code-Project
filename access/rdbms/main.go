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
	GetByIdAccount(id int) (response rdbmsstructure.Account,Error error)
	UpdateProfile(Account rdbmsstructure.Account) (Error error)
	DeleteAccount(id int) (Error error)

	//  QRCode
	GenerateQrCode(Qrcode rdbmsstructure.Qrcode) (Error error)
	GetIdTeamPage(teamPageId uint) (response rdbmsstructure.TeamPage, Error error)


	// Acconut
	Register(Account rdbmsstructure.Account) (Error error)
	Login(login rdbmsstructure.Account) (response rdbmsstructure.Account,Error error)
	GetAccount(id int) (response rdbmsstructure.Account,Error error)

	//  TeamPage
	InsertTeamPage(TeamPage rdbmsstructure.TeamPage) (Error error)
	UpdateTeamPage(TeamPage rdbmsstructure.TeamPage) (Error error)
	GetAllTeamPage(ownersId int) (response []rdbmsstructure.TeamPage ,Error error)
	GetByIdTeamPage(teamPageId string) (response rdbmsstructure.TeamPage ,Error error)
	DeleteTeamPage(teamPageId uint) (Error error)

	// InsertLogTeamPage -- LogTeamPage
	InsertLogTeamPage(TeamPage rdbmsstructure.LogTeamPage) (Error error)
	GetByIdLogTeamPage(teamPageId uint) (response []rdbmsstructure.LogTeamPage ,Error error)

	GetAllDataListLogTeamPage(teamPageId uint) (response []rdbmsstructure.LogTeamPage ,Error error)
}

func Create(env *environment.Properties) FactoryInterface {
	once.Do(func() {
		factory = gormInstance(env)
	})
	return factory
}