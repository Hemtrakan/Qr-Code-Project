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

// Owner
func (ctrl *APIControl) UpdateOption(OwnerId uint, req structure.UpdateOption) (Error error) {
	data, err := ctrl.access.RDBMS.GetQrCode(OwnerId, string(constant.OfficeEquipment))
	if err != nil {
		Error = err
		return
	}
	for _, m1 := range data {
		dataOps, err := ctrl.access.RDBMS.GetDataQrCodeOpsByQrCodeID(m1.QrCodeUUID.String())
		if err != nil {
			Error = err
			return
		}
		for _, m2 := range dataOps {
			Worksheet := structure.Worksheet{}
			err = json2.Unmarshal(m2.Operator, &Worksheet)
			dataWorksheet := structure.Worksheet{
				QrCodeID:        Worksheet.QrCodeID,
				Text:            Worksheet.Text,
				Option:          req.Option,
				Type:            Worksheet.Type,
				Ops:             Worksheet.Ops,
				OwnerId:         Worksheet.OwnerId,
				StatusWorksheet: Worksheet.StatusWorksheet,
			}
			jsonData, err := json2.Marshal(dataWorksheet)
			if err != nil {
				Error = err
				return
			}
			dataSave := rdbmsstructure.Ops{
				Model: gorm.Model{
					ID: m2.ID,
				},
				Operator: datatypes.JSON(jsonData),
			}

			err = ctrl.access.RDBMS.UpdateDataQrCodeOps(dataSave)
			if err != nil {
				Error = err
				return
			}
		}
	}
	return
}

func (ctrl *APIControl) OwnerGetWorksheet(OwnerID uint) (response []structure.Worksheet, Error error) {
	var responseArray []structure.Worksheet
	ops, err := ctrl.access.RDBMS.GetDataQrCodeOps()
	if err != nil {
		Error = err
		return
	}
	for _, qr := range ops {
		Worksheet := structure.Worksheet{}
		err = json2.Unmarshal(qr.Operator, &Worksheet)
		if err != nil {
			Error = err
			return
		}

		if OwnerID == Worksheet.OwnerId {
			var StatusWorksheetArray []structure.StatusWorksheet
			for _, m1 := range Worksheet.StatusWorksheet {
				var Equipments []structure.Equipment
				var Text *string
				if m1.Status != "" {
					if m1.Status != constant.WorksheetsStatus4 {
						Equipment := structure.Equipment{}
						for _, m2 := range m1.Equipments {
							Equipment = structure.Equipment{
								NameEquipment: m2.NameEquipment,
							}
							Equipments = append(Equipments, Equipment)
						}
						Text = m1.Text
					}
				}

				StatusWorksheet := structure.StatusWorksheet{
					Status:     m1.Status,
					UpdateAt:   m1.UpdateAt,
					Text:       Text,
					Equipments: Equipments,
				}
				StatusWorksheetArray = append(StatusWorksheetArray, StatusWorksheet)
			}
			var Ops *string
			Ops = Worksheet.Ops
			data := structure.Worksheet{
				ID:              qr.ID,
				QrCodeID:        Worksheet.QrCodeID,
				Option:          Worksheet.Option,
				Text:            Worksheet.Text,
				Type:            Worksheet.Type,
				Ops:             Ops,
				StatusWorksheet: StatusWorksheetArray,
			}
			responseArray = append(responseArray, data)
		}
	}
	response = responseArray
	return
}

