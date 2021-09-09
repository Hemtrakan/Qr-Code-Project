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

	if env.Flavor != environment.Production {
		db = db.Debug()
	}

	_ = db.AutoMigrate(
		&rdbmsstructure.Account{},
		&rdbmsstructure.QrCode{},
		&rdbmsstructure.HistoryInfo{},
		&rdbmsstructure.Ops{},
		//&rdbmsstructure.TestHistory{},
		//&rdbmsstructure.TestOps{},
		//&rdbmsstructure.TestQrCode{},
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

func (factory GORMFactory) GetAccountByLineId(lineId string) (response rdbmsstructure.Account, Error error) {
	var data rdbmsstructure.Account
	err := factory.client.Where("line_user_id = ?", lineId).Take(&data).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			Error = err
			return
		} else {
			Error = errors.New("ไม่มีข้อมูลผู้ใช้งานคนนี้")
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
	err := factory.client.Where("id = ? and sub_owner_id = ? ", OperatorId, OwnerId).First(&data).Error
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
	err := factory.client.Where("id = ?", id).Take(&data).Error
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

func (factory GORMFactory) GetAllAccountOwner() (response []rdbmsstructure.Account, Error error) {
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

	err := factory.client.Where("role = ?", constant.Owner).Order("created_at desc").Find(&data).Error
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
	err := factory.client.Where("sub_owner_id = ?", OwnerId).Order("created_at desc").Find(&data).Error
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

func (factory GORMFactory) GetAllAccountOperator() (response []rdbmsstructure.Account, Error error) {
	var data []rdbmsstructure.Account
	err := factory.client.Where("role = ?", constant.Operator).Order("created_at desc").Find(&data).Error
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
	err := factory.client.Preload("OpsAccount").Where("id = ?", OwnerId).Order("created_at desc").Find(&data).Error
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

func (factory GORMFactory) GetDataQrCode(QrCodeUUID string) (response []rdbmsstructure.QrCode, Error error) {
	var data []rdbmsstructure.QrCode
	db := factory.client

	err := db.Debug().Preload("DataOps", func(db *gorm.DB) *gorm.DB {
		return db.Order("ops.updated_at desc")
	}).Preload("DataHistory", func(db *gorm.DB) *gorm.DB {
		return db.Order("history_infos.updated_at desc")
	}).Where("qr_code_uuid= ?", QrCodeUUID).First(&data).Error
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
	return
}

func (factory GORMFactory) UpdateDataQrCode(Info rdbmsstructure.QrCode,HistoryInfo rdbmsstructure.HistoryInfo) (Error error){
	db := factory.client
	err := db.Where("qr_code_uuid = ?", Info.QrCodeUUID).Updates(&Info).Error
	if err != nil {
		return err
	}
	err = factory.client.Session(&gorm.Session{FullSaveAssociations: true}).Save(&HistoryInfo).Error
	if err != nil {
		return err
	}
	return
}

func (factory GORMFactory) InsertDataQrCodeById(QrCode rdbmsstructure.QrCode) (Error error) {
	db := factory.client.Where("qr_code_uuid = ?", QrCode.QrCodeUUID).Updates(&QrCode).Error
	if db != nil {
		return db
	}
	return
}

func (factory GORMFactory) UpdateHistoryInfoQrCodeById(HistoryInfo rdbmsstructure.HistoryInfo) (Error error) {
	db := factory.client.Session(&gorm.Session{FullSaveAssociations: true}).Save(&HistoryInfo).Error
	if db != nil {
		return db
	}
	return
}
func (factory GORMFactory) UpdateOpsQrCodeById(Ops rdbmsstructure.Ops) (Error error) {
	db := factory.client.Session(&gorm.Session{FullSaveAssociations: true}).Save(&Ops).Error
	if db != nil {
		return db
	}
	return
}

func (factory GORMFactory) UpdateQrCodeActive(QrCode rdbmsstructure.QrCode) (Error error) {
	var data rdbmsstructure.QrCode
	db := factory.client.Where("qr_code_uuid = ?", QrCode.QrCodeUUID).Take(&data).Error
	if db != nil {
		Error = db
		return
	}
	data.Active = QrCode.Active

	db = factory.client.Save(&data).Error
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
	err := factory.client.Where("owner_id = ? and template_name = ?", OwnerId, templateName).Order("created_at desc").Find(&data).Error
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
	err := factory.client.Where("owner_id = ?", OwnerId).Order("created_at desc").Find(&data).Error
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

func (factory GORMFactory) GetDataQrCodeInfo(QrCodeUUID string) (response rdbmsstructure.QrCode, Error error) {
	var data rdbmsstructure.QrCode
	err := factory.client.Where("qr_code_uuid = ?",QrCodeUUID).Take(&data).Error
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

func (factory GORMFactory) GetDataQrCodeOps() (response []rdbmsstructure.Ops, Error error) {
	var data []rdbmsstructure.Ops
	err := factory.client.Order("created_at DESC").Find(&data).Error
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

func (factory GORMFactory) GetDataQrCodeOpsById(ID uint) (response rdbmsstructure.Ops, Error error) {
	var data rdbmsstructure.Ops
	err := factory.client.Where("id = ?",ID).First(&data).Error
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

func (factory GORMFactory) UpdateDataQrCodeOps(ops rdbmsstructure.Ops) (Error error) {
	db := factory.client.Where("id = ?", ops.ID).Updates(&ops).Error
	if db != nil {
		return db
	}
	return
}

// -- TeamPage


func (factory GORMFactory) InsertQrCode(QrCodeUUID string, QrCode rdbmsstructure.QrCode) (Error error) {
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



// test
//func (factory GORMFactory) TestGetQrData() (response []rdbmsstructure.TestQrCode, Error error) {
//	var data []rdbmsstructure.TestQrCode
//	db := factory.client
//	err := db.Preload("DataHistory").Preload("DataOps").Where("qr_code_id = ?", 1).Find(&data).Error
//
//	if err != nil {
//		Error = errors.New("record not found")
//		return
//	}
//	response = data
//	return
//}
//
//func (factory GORMFactory) TestInsertHistory(History rdbmsstructure.TestHistory) (Error error) {
//	db := factory.client.Session(&gorm.Session{FullSaveAssociations: true}).Save(&History).Error
//	if db != nil {
//		return db
//	}
//	return nil
//}
//
//func (factory GORMFactory) TestInsertQR(Qr rdbmsstructure.TestQrCode) (Error error) {
//	db := factory.client.Session(&gorm.Session{FullSaveAssociations: true}).Save(&Qr).Error
//	if db != nil {
//		return db
//	}
//	return nil
//}
//
//func (factory GORMFactory) TestInsertOps(Ops rdbmsstructure.TestOps) (Error error) {
//	db := factory.client.Session(&gorm.Session{FullSaveAssociations: true}).Save(&Ops).Error
//	if db != nil {
//		return db
//	}
//	return nil
//}
