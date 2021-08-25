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
		&rdbmsstructure.History{},
	)
	return GORMFactory{env: env, client: db}
}

//  Account

func (factory GORMFactory) CheckUserRegister(Username, PhoneNumber, LineId string, UserId uint) (response *rdbmsstructure.Account, Error error) {
	var data *rdbmsstructure.Account
	err := factory.client.Where("username = ? ", Username).First(&data).Error
	if err == nil {
		Error = errors.New("ชื่อผู้ใช้นี้อยู่แล้ว")
		return
	}
	err = factory.client.Where("phone_number = ?", PhoneNumber).First(&data).Error
	if err == nil {
		Error = errors.New("เบอร์โทรศัพท์นี้มีอยู่แล้ว")
		return
	}
	err = factory.client.Where("line_id = ?", LineId).First(&data).Error
	if err == nil {
		Error = errors.New("LineId นี้มีอยู่แล้ว")
		return
	}
	response = data
	return
}

func (factory GORMFactory) Register(Account rdbmsstructure.Account) (Error error) {
	err := factory.client.Session(&gorm.Session{FullSaveAssociations: true}).Save(&Account).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRegistered) {
			Error = err
			return
		}
	}
	return
}

func (factory GORMFactory) Login(login rdbmsstructure.Account) (response rdbmsstructure.Account, Error error) {
	var data rdbmsstructure.Account
	err := factory.client.Where("username = ?", login.Username).First(&data).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			Error = err
		} else {
			Error = errors.New("record not found")
			return
		}
		return
	}
	response = data
	return
}

func (factory GORMFactory) CheckAccountId(id uint) (response *rdbmsstructure.Account, Error error) {
	var data *rdbmsstructure.Account
	err := factory.client.Where("id = ?", id).First(&data).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			Error = err
			return
		} else {
			Error = errors.New("record not found")
			return
		}
		return
	}
	response = data
	return
}

func (factory GORMFactory) GetOperatorById(OperatorId int, OwnerId uint) (response rdbmsstructure.Account, Error error) {
	var data rdbmsstructure.Account
	err := factory.client.Where("id = ? and sub_owner_id = ? ", OperatorId,OwnerId).First(&data).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			Error = err
		} else {
			Error = errors.New("record not found")
			return
		}
		return
	}
	response = data
	return
}

func (factory GORMFactory) GetAccount(id int) (response rdbmsstructure.Account, Error error) {
	var data rdbmsstructure.Account
	err := factory.client.Where("id = ?", id).First(&data).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			Error = err
		} else {
			Error = errors.New("record not found")
			return
		}
		return
	}
	response = data
	return
}

func (factory GORMFactory) GetAllAccountOwner() (response []rdbmsstructure.Account,  Error error) {
	var data []rdbmsstructure.Account
	//db := factory.client.Where("role = ?", constant.Owner)
	//if Firstname != nil {
	//	db = db.Where("first_name like ? ", fmt.Sprintf("%s", "%"+*Firstname+"%"))
	//}
	//if Lastname != nil {
	//	db = db.Where("last_name like ?", fmt.Sprintf("%s", "%"+*Lastname+"%"))
	//}
	//if Phonenumber != nil {
	//	db = db.Where("phone_number like ?", fmt.Sprintf("%s", "%"+*Phonenumber+"%"))
	//}
	//if Lineid != nil {
	//	db = db.Where("line_id like ?", fmt.Sprintf("%s", "%"+*Lineid+"%"))
	//}
	//
	//pagination := utility.Paging(&utility.Param{
	//	DB:      db,
	//	Page:    *page,
	//	Limit:   *limit,
	//	OrderBy: []string{"created_at asc"},
	//}, &data)
	//paginator = *pagination
	//
	//if db.Error != nil {
	//	Error = db.Error
	//	return
	//}
	//response = data

	err := factory.client.Where("role = ?", constant.Owner).Order("created_at asc").Find(&data).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			Error = err
		} else {
			Error = errors.New("record not found")
			return
		}
		return
	}
	response = data
	return
}