func (ctrl *APIControl) OwnerGetWorksheetById(ReportID uint) (res structure.Worksheet, Error error) {
	ops, err := ctrl.access.RDBMS.GetDataQrCodeOpsById(ReportID)
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
	for index, m1 := range Worksheet.StatusWorksheet {
		if index == 0 {
			StatusWorksheet1 := structure.StatusWorksheet{
				Status:     m1.Status,
				UpdateAt:   m1.UpdateAt,
				Text:       m1.Text,
				Equipments: m1.Equipments,
			}
			StatusWorksheetArray = append(StatusWorksheetArray, StatusWorksheet1)
		}
		if index == 1 {
			StatusWorksheet2 := structure.StatusWorksheet{
				Status:     m1.Status,
				UpdateAt:   m1.UpdateAt,
				Text:       m1.Text,
				Equipments: m1.Equipments,
			}
			StatusWorksheetArray = append(StatusWorksheetArray, StatusWorksheet2)
		}
		if index == 2 {
			StatusWorksheet3 := structure.StatusWorksheet{
				Status:     m1.Status,
				UpdateAt:   m1.UpdateAt,
				Text:       m1.Text,
				Equipments: m1.Equipments,
			}
			StatusWorksheetArray = append(StatusWorksheetArray, StatusWorksheet3)
		}
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

func (ctrl *APIControl) OwnerWorksheet(reportId uint, req structure.ReportID) (Error error) {
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
	for _, m1 := range Worksheet.StatusWorksheet {
		if m1.Status == constant.WorksheetsStatus2 {
			Error = errors.New("มีช่างรับงานนี้ไปแล้ว")
			return
		}
		StatusWorksheet1 := structure.StatusWorksheet{
			Status:     m1.Status,
			UpdateAt:   m1.UpdateAt,
			Text:       nil,
			Equipments: nil,
		}
		StatusWorksheetArray = append(StatusWorksheetArray, StatusWorksheet1)

	}
	StatusWorksheet2 := structure.StatusWorksheet{
		Status:     constant.WorksheetsStatus2,
		UpdateAt:   &setTime,
		Text:       nil,
		Equipments: nil,
	}
	StatusWorksheetArray = append(StatusWorksheetArray, StatusWorksheet2)

	data := structure.Worksheet{
		QrCodeID:        Worksheet.QrCodeID,
		Text:            Worksheet.Text,
		Type:            Worksheet.Type,
		Ops:             &line.Username,
		OwnerId:         Worksheet.OwnerId,
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

func (ctrl *APIControl) OwnerGetUpdateWorksheet(QrCodeId string) (res structure.GetWorksheet, Error error) {
	qrCode, err := ctrl.access.RDBMS.GetDataQrCodeInfo(QrCodeId)
	if err != nil {
		Error = err
		return
	}

	var responseArray []structure.Worksheet
	Worksheet := structure.Worksheet{}
	ops, err := ctrl.access.RDBMS.GetDataQrCodeOpsByQrCodeID(QrCodeId)
	if err != nil {
		Error = err
		return
	}
	for _, qr := range ops {
		if qr.QrCodeID.String() == QrCodeId {
			err = json2.Unmarshal(qr.Operator, &Worksheet)
			if err != nil {
				Error = err
				return
			}
			var Text *string
			var StatusWorksheetArray []structure.StatusWorksheet

			s1 := len(Worksheet.StatusWorksheet)
			if s1 == 2 {
				for _, m1 := range Worksheet.StatusWorksheet {
					var Equipments []structure.Equipment
					if m1.Status != "" {
						if m1.Status != constant.WorksheetsStatus4 {
							Equipment := structure.Equipment{}
							for _, m2 := range m1.Equipments {
								Equipment = structure.Equipment{
									NameEquipment: m2.NameEquipment,
								}
								Equipments = append(Equipments, Equipment)
							}
						}
					}
					Text = m1.Text
					StatusWorksheet := structure.StatusWorksheet{
						Status:     m1.Status,
						UpdateAt:   m1.UpdateAt,
						Text:       Text,
						Equipments: Equipments,
					}
					StatusWorksheetArray = append(StatusWorksheetArray, StatusWorksheet)
				}
				Ops := *Worksheet.Ops
				data := structure.Worksheet{
					ID:              qr.ID,
					QrCodeID:        Worksheet.QrCodeID,
					Text:            Worksheet.Text,
					Type:            Worksheet.Type,
					Ops:             &Ops,
					StatusWorksheet: StatusWorksheetArray,
				}
				responseArray = append(responseArray, data)
			}
		}
	}

	data := structure.GetWorksheet{
		Info:      qrCode.Info,
		Worksheet: responseArray,
	}

	res = data
	return
}

func (ctrl *APIControl) OwnerUpdateWorksheet(OwnerId, reportId uint, req structure.UpdateWorksheet) (Error error) {
	owner, err := ctrl.access.RDBMS.GetAccount(int(OwnerId))
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

	if len(Worksheet.StatusWorksheet) == 1 {
		Error = errors.New("ยังไม่สามารถจบงานนี้ได้")
		return
	}

	for index, m1 := range Worksheet.StatusWorksheet {
		if index == 0 {
			StatusWorksheet1 := structure.StatusWorksheet{
				Status:     m1.Status,
				UpdateAt:   m1.UpdateAt,
				Text:       nil,
				Equipments: nil,
			}
			StatusWorksheetArray = append(StatusWorksheetArray, StatusWorksheet1)
		}
		if index == 1 {
			StatusWorksheet2 := structure.StatusWorksheet{
				Status:     m1.Status,
				UpdateAt:   m1.UpdateAt,
				Text:       nil,
				Equipments: nil,
			}
			StatusWorksheetArray = append(StatusWorksheetArray, StatusWorksheet2)
		}
		if index == 2 {
			Error = errors.New("งานนี้ถูกยกเลิก หรือ ส่งงานแล้ว")
			return
		}
	}
	var Equipments []structure.Equipment
	for _, em := range req.Equipments {
		Equipment := structure.Equipment{
			NameEquipment: em.NameEquipment,
		}
		Equipments = append(Equipments, Equipment)
	}
	StatusWorksheet3 := structure.StatusWorksheet{
		Status:     constant.WorksheetsStatus3,
		UpdateAt:   &setTime,
		Text:       &req.Text,
		Equipments: Equipments,
	}
	StatusWorksheetArray = append(StatusWorksheetArray, StatusWorksheet3)

	data := structure.Worksheet{
		QrCodeID:        Worksheet.QrCodeID,
		Text:            Worksheet.Text,
		Type:            Worksheet.Type,
		Ops:             &owner.Username,
		OwnerId:         Worksheet.OwnerId,
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
		UserId:   OwnerId,
	}

	err = ctrl.access.RDBMS.UpdateDataQrCodeOps(dataOps)
	if err != nil {
		Error = err
		return
	}
	return
}

func (ctrl *APIControl) OwnerDeleteWorksheet(OwnerId, reportId uint, req structure.UpdateWorksheet) (Error error) {
	owner, err := ctrl.access.RDBMS.GetAccount(int(OwnerId))
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
	for index, m1 := range Worksheet.StatusWorksheet {
		if index == 0 {
			StatusWorksheet1 := structure.StatusWorksheet{
				Status:     m1.Status,
				UpdateAt:   m1.UpdateAt,
				Text:       nil,
				Equipments: nil,
			}
			StatusWorksheetArray = append(StatusWorksheetArray, StatusWorksheet1)
		}
		if index == 1 {
			if m1.Status == constant.WorksheetsStatus2 {

				StatusWorksheet2 := structure.StatusWorksheet{
					Status:     m1.Status,
					UpdateAt:   m1.UpdateAt,
					Text:       nil,
					Equipments: nil,
				}
				StatusWorksheetArray = append(StatusWorksheetArray, StatusWorksheet2)
			} else {
				Error = errors.New("งานนี้ถูกยกเลิก หรือ ส่งงานแล้ว")
				return
			}
		}
		if index == 2 {
			Error = errors.New("งานนี้ถูกยกเลิก หรือ ส่งงานแล้ว")
			return
		}

	}
	StatusWorksheet3 := structure.StatusWorksheet{
		Status:     constant.WorksheetsStatus4,
		UpdateAt:   &setTime,
		Text:       &req.Text,
		Equipments: nil,
	}
	StatusWorksheetArray = append(StatusWorksheetArray, StatusWorksheet3)

	data := structure.Worksheet{
		QrCodeID:        Worksheet.QrCodeID,
		Text:            Worksheet.Text,
		Type:            Worksheet.Type,
		Ops:             &owner.Username,
		OwnerId:         Worksheet.OwnerId,
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
		UserId:   OwnerId,
	}

	err = ctrl.access.RDBMS.UpdateDataQrCodeOps(dataOps)
	if err != nil {
		Error = err
		return
	}
	return
}

// Ops
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
	opsLine, err := ctrl.access.RDBMS.GetAccountByLineId(lineId)
	if err != nil {
		Error = err
		return
	}

	var responseArray []structure.Worksheet

	ops, err := ctrl.access.RDBMS.GetDataQrCodeOps()
	if err != nil {
		Error = err
		return
	}
	for _, qr := range ops {
		Worksheet := structure.Worksheet{}
		err = json2.Unmarshal(qr.Operator, &Worksheet)
		if err != nil {
			Error = err
			return
		}

		if *opsLine.SubOwnerId == Worksheet.OwnerId {
			var StatusWorksheetArray []structure.StatusWorksheet
			var Ops *string
			Ops = Worksheet.Ops
			if Worksheet.Option == true {
				for _, m1 := range Worksheet.StatusWorksheet {
					var Equipments []structure.Equipment
					var Text *string
					if m1.Status != "" {
						if m1.Status != constant.WorksheetsStatus4 {
							Equipment := structure.Equipment{}
							for _, m2 := range m1.Equipments {
								Equipment = structure.Equipment{
									NameEquipment: m2.NameEquipment,
								}
								Equipments = append(Equipments, Equipment)
							}
							Text = m1.Text
						}
					}

					StatusWorksheet := structure.StatusWorksheet{
						Status:     m1.Status,
						UpdateAt:   m1.UpdateAt,
						Text:       Text,
						Equipments: Equipments,
					}
					StatusWorksheetArray = append(StatusWorksheetArray, StatusWorksheet)
				}
				data := structure.Worksheet{
					ID:              qr.ID,
					Option:          Worksheet.Option,
					QrCodeID:        Worksheet.QrCodeID,
					Text:            Worksheet.Text,
					Type:            Worksheet.Type,
					Ops:             Ops,
					StatusWorksheet: StatusWorksheetArray,
				}
				responseArray = append(responseArray, data)

			} else {
				for _, m1 := range Worksheet.StatusWorksheet {
					var Equipments []structure.Equipment
					var Text *string
					if m1.Status != "" {
						if m1.Status != constant.WorksheetsStatus4 {
							Equipment := structure.Equipment{}
							for _, m2 := range m1.Equipments {
								Equipment = structure.Equipment{
									NameEquipment: m2.NameEquipment,
								}
								Equipments = append(Equipments, Equipment)
							}
							Text = m1.Text
						}
					}

					StatusWorksheet := structure.StatusWorksheet{
						Status:     m1.Status,
						UpdateAt:   m1.UpdateAt,
						Text:       Text,
						Equipments: Equipments,
					}
					StatusWorksheetArray = append(StatusWorksheetArray, StatusWorksheet)
				}

				if Ops != nil && *Ops == opsLine.Username {
					data := structure.Worksheet{
						ID:              qr.ID,
						Option:          Worksheet.Option,
						QrCodeID:        Worksheet.QrCodeID,
						Text:            Worksheet.Text,
						Type:            Worksheet.Type,
						Ops:             Ops,
						StatusWorksheet: StatusWorksheetArray,
					}
					responseArray = append(responseArray, data)
				}
			}
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
	for index, m1 := range Worksheet.StatusWorksheet {
		if index == 0 {
			StatusWorksheet1 := structure.StatusWorksheet{
				Status:     m1.Status,
				UpdateAt:   m1.UpdateAt,
				Text:       m1.Text,
				Equipments: m1.Equipments,
			}
			StatusWorksheetArray = append(StatusWorksheetArray, StatusWorksheet1)
		}
		if index == 1 {
			StatusWorksheet2 := structure.StatusWorksheet{
				Status:     m1.Status,
				UpdateAt:   m1.UpdateAt,
				Text:       m1.Text,
				Equipments: m1.Equipments,
			}
			StatusWorksheetArray = append(StatusWorksheetArray, StatusWorksheet2)
		}
		if index == 2 {
			StatusWorksheet3 := structure.StatusWorksheet{
				Status:     m1.Status,
				UpdateAt:   m1.UpdateAt,
				Text:       m1.Text,
				Equipments: m1.Equipments,
			}
			StatusWorksheetArray = append(StatusWorksheetArray, StatusWorksheet3)
		}
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
	var Option bool
	ops, err := ctrl.access.RDBMS.GetDataQrCodeOpsByQrCodeID(req.QrCodeID.String())
	for _, m1 := range ops {
		Worksheet := structure.Worksheet{}
		err = json2.Unmarshal(m1.Operator, &Worksheet)
		if err != nil {
			Error = err
			return
		}
		Option = Worksheet.Option
	}

	var QrCodeID uuid.UUID
	var QrCodeRefer uint
	var OwnerId uint
	for _, qr := range QrCode {
		if qr.TemplateName == string(constant.OfficeEquipment) {
			QrCodeID = qr.QrCodeUUID
			QrCodeRefer = qr.ID
			OwnerId = qr.OwnerId
		}
	}

	setTime := time.Now()
	var StatusWorksheetArray []structure.StatusWorksheet
	StatusWorksheet1 := structure.StatusWorksheet{
		Status:     constant.WorksheetsStatus1,
		UpdateAt:   &setTime,
		Text:       nil,
		Equipments: nil,
	}
	StatusWorksheetArray = append(StatusWorksheetArray, StatusWorksheet1)

	Worksheet := structure.Worksheet{
		QrCodeID:        req.QrCodeID,
		Option:          Option,
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
	for _, m1 := range Worksheet.StatusWorksheet {
		if m1.Status == constant.WorksheetsStatus2 {
			Error = errors.New("มีช่างรับงานนี้ไปแล้ว")
			return
		}
		StatusWorksheet1 := structure.StatusWorksheet{
			Status:     m1.Status,
			UpdateAt:   m1.UpdateAt,
			Text:       nil,
			Equipments: nil,
		}
		StatusWorksheetArray = append(StatusWorksheetArray, StatusWorksheet1)

	}
	StatusWorksheet2 := structure.StatusWorksheet{
		Status:     constant.WorksheetsStatus2,
		UpdateAt:   &setTime,
		Text:       nil,
		Equipments: nil,
	}
	StatusWorksheetArray = append(StatusWorksheetArray, StatusWorksheet2)

	data := structure.Worksheet{
		QrCodeID:        Worksheet.QrCodeID,
		Text:            Worksheet.Text,
		Option:          Worksheet.Option,
		Type:            Worksheet.Type,
		Ops:             &line.Username,
		OwnerId:         Worksheet.OwnerId,
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
	ops, err := ctrl.access.RDBMS.GetDataQrCodeOpsByQrCodeID(QrCodeId)
	if err != nil {
		Error = err
		return
	}
	for _, qr := range ops {
		if qr.QrCodeID.String() == QrCodeId {
			err = json2.Unmarshal(qr.Operator, &Worksheet)
			if err != nil {
				Error = err
				return
			}
			var Text *string
			var StatusWorksheetArray []structure.StatusWorksheet

			s1 := len(Worksheet.StatusWorksheet)
			if s1 == 2 {
				for _, m1 := range Worksheet.StatusWorksheet {
					var Equipments []structure.Equipment
					if m1.Status != "" {
						if m1.Status != constant.WorksheetsStatus4 {
							Equipment := structure.Equipment{}
							for _, m2 := range m1.Equipments {
								Equipment = structure.Equipment{
									NameEquipment: m2.NameEquipment,
								}
								Equipments = append(Equipments, Equipment)
							}
						}
					}
					Text = m1.Text
					StatusWorksheet := structure.StatusWorksheet{
						Status:     m1.Status,
						UpdateAt:   m1.UpdateAt,
						Text:       Text,
						Equipments: Equipments,
					}
					StatusWorksheetArray = append(StatusWorksheetArray, StatusWorksheet)
				}
				Ops := *Worksheet.Ops
				data := structure.Worksheet{
					ID:              qr.ID,
					QrCodeID:        Worksheet.QrCodeID,
					Text:            Worksheet.Text,
					Type:            Worksheet.Type,
					Ops:             &Ops,
					StatusWorksheet: StatusWorksheetArray,
				}
				responseArray = append(responseArray, data)
			}
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

	if len(Worksheet.StatusWorksheet) == 1 {
		Error = errors.New("ยังไม่สามารถจบงานนี้ได้")
		return
	}

	for index, m1 := range Worksheet.StatusWorksheet {
		if index == 0 {
			StatusWorksheet1 := structure.StatusWorksheet{
				Status:     m1.Status,
				UpdateAt:   m1.UpdateAt,
				Text:       nil,
				Equipments: nil,
			}
			StatusWorksheetArray = append(StatusWorksheetArray, StatusWorksheet1)
		}
		if index == 1 {
			StatusWorksheet2 := structure.StatusWorksheet{
				Status:     m1.Status,
				UpdateAt:   m1.UpdateAt,
				Text:       nil,
				Equipments: nil,
			}
			StatusWorksheetArray = append(StatusWorksheetArray, StatusWorksheet2)
		}
		if index == 2 {
			Error = errors.New("งานนี้ถูกยกเลิก หรือ ส่งงานแล้ว")
			return
		}
	}
	var Equipments []structure.Equipment
	for _, em := range req.Equipments {
		Equipment := structure.Equipment{
			NameEquipment: em.NameEquipment,
		}
		Equipments = append(Equipments, Equipment)
	}
	StatusWorksheet3 := structure.StatusWorksheet{
		Status:     constant.WorksheetsStatus3,
		UpdateAt:   &setTime,
		Text:       &req.Text,
		Equipments: Equipments,
	}
	StatusWorksheetArray = append(StatusWorksheetArray, StatusWorksheet3)

	data := structure.Worksheet{
		QrCodeID:        Worksheet.QrCodeID,
		Text:            Worksheet.Text,
		Option:          Worksheet.Option,
		Type:            Worksheet.Type,
		Ops:             &line.Username,
		OwnerId:         Worksheet.OwnerId,
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
	for index, m1 := range Worksheet.StatusWorksheet {
		if index == 0 {
			StatusWorksheet1 := structure.StatusWorksheet{
				Status:     m1.Status,
				UpdateAt:   m1.UpdateAt,
				Text:       nil,
				Equipments: nil,
			}
			StatusWorksheetArray = append(StatusWorksheetArray, StatusWorksheet1)
		}
		if index == 1 {
			if m1.Status == constant.WorksheetsStatus2 {

				StatusWorksheet2 := structure.StatusWorksheet{
					Status:     m1.Status,
					UpdateAt:   m1.UpdateAt,
					Text:       nil,
					Equipments: nil,
				}
				StatusWorksheetArray = append(StatusWorksheetArray, StatusWorksheet2)
			} else {
				Error = errors.New("งานนี้ถูกยกเลิก หรือ ส่งงานแล้ว")
				return
			}
		}
		if index == 2 {
			Error = errors.New("งานนี้ถูกยกเลิก หรือ ส่งงานแล้ว")
			return
		}

	}
	StatusWorksheet3 := structure.StatusWorksheet{
		Status:     constant.WorksheetsStatus4,
		UpdateAt:   &setTime,
		Text:       &req.Text,
		Equipments: nil,
	}
	StatusWorksheetArray = append(StatusWorksheetArray, StatusWorksheet3)

	data := structure.Worksheet{
		QrCodeID:        Worksheet.QrCodeID,
		Text:            Worksheet.Text,
		Option:          Worksheet.Option,
		Type:            Worksheet.Type,
		Ops:             &line.Username,
		OwnerId:         Worksheet.OwnerId,
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
