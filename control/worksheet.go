package control

import (
	json2 "encoding/json"
	"errors"
	"github.com/gofrs/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"qrcode/access/constant"
	rdbmsstructure "qrcode/access/rdbms/structure"
	"qrcode/present/structure"
	"time"
)

func (ctrl *APIControl) GetTypeWorksheet() structure.TypeWorksheet {
	templates := constant.Worksheet
	var arrTemplates structure.TypeWorksheet
	for _, item := range templates {
		Type := structure.TypeWorksheets{
			TypeWorksheet: item,
		}
		arrTemplates.Data = append(arrTemplates.Data, Type)
	}
	return arrTemplates
}

func (ctrl *APIControl) GetWorksheet(lineId string) (response []structure.Worksheet, Error error) {
	owner, err := ctrl.access.RDBMS.GetAccountByLineId(lineId)
	if err != nil {
		Error = err
		return
	}

	var responseArray []structure.Worksheet
	Worksheet := structure.Worksheet{}
	ops, err := ctrl.access.RDBMS.GetDataQrCodeOps()
	if err != nil {
		Error = err
		return
	}
	for _, qr := range ops {
		err = json2.Unmarshal(qr.Operator, &Worksheet)
		if err != nil {
			Error = err
			return
		}
		if *owner.SubOwnerId == Worksheet.OwnerId {
			var StatusWorksheetArray []structure.StatusWorksheet

			for _, m1 := range Worksheet.StatusWorksheet {
				StatusWorksheet := structure.StatusWorksheet{
					StatusWorksheet1: structure.StatusWorksheet1{
						Status:   m1.StatusWorksheet1.Status,
						UpdateAt: m1.StatusWorksheet1.UpdateAt,
					},
					StatusWorksheet2: structure.StatusWorksheet2{
						Status:   m1.StatusWorksheet2.Status,
						UpdateAt: m1.StatusWorksheet2.UpdateAt,
					},
					StatusWorksheet3: structure.StatusWorksheet3{
						Status:   m1.StatusWorksheet3.Status,
						UpdateAt: m1.StatusWorksheet3.UpdateAt,
						Text:     m1.StatusWorksheet3.Text,
					},
				}
				StatusWorksheetArray = append(StatusWorksheetArray, StatusWorksheet)
			}

			data := structure.Worksheet{
				ID:              qr.ID,
				QrCodeID:        Worksheet.QrCodeID,
				Text:            Worksheet.Text,
				Type:            Worksheet.Type,
				Ops:             Worksheet.Ops,
				StatusWorksheet: StatusWorksheetArray,
			}
			responseArray = append(responseArray, data)
		}
	}
	response = responseArray
	return
}

func (ctrl *APIControl) GetWorksheetById(req *structure.ReportID) (res structure.Worksheet, Error error) {
	_, err := ctrl.access.RDBMS.GetAccountByLineId(req.LineUserId)
	if err != nil {
		Error = err
		return
	}

	ops, err := ctrl.access.RDBMS.GetDataQrCodeOpsById(req.ReportID)
	if err != nil {
		Error = err
		return
	}
	Worksheet := structure.Worksheet{}
	err = json2.Unmarshal(ops.Operator, &Worksheet)
	if err != nil {
		Error = err
		return
	}
	qr, err := ctrl.access.RDBMS.GetDataQrCodeInfo(Worksheet.QrCodeID.String())
	info, err := json2.Marshal(qr.Info)

	var StatusWorksheetArray []structure.StatusWorksheet
	for _, m1 := range Worksheet.StatusWorksheet {
		StatusWorksheet := structure.StatusWorksheet{
			StatusWorksheet1: structure.StatusWorksheet1{
				Status:   m1.StatusWorksheet1.Status,
				UpdateAt: m1.StatusWorksheet1.UpdateAt,
			},
			StatusWorksheet2: structure.StatusWorksheet2{
				Status:   m1.StatusWorksheet2.Status,
				UpdateAt: m1.StatusWorksheet2.UpdateAt,
			},
			StatusWorksheet3: structure.StatusWorksheet3{
				Status:   m1.StatusWorksheet3.Status,
				UpdateAt: m1.StatusWorksheet3.UpdateAt,
				Text:     m1.StatusWorksheet3.Text,
			},
		}
		StatusWorksheetArray = append(StatusWorksheetArray, StatusWorksheet)
	}

	data := structure.Worksheet{
		ID:              ops.ID,
		QrCodeID:        Worksheet.QrCodeID,
		Info:            datatypes.JSON(info),
		Text:            Worksheet.Text,
		Type:            Worksheet.Type,
		Ops:             Worksheet.Ops,
		StatusWorksheet: StatusWorksheetArray,
	}
	res = data

	return
}