func (factory GORMFactory) GetAllAccountOperatorByOwnerID(OwnerId uint) (response []rdbmsstructure.Account, Error error) {
	var data []rdbmsstructure.Account
	err := factory.client.Where("sub_owner_id = ?", OwnerId).Order("created_at asc").Find(&data).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			Error = err
		} else {
			Error = errors.New("record not found")
			return
		}
		return
	}
	response = data
	return
}

func (factory GORMFactory) GetAllAccountOperator() (response []rdbmsstructure.Account,  Error error) {
	var data []rdbmsstructure.Account
	err := factory.client.Where("role = ?", constant.Operator).Order("created_at asc").Find(&data).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			Error = err
		} else {
			Error = errors.New("record not found")
			return
		}
		return
	}
	response = data
	//db := factory.client.Where("role = ?", constant.Operator)
	//if Firstname != nil {
	//	db = db.Where("first_name like ? ", fmt.Sprintf("%s", "%"+*Firstname+"%"))
	//}
	//if Lastname != nil {
	//	db = db.Where("last_name like ?", fmt.Sprintf("%s", "%"+*Lastname+"%"))
	//}
	//if Phonenumber != nil {
	//	db = db.Where("phone_number like ?", fmt.Sprintf("%s", "%"+*Phonenumber+"%"))
	//}
	//if Lineid != nil {
	//	db = db.Where("line_id like ?", fmt.Sprintf("%s", "%"+*Lineid+"%"))
	//}
	//
	//pagination := utility.Paging(&utility.Param{
	//	DB:      db,
	//	Page:    *page,
	//	Limit:   *limit,
	//	OrderBy: []string{"created_at asc"},
	//}, &data)
	//paginator = *pagination
	//
	//if db.Error != nil {
	//	Error = db.Error
	//	return
	//}
	//response = data
	return
}

func (factory GORMFactory) GetSubOwner(OwnerId int) (response []rdbmsstructure.Account, Error error) {
	var data []rdbmsstructure.Account
	err := factory.client.Where("sub_owner_id = ?", OwnerId).Order("created_at asc").Find(&data).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			Error = err
		} else {
			Error = errors.New("record not found")
			return
		}
		return
	}
	response = data
	return
}

func (factory GORMFactory) GetOwnerByIdOps(OperatorId int) (response rdbmsstructure.Account, Error error) {
	var data rdbmsstructure.Account
	//err := factory.client.Raw("SELECT * FROM accounts as ownerAccount INNER JOIN accounts as OpsAccount ON ownerAccount.id = OpsAccount.sub_owner_id").Where("sub_owner_id = ?", OperatorId).Find(&data).Error
	err := factory.client.Where("id = ?", OperatorId).Find(&data).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			Error = err
		} else {
			Error = errors.New("record not found")
			return
		}
		return
	}
	response = data
	return
}

func (factory GORMFactory) UpdateProfile(Account rdbmsstructure.Account) (Error error) {
	err := factory.client.Where("id = ?", Account.ID).Updates(&Account).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			Error = err
			return
		} else {
			Error = errors.New("record not found")
			return
		}
		return
	}
	return
}

func (factory GORMFactory) DeleteAccount(id uint) (Error error) {
	var data rdbmsstructure.Account
	err := factory.client.Where("id = ?", id).Delete(&data).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			Error = err
		} else {
			Error = errors.New("record not found")
			return
		}
		return
	}
	return
}

func (factory GORMFactory) DeleteAccountByOwner(OwnerId uint, OperatorId int) (Error error) {
	var data rdbmsstructure.Account
	err := factory.client.Where("id = ?", OperatorId).Where("sub_owner_id = ?", OwnerId).First(&data).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			Error = err
		} else {
			Error = errors.New("record not found")
			return
		}
		return
	}

	err = factory.client.Where("id = ?", OperatorId).Where("sub_owner_id = ?", OwnerId).Delete(&data).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			Error = err
		} else {
			Error = errors.New("record not found")
		}
		return
	}
	return
}

// -- QR-Code

func (factory GORMFactory) GetDataQrCode(QrCodeUUID string) (response rdbmsstructure.QrCode, Error error) {
	var data rdbmsstructure.QrCode
	err := factory.client.Where("qr_code_uuid= ?", QrCodeUUID).First(&data).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			Error = err
		} else {
			Error = errors.New("record not found")
			return
		}
		return
	}
	response = data
	return
}

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

