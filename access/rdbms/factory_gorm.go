package rdbms

import (
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"qrcode/access/constant"
	rdbmsstructure "qrcode/access/rdbms/structure"
	"qrcode/environment"
)

type GORMFactory struct {
	env    *environment.Properties
	client *gorm.DB
}

func gormInstance(env *environment.Properties) GORMFactory {
	databaseSet := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		env.GormHost, env.GormPort, env.GormUser, env.GormName, env.GormPass, "disable")

	db, err := gorm.Open(postgres.Open(databaseSet), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("failed to connect database : %s", err.Error()))
		//panic(fmt.Sprintf("failed to connect database : %s", err.Error()))
	}

	_ = db.AutoMigrate(
		&rdbmsstructure.Qrcode{},
		&rdbmsstructure.Account{},
		&rdbmsstructure.Location{},
		&rdbmsstructure.TeamPage{},
		&rdbmsstructure.LogTeamPage{},
	)
	return GORMFactory{env: env, client: db}
}

func (factory GORMFactory) GenerateQrCode(Qrcode rdbmsstructure.Qrcode) (Error error) {
	db := factory.client.Session(&gorm.Session{FullSaveAssociations: true}).Save(&Qrcode).Error
	if db != nil {
		return db
	}
	return nil
	return
}

func (factory GORMFactory) GetIdTeamPage(teamPageId uint) (response rdbmsstructure.TeamPage, Error error) {
	var data rdbmsstructure.TeamPage
	err := factory.client.Where("id = ?", teamPageId).Find(&data).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			Error = err
		} else {
			return
		}
		return
	}
	response = data
	return
}

//  Account
func (factory GORMFactory) Register(Account rdbmsstructure.Account) (Error error) {
	db := factory.client.Session(&gorm.Session{FullSaveAssociations: true}).Save(&Account).Error
	if db != nil {
		return db
	}
	return nil
}

func (factory GORMFactory) Login(login rdbmsstructure.Account) (response rdbmsstructure.Account, Error error) {
	var data rdbmsstructure.Account
	err := factory.client.Where("username = ?", login.Username).First(&data).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			Error = err
		} else {
			return
		}
		return
	}
	response = data
	return
}

func (factory GORMFactory) GetAccount(id int) (response rdbmsstructure.Account,Error error){
	var data rdbmsstructure.Account
	err := factory.client.Where("id = ?", id).First(&data).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			Error = err
		} else {
			return
		}
		return
	}
	response = data
	return
}

func (factory GORMFactory) GetAllAccountOwner() (response []rdbmsstructure.Account, Error error) {
	var data []rdbmsstructure.Account
	err := factory.client.Where("role = ?", constant.Owner).Find(&data).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			Error = err
		} else {
			return
		}
		return
	}
	response = data
	return
}

func (factory GORMFactory) GetAllAccountOperator() (response []rdbmsstructure.Account, Error error) {
	var data []rdbmsstructure.Account
	err := factory.client.Where("role = ?", constant.Operator).Find(&data).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			Error = err
		} else {
			return
		}
		return
	}
	response = data
	return
}

func (factory GORMFactory) GetByIdAccount(id int) (response rdbmsstructure.Account, Error error) {
	panic("implement me")
}

func (factory GORMFactory) UpdateProfile(Account rdbmsstructure.Account) (Error error) {
	panic("implement me")
}

func (factory GORMFactory) DeleteAccount(id int) (Error error) {
	panic("implement me")
}

// -- TeamPage

func (factory GORMFactory) InsertTeamPage(TeamPage rdbmsstructure.TeamPage) (Error error){
	db := factory.client.Session(&gorm.Session{FullSaveAssociations: true}).Save(&TeamPage).Error
	if db != nil {
		return db
	}
	return nil
}
func (factory GORMFactory) UpdateTeamPage(TeamPage rdbmsstructure.TeamPage) (Error error){
	db := factory.client.Where("id = ?", TeamPage.ID).Updates(&TeamPage).Error
	if db != nil {
		return db
	}
	return nil
}

func (factory GORMFactory) GetByIdTeamPage(teamPageId string) (response rdbmsstructure.TeamPage, Error error) {
	var data rdbmsstructure.TeamPage
	err := factory.client.Where("uuid = ?", teamPageId).Find(&data).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			Error = err
		} else {
			return
		}
		return
	}
	response = data
	return
}

func (factory GORMFactory) DeleteTeamPage(teamPageId uint) (Error error) {
	var data rdbmsstructure.TeamPage
	err := factory.client.Where("id = ?", teamPageId).Delete(&data).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			Error = err
		} else {
			return
		}
		return
	}
	return
}

func (factory GORMFactory) GetAllTeamPage(ownersId int) (response []rdbmsstructure.TeamPage, Error error) {
	var data []rdbmsstructure.TeamPage
	err := factory.client.Where("owners_id = ?", ownersId).Find(&data).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			Error = err
		} else {
			return
		}
		return
	}
	response = data
	return
}



// InsertLogTeamPage -- LogTeamPage
func (factory GORMFactory) InsertLogTeamPage(LogTeamPage rdbmsstructure.LogTeamPage) (Error error) {
	db := factory.client.Session(&gorm.Session{FullSaveAssociations: true}).Save(&LogTeamPage).Error
	if db != nil {
		return db
	}
	return
}

func (factory GORMFactory) GetByIdLogTeamPage(teamPageId uint) (response []rdbmsstructure.LogTeamPage, Error error) {
	var data []rdbmsstructure.LogTeamPage
	err := factory.client.Where("team_page_id = ?", teamPageId).Find(&data).Error // todo ยังไม่เสร็จ
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			Error = err
		} else {
			return
		}
		return
	}
	response = data
	return
}


func (factory GORMFactory) GetAllDataListLogTeamPage(teamPageId uint) (response []rdbmsstructure.LogTeamPage, Error error) {
	var data []rdbmsstructure.LogTeamPage
	err := factory.client.Where("team_page_id = ?", teamPageId).Order("id desc").Find(&data).Error // todo ยังไม่เสร็จ
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			Error = err
		} else {
			return
		}
		return
	}
	response = data
	return
}