func (ctrl *APIControl) InsertWorksheet(req *structure.InsertWorksheet) (Error error) {
	QrCode, err := ctrl.access.RDBMS.GetDataQrCode(req.QrCodeID.String())
	if err != nil {
		Error = err
		return
	}
	var QrCodeID uuid.UUID
	var QrCodeRefer uint
	var OwnerId uint
	for _, qr := range QrCode {
		QrCodeID = qr.QrCodeUUID
		QrCodeRefer = qr.ID
		OwnerId = qr.OwnerId
	}

	var StatusWorksheetArray []structure.StatusWorksheet
	StatusWorksheet := structure.StatusWorksheet{
		StatusWorksheet1: structure.StatusWorksheet1{
			Status:   constant.WorksheetsStatus1,
			UpdateAt: time.Now(),
		},
		StatusWorksheet2: structure.StatusWorksheet2{},
		StatusWorksheet3: structure.StatusWorksheet3{},
	}
	StatusWorksheetArray = append(StatusWorksheetArray, StatusWorksheet)

	Worksheet := structure.Worksheet{
		QrCodeID:        req.QrCodeID,
		Text:            req.Text,
		Type:            req.Type,
		OwnerId:         OwnerId,
		StatusWorksheet: StatusWorksheetArray,
	}
	json, err := json2.Marshal(Worksheet)
	if err != nil {
		Error = err
		return
	}
	data := rdbmsstructure.Ops{
		QrCodeID:    QrCodeID,
		Operator:    datatypes.JSON(json),
		UserId:      0,
		QrCodeRefer: QrCodeRefer,
	}

	err = ctrl.access.RDBMS.UpdateOpsQrCodeById(data)

	return
}

func (ctrl *APIControl) Worksheet(reportId uint, req structure.ReportID) (Error error) {
	line, err := ctrl.access.RDBMS.GetAccountByLineId(req.LineUserId)
	if err != nil {
		Error = err
		return
	}

	ops, err := ctrl.access.RDBMS.GetDataQrCodeOpsById(reportId)
	if err != nil {
		Error = err
		return
	}
	Worksheet := structure.Worksheet{}
	err = json2.Unmarshal(ops.Operator, &Worksheet)
	if err != nil {
		Error = err
		return
	}

	setTime := time.Now()
	var StatusWorksheetArray []structure.StatusWorksheet
	for _, m1 :=range Worksheet.StatusWorksheet{
		if m1.StatusWorksheet2.Status != "" {
			Error = errors.New("มีซ่อมรับงานนี้ไปแล้ว")
			return
		}
		StatusWorksheet := structure.StatusWorksheet{
			StatusWorksheet1: structure.StatusWorksheet1{
				Status:   m1.StatusWorksheet1.Status,
				UpdateAt: m1.StatusWorksheet1.UpdateAt,
			},
			StatusWorksheet2: structure.StatusWorksheet2{
				Status:   constant.WorksheetsStatus2,
				UpdateAt: &setTime,
			},
			StatusWorksheet3: structure.StatusWorksheet3{},
		}
		StatusWorksheetArray = append(StatusWorksheetArray, StatusWorksheet)
	}




	data := structure.Worksheet{
		QrCodeID: Worksheet.QrCodeID,
		Text:     Worksheet.Text,
		Type:     Worksheet.Type,
		Ops:      &line.Username,
		OwnerId:  Worksheet.OwnerId,
		StatusWorksheet: StatusWorksheetArray,
	}
	jsonData, err := json2.Marshal(data)
	if err != nil {
		Error = err
		return
	}

	dataOps := rdbmsstructure.Ops{
		Model: gorm.Model{
			ID: reportId,
		},
		QrCodeID: Worksheet.QrCodeID,
		Operator: datatypes.JSON(jsonData),
		UserId:   line.ID,
	}

	err = ctrl.access.RDBMS.UpdateDataQrCodeOps(dataOps)
	if err != nil {
		Error = err
		return
	}
	return
}

