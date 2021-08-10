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
		&rdbmsstructure.Account{},
		&rdbmsstructure.QrCode{},
	)
	return GORMFactory{env: env, client: db}
}


//
//func (factory GORMFactory) GetIdTeamPage(teamPageId uint) (response rdbmsstructure.Template, Error error) {
//	var data rdbmsstructure.Template
//	err := factory.client.Where("id = ?", teamPageId).Find(&data).Error
//	if err != nil {
//		if !errors.Is(err, gorm.ErrRecordNotFound) {
//			Error = err
//		} else {
//			return
//		}
//		return
//	}
//	response = data
//	return
//}

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

func (factory GORMFactory) UpdateProfile(Account rdbmsstructure.Account) (Error error) {
	db := factory.client.Where("id = ?", Account.ID).Updates(&Account).Error
	if db != nil {
		return db
	}
	return nil
}

func (factory GORMFactory) DeleteAccount(id int) (Error error) {
	var data rdbmsstructure.Account
	err := factory.client.Where("id = ?", id).Delete(&data).Error
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

// -- QR-Code

func (factory GORMFactory) CreateQrCode(QrCode rdbmsstructure.QrCode) (Error error) {
	db := factory.client.Session(&gorm.Session{FullSaveAssociations: true}).Save(&QrCode).Error
	if db != nil {
		return db
	}
	return nil
}

func (factory GORMFactory) UpdateQrCodeById(QrCode rdbmsstructure.QrCode) (Error error) {
	db := factory.client.Where("qr_code_uuid = ?", QrCode.QrCodeUUID).Updates(&QrCode).Error
	if db != nil {
		return db
	}
	return nil
}

func (factory GORMFactory) UpdateQrCode(QrCode rdbmsstructure.QrCode) (Error error) {
	db := factory.client.Where("id = ?", QrCode.ID).Updates(&QrCode).Error
	if db != nil {
		return db
	}
	return nil
}

func (factory GORMFactory) GetQrCode(OwnerId uint,templateName string) (response []rdbmsstructure.QrCode,Error error) {
	var data []rdbmsstructure.QrCode
	err := factory.client.Where("owner_id = ? and template_name = ?", OwnerId,templateName).Find(&data).Error
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

func (factory GORMFactory) GetQrCodeById(OwnerId int) (response []rdbmsstructure.QrCode,Error error) {
	var data []rdbmsstructure.QrCode
	err := factory.client.Where("owner_id = ?", OwnerId).Find(&data).Error
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



// -- TeamPage
//
//func (factory GORMFactory) InsertTeamPage(TeamPage rdbmsstructure.Template) (response rdbmsstructure.Template,Error error){
//	err := factory.client.Session(&gorm.Session{FullSaveAssociations: true}).Save(&TeamPage).Error
//	if err != nil {
//		Error = err
//		return
//	}
//	response = TeamPage
//	return
//}
//func (factory GORMFactory) UpdateTeamPage(TeamPage rdbmsstructure.Template) (Error error){
//	db := factory.client.Where("id = ?", TeamPage.ID).Updates(&TeamPage).Error
//	if db != nil {
//		return db
//	}
//	return nil
//}
//
//func (factory GORMFactory) GetByIdTeamPage(teamPageId string) (response rdbmsstructure.Template, Error error) {
//	var data rdbmsstructure.Template
//	err := factory.client.Where("uuid = ?", teamPageId).Find(&data).Error
//	if err != nil {
//		if !errors.Is(err, gorm.ErrRecordNotFound) {
//			Error = err
//		} else {
//			return
//		}
//		return
//	}
//	response = data
//	return
//}
//
//func (factory GORMFactory) DeleteTeamPage(teamPageId uint) (Error error) {
//	var data rdbmsstructure.Template
//	err := factory.client.Where("id = ?", teamPageId).Delete(&data).Error
//	if err != nil {
//		if !errors.Is(err, gorm.ErrRecordNotFound) {
//			Error = err
//		} else {
//			return
//		}
//		return
//	}
//	return
//}
//
//func (factory GORMFactory) GetAllTeamPage(ownersId int) (response []rdbmsstructure.Template, Error error) {
//	var data []rdbmsstructure.Template
//	err := factory.client.Where("owners_id = ?", ownersId).Find(&data).Error
//	if err != nil {
//		if !errors.Is(err, gorm.ErrRecordNotFound) {
//			Error = err
//		} else {
//			return
//		}
//		return
//	}
//	response = data
//	return
//}
//


// InsertLogTeamPage -- LogTeamPage
//func (factory GORMFactory) InsertLogTeamPage(LogTeamPage rdbmsstructure.LogTeamPage) (Error error) {
//	db := factory.client.Session(&gorm.Session{FullSaveAssociations: true}).Save(&LogTeamPage).Error
//	if db != nil {
//		return db
//	}
//	return
//}
//
//func (factory GORMFactory) GetByIdLogTeamPage(teamPageId uint) (response []rdbmsstructure.LogTeamPage, Error error) {
//	var data []rdbmsstructure.LogTeamPage
//	err := factory.client.Where("team_page_id = ?", teamPageId).Find(&data).Error // todo ยังไม่เสร็จ
//	if err != nil {
//		if !errors.Is(err, gorm.ErrRecordNotFound) {
//			Error = err
//		} else {
//			return
//		}
//		return
//	}
//	response = data
//	return
//}
//
//
//func (factory GORMFactory) GetAllDataListLogTeamPage(teamPageId uint) (response []rdbmsstructure.LogTeamPage, Error error) {
//	var data []rdbmsstructure.LogTeamPage
//	err := factory.client.Where("team_page_id = ?", teamPageId).Order("id desc").Find(&data).Error // todo ยังไม่เสร็จ
//	if err != nil {
//		if !errors.Is(err, gorm.ErrRecordNotFound) {
//			Error = err
//		} else {
//			return
//		}
//		return
//	}
//	response = data
//	return
//}