func (factory GORMFactory) CountCode(OwnerId uint, templateName, Code string) (response []rdbmsstructure.QrCode, Error error) {
	var data []rdbmsstructure.QrCode
	err := factory.client.Where("owner_id = ? and template_name = ? and code = ?", OwnerId, templateName, Code).Find(&data).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			Error = err
		} else {
			Error = errors.New("record not found")
			return
		}
		return
	}
	response = data
	return
}

func (factory GORMFactory) CheckCode(OwnerId uint, templateName, Code string) (response rdbmsstructure.QrCode, Error error) {
	var data rdbmsstructure.QrCode
	err := factory.client.Where("owner_id = ? and template_name = ? and code = ?", OwnerId, templateName, Code).Find(&data).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			Error = err
		} else {
			Error = errors.New("record not found")
			return
		}
		return
	}
	response = data
	return
}

func (factory GORMFactory) GetQrCode(OwnerId uint, templateName string) (response []rdbmsstructure.QrCode, Error error) {
	var data []rdbmsstructure.QrCode
	err := factory.client.Where("owner_id = ? and template_name = ?", OwnerId, templateName).Order("created_at asc").Find(&data).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			Error = err
		} else {
			Error = errors.New("record not found")
			return
		}
		return
	}
	response = data
	return
}


func (factory GORMFactory) GetAllQrCode() (response []rdbmsstructure.QrCode, Error error) {
	var data []rdbmsstructure.QrCode
	err := factory.client.Order("created_at asc").Find(&data).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			Error = err
			return
		} else {
			Error = errors.New("record not found")
			return
		}
		return
	}
	response = data
	return
}

func (factory GORMFactory) GetQrCodeByOwnerId(OwnerId int) (response []rdbmsstructure.QrCode, Error error) {
	var data []rdbmsstructure.QrCode
	err := factory.client.Where("owner_id = ?", OwnerId).Order("created_at asc").Find(&data).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			Error = err
			return
		} else {
			Error = errors.New("record not found")
			return
		}
		return
	}
	response = data
	return
}

func (factory GORMFactory) GetQrCodeByQrCodeId(OwnerId int, QrCodeId string) (response rdbmsstructure.QrCode, Error error) {
	var data rdbmsstructure.QrCode
	err := factory.client.Where("owner_id = ? AND qr_code_uuid = ?", OwnerId, QrCodeId).First(&data).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			Error = err
		} else {
			Error = errors.New("record not found")
			return
		}
		return
	}
	response = data
	return
}

func (factory GORMFactory) DeleteQrCode(QrCodeUUID string) (Error error) {
	var data rdbmsstructure.QrCode
	err := factory.client.Where("qr_code_uuid = ?", QrCodeUUID).Delete(&data).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			Error = err
			return
		} else {
			Error = errors.New("record not found")
			return
		}
		return
	}
	return
}

func (factory GORMFactory) UpdateStatusQrCode(QrCode rdbmsstructure.QrCode) (Error error) {
	err := factory.client.Where("qr_code_uuid = ?", QrCode.QrCodeUUID).Updates(&QrCode).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			Error = err
			return
		} else {
			Error = errors.New("record not found")
			return
		}
		return
	}
	return
}

// -- TeamPage

func (factory GORMFactory) GetHistory(QrCodeUUID string) (response []rdbmsstructure.History, Error error) {
	var data []rdbmsstructure.History
	err := factory.client.Where("qr_code_uuid= ?", QrCodeUUID).First(&data).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			Error = err
		} else {
			Error = errors.New("record not found")
			return
		}
		return
	}
	response = data
	return
}

func (factory GORMFactory) InsertQrCode(QrCodeUUID string,QrCode rdbmsstructure.QrCode) (Error error) {
	err := factory.client.Where("qr_code_uuid = ?", QrCodeUUID).Updates(QrCode).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			Error = err
			return
		} else {
			Error = errors.New("record not found")
			return
		}
		return
	}
	return
}

func (factory GORMFactory) InsertHistory(History rdbmsstructure.History) (Error error) {
	db := factory.client.Session(&gorm.Session{FullSaveAssociations: true}).Save(&History).Error
	if db != nil {
		return db
	}
	return nil
}