func (ctrl *APIControl) GetUpdateWorksheet(QrCodeId string) (res structure.GetWorksheet, Error error) {
	qrCode, err := ctrl.access.RDBMS.GetDataQrCodeInfo(QrCodeId)
	if err != nil {
		Error = err
		return
	}

	var responseArray []structure.Worksheet
	Worksheet := structure.Worksheet{}
	ops, err := ctrl.access.RDBMS.GetDataQrCodeOps()
	if err != nil {
		Error = err
		return
	}
	for _, qr := range ops {

		err = json2.Unmarshal(qr.Operator, &Worksheet)
		if err != nil {
			Error = err
			return
		}
		if qr.QrCodeID.String() == QrCodeId {
			var StatusWorksheetArray []structure.StatusWorksheet
			for _, m1 := range Worksheet.StatusWorksheet {
				StatusWorksheet := structure.StatusWorksheet{
					StatusWorksheet1: structure.StatusWorksheet1{
						Status:   m1.StatusWorksheet1.Status,
						UpdateAt: m1.StatusWorksheet1.UpdateAt,
					},
					StatusWorksheet2: structure.StatusWorksheet2{
						Status:   m1.StatusWorksheet2.Status,
						UpdateAt: m1.StatusWorksheet2.UpdateAt,
					},
					StatusWorksheet3: structure.StatusWorksheet3{
						Status:   m1.StatusWorksheet3.Status,
						UpdateAt: m1.StatusWorksheet3.UpdateAt,
						Text:     m1.StatusWorksheet3.Text,
					},
				}
				StatusWorksheetArray = append(StatusWorksheetArray, StatusWorksheet)
			}

			data := structure.Worksheet{
				ID:              qr.ID,
				QrCodeID:        Worksheet.QrCodeID,
				Text:            Worksheet.Text,
				Type:            Worksheet.Type,
				Ops:             Worksheet.Ops,
				StatusWorksheet: StatusWorksheetArray,
			}
			responseArray = append(responseArray, data)
		}
	}

	data := structure.GetWorksheet{
		Info:      qrCode.Info,
		Worksheet: responseArray,
	}

	res = data
	return
}

func (ctrl *APIControl) UpdateWorksheet(reportId uint, req structure.UpdateWorksheet) (Error error) {
	line, err := ctrl.access.RDBMS.GetAccountByLineId(req.LineUserId)
	if err != nil {
		Error = err
		return
	}

	ops, err := ctrl.access.RDBMS.GetDataQrCodeOpsById(reportId)
	if err != nil {
		Error = err
		return
	}
	Worksheet := structure.Worksheet{}
	err = json2.Unmarshal(ops.Operator, &Worksheet)
	if err != nil {
		Error = err
		return
	}

	setTime := time.Now()
	var StatusWorksheetArray []structure.StatusWorksheet
	for _ , m1 := range Worksheet.StatusWorksheet{
		if m1.StatusWorksheet3.Status != "" {
			Error = errors.New("มีการแก้ไขปัญหาเรียบร้อยแล้ว")
			return
		}
		StatusWorksheet := structure.StatusWorksheet{
			StatusWorksheet1: structure.StatusWorksheet1{
				Status:   m1.StatusWorksheet1.Status,
				UpdateAt: m1.StatusWorksheet1.UpdateAt,
			},
			StatusWorksheet2: structure.StatusWorksheet2{
				Status:   m1.StatusWorksheet2.Status,
				UpdateAt: m1.StatusWorksheet2.UpdateAt,
			},
			StatusWorksheet3: structure.StatusWorksheet3{
				Status:   constant.WorksheetsStatus3,
				UpdateAt: &setTime,
				Text:     req.Text,
			},
		}
		StatusWorksheetArray = append(StatusWorksheetArray, StatusWorksheet)
	}


	data := structure.Worksheet{
		QrCodeID: Worksheet.QrCodeID,
		Text:     Worksheet.Text,
		Type:     Worksheet.Type,
		Ops:      &line.Username,
		OwnerId:  Worksheet.OwnerId,
		StatusWorksheet: StatusWorksheetArray,
	}
	jsonData, err := json2.Marshal(data)
	if err != nil {
		Error = err
		return
	}
	dataOps := rdbmsstructure.Ops{
		Model: gorm.Model{
			ID: reportId,
		},
		QrCodeID: Worksheet.QrCodeID,
		Operator: datatypes.JSON(jsonData),
		UserId:   line.ID,
	}

	err = ctrl.access.RDBMS.UpdateDataQrCodeOps(dataOps)
	if err != nil {
		Error = err
		return
	}
	return
}

func (ctrl *APIControl) DeleteWorksheet(reportId uint, req structure.UpdateWorksheet) (Error error) {
	line, err := ctrl.access.RDBMS.GetAccountByLineId(req.LineUserId)
	if err != nil {
		Error = err
		return
	}

	ops, err := ctrl.access.RDBMS.GetDataQrCodeOpsById(reportId)
	if err != nil {
		Error = err
		return
	}
	Worksheet := structure.Worksheet{}
	err = json2.Unmarshal(ops.Operator, &Worksheet)
	if err != nil {
		Error = err
		return
	}


	setTime := time.Now()
	var StatusWorksheetArray []structure.StatusWorksheet
	for _ , m1 := range Worksheet.StatusWorksheet{
		if m1.StatusWorksheet3.Status != "" {
			Error = errors.New("มีการแก้ไขปัญหาเรียบร้อยแล้ว")
			return
		}
		StatusWorksheet := structure.StatusWorksheet{
			StatusWorksheet1: structure.StatusWorksheet1{
				Status:   m1.StatusWorksheet1.Status,
				UpdateAt: m1.StatusWorksheet1.UpdateAt,
			},
			StatusWorksheet2: structure.StatusWorksheet2{
				Status:   m1.StatusWorksheet2.Status,
				UpdateAt: m1.StatusWorksheet2.UpdateAt,
			},
			StatusWorksheet3: structure.StatusWorksheet3{
				Status:   constant.WorksheetsStatus4,
				UpdateAt: &setTime,
				Text:     req.Text,
			},
		}
		StatusWorksheetArray = append(StatusWorksheetArray, StatusWorksheet)
	}


	data := structure.Worksheet{
		QrCodeID: Worksheet.QrCodeID,
		Text:     Worksheet.Text,
		Type:     Worksheet.Type,
		Ops:      &line.Username,
		OwnerId:  Worksheet.OwnerId,
		StatusWorksheet: StatusWorksheetArray,
	}
	jsonData, err := json2.Marshal(data)
	if err != nil {
		Error = err
		return
	}
	dataOps := rdbmsstructure.Ops{
		Model: gorm.Model{
			ID: reportId,
		},
		QrCodeID: Worksheet.QrCodeID,
		Operator: datatypes.JSON(jsonData),
		UserId:   line.ID,
	}

	err = ctrl.access.RDBMS.UpdateDataQrCodeOps(dataOps)
	if err != nil {
		Error = err
		return
	}
	return
}
