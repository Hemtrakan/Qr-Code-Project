package control

import (
	"archive/zip"
	"encoding/json"
	"errors"
	uuid2 "github.com/gofrs/uuid"
	"github.com/yeqown/go-qrcode"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"io"
	"os"
	"qrcode/access/constant"
	rdbmsstructure "qrcode/access/rdbms/structure"
	"qrcode/present/structure"
	"qrcode/utility"
	"strconv"
	"strings"
	"time"
)

// todo ส่วนของการ เพิ่มข้อมูล แก้ไขข้อมูล ลง Qr-Code

func (ctrl *APIControl) InsertDataQrCode(req *structure.InsertDataQrCode) (Error error) {
	check, err := ctrl.access.RDBMS.GetQrCodeByQrCodeId(int(req.OwnerId), req.QrCodeId.String())
	if err != nil {
		Error = errors.New("ไม่พบ QrCode นี้อยู่ในระบบ")
		return
	}
	if check.TemplateName != "" {
		Error = errors.New("QrCode ได้ถูกตั้งค่า Template แล้ว")
		return
	}

	b, err := json.Marshal(req.Info)
	if err != nil {
		Error = err
		return
	}

	data := rdbmsstructure.QrCode{
		TemplateName: req.TemplateName,
		Info:         datatypes.JSON(b),
		QrCodeUUID:   req.QrCodeId,
		First:        true,
	}

	err = ctrl.access.RDBMS.InsertDataQrCodeById(data)
	if err != nil {
		Error = err
		return
	}
	return
}

func (ctrl *APIControl) UpdateDataQrCode(req *structure.UpdateDataQrCode) (Error error) {
	OldInfo, err := ctrl.access.RDBMS.GetQrCodeByQrCodeId(int(req.OwnerId), req.QrCodeId.String())
	if err != nil {
		Error = errors.New("ไม่พบ QrCode นี้อยู่ในระบบ")
		return
	}

	var UserId uint
	if req.LineUserId != "" {
		ops, err := ctrl.access.RDBMS.GetAccountByLineId(req.LineUserId)
		if err != nil {
			Error = err
			return
		}
		UserId = ops.ID
	}else {
		UserId = req.OwnerId
	}
	infoJson, err := json.Marshal(req.Info)
	if err != nil {
		Error = err
		return
	}
	infoQr := rdbmsstructure.QrCode{
		Info:       datatypes.JSON(infoJson),
		QrCodeUUID: req.QrCodeId,
		First:      true,
	}

	HistoryJson, err := json.Marshal(OldInfo.Info)
	if err != nil {
		Error = err
		return
	}

	HistoryInfo := rdbmsstructure.HistoryInfo{
		Model: gorm.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		QrCodeID:    req.QrCodeId,
		HistoryInfo: datatypes.JSON(HistoryJson),
		UserId:      UserId, // id คนที่มาอัพเดทข้อมูล
		QrCodeRefer: OldInfo.ID,
	}

	err = ctrl.access.RDBMS.UpdateDataQrCode(infoQr, HistoryInfo)
	if err != nil {
		Error = err
		return
	}

	return
}

func (ctrl *APIControl) UpdateHistoryInfoDataQrCode(req *structure.UpdateHistoryInfoDataQrCode) (Error error) {
	check, err := ctrl.access.RDBMS.GetQrCodeByQrCodeId(int(req.OwnerId), req.QrCodeId.String())
	if err != nil {
		Error = errors.New("ไม่พบ QrCode นี้อยู่ในระบบ")
		return
	}

	b, err := json.Marshal(req.HistoryInfo)
	if err != nil {
		Error = err
		return
	}

	data := rdbmsstructure.HistoryInfo{
		Model: gorm.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		QrCodeID:    req.QrCodeId,
		HistoryInfo: datatypes.JSON(b),
		UserId:      req.UserId, // id คนที่มาอัพเดทข้อมูล
		QrCodeRefer: check.ID,
	}

	err = ctrl.access.RDBMS.UpdateHistoryInfoQrCodeById(data)
	if err != nil {
		Error = err
		return
	}
	return
}

func (ctrl *APIControl) UpdateOpsDataQrCode(req *structure.UpdateOpsDataQrCode) (Error error) {
	check, err := ctrl.access.RDBMS.GetQrCodeByQrCodeId(int(req.OwnerId), req.QrCodeId.String())
	if err != nil {
		Error = errors.New("ไม่พบ QrCode นี้อยู่ในระบบ")
		return
	}

	b, err := json.Marshal(req.Ops)
	if err != nil {
		Error = err
		return
	}

	data := rdbmsstructure.Ops{
		Model: gorm.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		QrCodeID:    req.QrCodeId,
		Operator:    datatypes.JSON(b),
		UserId:      req.UserId, // id คนที่มาอัพเดทข้อมูล
		QrCodeRefer: check.ID,
	}

	err = ctrl.access.RDBMS.UpdateOpsQrCodeById(data)
	if err != nil {
		Error = err
		return
	}
	return
}

func (ctrl *APIControl) GetDataQrCode(QrCodeId string) (response structure.GetDataQrCode, Error error) {
	data, err := ctrl.access.RDBMS.GetDataQrCode(QrCodeId)
	if err != nil {
		Error = errors.New("ไม่มี QrCode นี้อยู่ในระบบ")
		return
	}
	var opsArray []structure.GetOps
	var HistoryArray []structure.GetHistory
	for _, qr := range data {
		res, _ := ctrl.access.RDBMS.GetAccount(int(qr.OwnerId))
		for _, DataOps := range qr.DataOps {
			User, _ := ctrl.access.RDBMS.GetAccount(int(DataOps.UserId))
			var Username string
			var Role string
			if User.Username == "" {
				Username = "ผู้ใช้งานทั้วไป"
			} else {
				Username = User.Username
			}
			if User.Role == "" {
				Role = "Viewer"
			} else {
				Role = User.Role
			}
			ops := structure.GetOps{
				Ops:  DataOps.Operator,
				User: Username,
				Role: Role,
			}
			opsArray = append(opsArray, ops)
		}

		for _, DataHistory := range qr.DataHistory {
			User, _ := ctrl.access.RDBMS.GetAccount(int(DataHistory.UserId))
			History := structure.GetHistory{
				HistoryInfo: DataHistory.HistoryInfo,
				User:        User.Username,
				UpdatedAt:   DataHistory.CreatedAt,
				Role:        User.Role,
			}
			HistoryArray = append(HistoryArray, History)
		}

		response = structure.GetDataQrCode{
			QrCodeId:     qr.QrCodeUUID.String(),
			Info:         qr.Info,
			OwnerId:      int(qr.OwnerId),
			TemplateName: qr.TemplateName,
			CodeName:     qr.Code + "-" + qr.Count,
			HistoryInfo:  HistoryArray,
			Ops:          opsArray,
			OwnerName:    res.FirstName + " " + res.LastName,
		}
	}
	return
}

//  todo ส่วนของการ CRUD QrCode

func (ctrl *APIControl) CreateQrCode(req structure.GenQrCode) (Error error) {
	req.CodeName = strings.Trim(req.CodeName, "\t \n")
	req.TemplateName = strings.Trim(req.TemplateName, "\t \n")
	if !(len(req.CodeName) <= 20) {
		Error = errors.New("CodeName ต้องไม่เกิน 20 ตัว")
		return
	}
	if req.CodeName == "" {
		Error = errors.New("CodeName ต้องไม่ว่าง")
		return
	}

	ownerId := int(req.OwnerId)
	data, err := ctrl.access.RDBMS.GetAccount(ownerId)
	if err != nil {
		Error = errors.New("ไม่มีผู้ใช้คนนี้อยู่ในระบบ")
		return
	}
	if data.Role != string(constant.Owner) {
		Error = errors.New("ผู้ใช้คนนี้ไม่มีสิทธิ์ในการสร้าง QR-Code")
		return
	}
	structureInfo, err := utility.CheckTemplate(req.TemplateName)
	if err != nil {
		Error = err
		return
	}
	byteInfo, err := json.Marshal(structureInfo)
	if err != nil {
		Error = err
		return
	}
	// สร้าง QR-Code
	count, err := ctrl.access.RDBMS.CountCode(req.OwnerId, req.TemplateName, req.CodeName)
	if err != nil {
		Error = err
		return
	}

	var counts = 0
	check, err := ctrl.access.RDBMS.CheckCode(req.OwnerId, req.TemplateName, req.CodeName)
	if err != nil {
		Error = err
		return
	}
	if check.Code == req.CodeName {
		counts = len(count)
	}

	var arrayString []string
	for i := 0 + 1; i <= req.Amount; i++ {
		uuid, err := uuid2.NewV4()
		if err != nil {
			Error = err
			return
		}
		number := strconv.Itoa(counts + i)

		save := rdbmsstructure.QrCode{
			OwnerId:      req.OwnerId,
			TemplateName: req.TemplateName,
			Info:         datatypes.JSON(byteInfo),
			QrCodeUUID:   uuid,
			Code:         req.CodeName,
			Count:        number,
			First:        false,
			Active:       true,
		}

		err = ctrl.access.RDBMS.CreateQrCode(save)
		if err != nil {
			Error = err
			return
		}
		arrayString = append(arrayString, uuid.String())
	}

	return
}

func (ctrl *APIControl) GetAllQrCode() (response []structure.GetQrCode, Error error) {
	var getQrCodeArray []structure.GetQrCode

	data, err := ctrl.access.RDBMS.GetAllQrCode()
	if err != nil {
		Error = errors.New("ไม่มีข้อมูล")
		return
	}
	for _, res := range data {
		owner, err := ctrl.access.RDBMS.GetAccount(int(res.OwnerId))
		if err != nil {
			Error = err
			return
		}
		resGetQrCode := structure.GetQrCode{
			OwnerId:       res.OwnerId,
			OwnerUsername: owner.Username,
			CreatedAt:     res.CreatedAt,
			UpdatedAt:     res.UpdatedAt,
			TemplateName:  res.TemplateName,
			QrCodeId:      res.QrCodeUUID.String(),
			CodeName:      res.Code + "-" + res.Count,
			URL:           ctrl.access.ENV.URLQRCode + res.QrCodeUUID.String(),
			Active:        res.Active,
		}
		getQrCodeArray = append(getQrCodeArray, resGetQrCode)
	}
	response = getQrCodeArray

	return
}

func (ctrl *APIControl) GetQrCodeById(OwnerId int) (response []structure.GetQrCode, Error error) {
	var getQrCodeArray []structure.GetQrCode
	check, err := ctrl.access.RDBMS.GetAccount(OwnerId)
	if err != nil {
		Error = errors.New("ไม่มีผู้ใช้คนนี้ในระบบ")
		return
	}
	if !(check.Role == string(constant.Owner)) {
		Error = errors.New("สิทธิ์ผู้ใช้งานไม่ถูกต้อง")
		return
	}
	data, err := ctrl.access.RDBMS.GetQrCodeByOwnerId(OwnerId)
	if err != nil {
		Error = err
		return
	}
	for _, res := range data {
		resGetQrCode := structure.GetQrCode{
			OwnerId:       res.OwnerId,
			OwnerUsername: check.Username,
			CreatedAt:     res.CreatedAt,
			UpdatedAt:     res.UpdatedAt,
			TemplateName:  res.TemplateName,
			QrCodeId:      res.QrCodeUUID.String(),
			CodeName:      res.Code + "-" + res.Count,
			URL:           ctrl.access.ENV.URLQRCode + res.QrCodeUUID.String(),
			Active:        res.Active,
		}
		getQrCodeArray = append(getQrCodeArray, resGetQrCode)
	}
	response = getQrCodeArray
	return
}

func (ctrl *APIControl) UpdateStatusQrCode(QrCodeId string, req structure.StatusQrCode) (Error error) {
	res, err := ctrl.access.RDBMS.GetDataQrCode(QrCodeId)
	if err != nil {
		Error = errors.New("Qr-Code ที่จะเปลี่ยนสถานะไม่มีอยู่ในระบบ")
		return
	}
	for _, data := range res {
		Qr := rdbmsstructure.QrCode{
			QrCodeUUID: data.QrCodeUUID,
			Active:     *req.Active,
		}
		err = ctrl.access.RDBMS.UpdateQrCodeActive(Qr)
		if err != nil {
			Error = err
			return
		}
	}
	return
}

func (ctrl *APIControl) DeleteQrCode(req structure.DelQrCode) (Error error) {
	if len(req.QrCodeId) == 0 {
		Error = errors.New("ไม่มี Qr-Code ถูกส่งมา")
	}
	for _, del := range req.QrCodeId {
		id, err := ctrl.access.RDBMS.GetDataQrCode(del)
		if err != nil {
			Error = errors.New("Qr-Code ที่จะลบไม่มีอยู่ในระบบ")
			return
		}
		for _, m1 := range id {
			err = ctrl.access.RDBMS.DeleteQrCode(m1.ID)
			if err != nil {
				Error = err
				return
			}
		}
	}
	return
}

//  todo Exposed file QrCode

func (ctrl *APIControl) AddFileZipById(req structure.FileZip) (file string, Error error) {
	var arrayFileName []structure.ArrayFileName
	if len(req.QrCodeId) == 0 {
		Error = errors.New("QrCodeId ต้องไม่ว่าง")
		return
	}
	check, err := ctrl.access.RDBMS.GetAccount(req.OwnerId)
	if err != nil {
		Error = errors.New("ไม่มีผู้ใช้คนนี้อยู่ในระบบ")
		return
	}
	if check.Role != string(constant.Owner) {
		Error = errors.New("ผู้ใช้คนนี้ไม่มีสิทธิ์ในการสร้าง QR-Code")
		return
	}
	for _, QrCodeId := range req.QrCodeId {
		data, err := ctrl.access.RDBMS.GetQrCodeByQrCodeId(req.OwnerId, QrCodeId)
		if err != nil {
			Error = errors.New("ไม่มี Qr-Code ที่นี้อยู่ในระบบ")
			return
		}
		filename := data.Code + "-" + data.Count
		//todo สร้าง QrCode
		qrc, err := qrcode.New(ctrl.access.ENV.URLQRCode + data.QrCodeUUID.String())
		if err != nil {
			Error = err
			return
		}
		// save file
		pathQr := string(constant.SaveFileLocationQrCode) + "/" + filename + ".jpg"
		if err = qrc.Save(pathQr); err != nil {
			Error = err
			return
		}

		// todo เขียนข้อความ
		ResQr := string(constant.SaveFileLocationExposed) + "/" + filename + ".png"
		RequestTextOnImg := utility.Request{
			BgImgPath: pathQr,
			Text:      filename,
			PathSave:  ResQr,
		}
		err = utility.TextOnImg(RequestTextOnImg)
		if err != nil {
			Error = err
			return
		}

		files := structure.ArrayFileName{
			FileName: filename,
		}
		arrayFileName = append(arrayFileName, files)
	}
	output := string(constant.SaveFileLocationZipFile) + "/" + "FileZip.zip"
	if err := ZipFilesById(output, arrayFileName); err != nil {
		Error = err
		return
	}
	file = output
	return
}

func (ctrl *APIControl) AddFileZipByOwner(req structure.FileZipByOwner) (file string, Error error) {
	ownerId := int(req.OwnerId)
	dataOwnerId, err := ctrl.access.RDBMS.GetAccount(ownerId)
	if err != nil {
		Error = errors.New("ไม่มีผู้ใช้คนนี้อยู่ในระบบ")
		return
	}
	if dataOwnerId.Role != string(constant.Owner) {
		Error = errors.New("ผู้ใช้คนนี้ไม่มีสิทธิ์ในการสร้าง QR-Code")
		return
	}
	data, err := ctrl.access.RDBMS.GetQrCodeByOwnerId(ownerId)
	if err != nil {
		Error = err
		return
	}
	if len(data) == 0 {
		Error = errors.New("ยังไม่สร้าง Qr-Code ใน template ที่ถูกเลือก")
		return
	}
	var pathQr string
	var arrayFileName []structure.ArrayFileName
	for _, res := range data {
		URL := ctrl.access.ENV.URLQRCode + res.QrCodeUUID.String()
		qrc, err := qrcode.New(URL)
		if err != nil {
			Error = err
			return
		}
		filename := res.Code + "-" + res.Count
		pathQr = string(constant.SaveFileLocationQrCode) + "/" + filename + ".jpg"
		// save file
		if err = qrc.Save(pathQr); err != nil {
			Error = err
			return
		}
		ResQr := string(constant.SaveFileLocationExposed) + "/" + filename + ".png"
		RequestTextOnImg := utility.Request{
			BgImgPath: pathQr,
			Text:      filename,
			PathSave:  ResQr,
		}
		err = utility.TextOnImg(RequestTextOnImg)
		if err != nil {
			Error = err
			return
		}
		files := structure.ArrayFileName{
			FileName: filename,
		}
		arrayFileName = append(arrayFileName, files)
	}
	output := string(constant.SaveFileLocationZipFile) + "/" + "FileZip.zip"
	if err := ZipFilesByTemplateName(output, arrayFileName); err != nil {
		Error = err
		return
	}
	file = output
	return
}

func (ctrl *APIControl) AddFileZipByTemplateName(req structure.FileZipByTemplateName) (file string, Error error) {
	req.TemplateName = strings.Trim(req.TemplateName, "\t \n")
	_, err := utility.CheckTemplate(req.TemplateName)
	if err != nil {
		Error = err
		return
	}
	ownerId := int(req.OwnerId)
	dataOwnerId, err := ctrl.access.RDBMS.GetAccount(ownerId)
	if err != nil {
		Error = errors.New("ไม่มีผู้ใช้คนนี้อยู่ในระบบ")
		return
	}
	if dataOwnerId.Role != string(constant.Owner) {
		Error = errors.New("ผู้ใช้คนนี้ไม่มีสิทธิ์ในการสร้าง QR-Code")
		return
	}
	data, err := ctrl.access.RDBMS.GetQrCode(req.OwnerId, req.TemplateName)
	if err != nil {
		Error = err
		return
	}
	if len(data) == 0 {
		Error = errors.New("ยังไม่สร้าง Qr-Code ใน template ที่ถูกเลือก")
		return
	}
	var pathQr string
	var arrayFileName []structure.ArrayFileName
	for _, res := range data {
		qrc, err := qrcode.New(string(ctrl.access.ENV.URLQRCode) + res.QrCodeUUID.String())
		if err != nil {
			Error = err
			return
		}
		filename := res.Code + "-" + res.Count
		pathQr = string(constant.SaveFileLocationQrCode) + "/" + filename + ".jpg"
		// save file
		if err = qrc.Save(pathQr); err != nil {
			Error = err
			return
		}
		ResQr := string(constant.SaveFileLocationExposed) + "/" + filename + ".png"
		RequestTextOnImg := utility.Request{
			BgImgPath: pathQr,
			Text:      filename,
			PathSave:  ResQr,
		}
		err = utility.TextOnImg(RequestTextOnImg)
		if err != nil {
			Error = err
			return
		}

		files := structure.ArrayFileName{
			FileName: filename,
		}
		arrayFileName = append(arrayFileName, files)
	}
	output := string(constant.SaveFileLocationZipFile) + "/" + "FileZip.zip"
	if err := ZipFilesByTemplateName(output, arrayFileName); err != nil {
		Error = err
		return
	}
	file = output
	return
}

func ZipFilesById(filename string, files []structure.ArrayFileName) error {
	newZipFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer newZipFile.Close()

	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()

	// Add files to zip
	for _, file := range files {
		pathFile := string(constant.SaveFileLocationExposed) + "/" + file.FileName + ".png"
		if err = AddFileToZip(zipWriter, pathFile); err != nil {
			return err
		}
	}
	return nil
}

func ZipFilesByTemplateName(filename string, files []structure.ArrayFileName) error {
	newZipFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer newZipFile.Close()

	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()

	// Add files to zip
	for _, file := range files {
		pathFile := string(constant.SaveFileLocationExposed) + "/" + file.FileName + ".png"
		if err = AddFileToZip(zipWriter, pathFile); err != nil {
			return err
		}
	}
	return nil
}

func AddFileToZip(zipWriter *zip.Writer, filename string) error {
	fileToZip, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fileToZip.Close()

	// Get the file information
	info, err := fileToZip.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	// Using FileInfoHeader() above only uses the basename of the file. If we want
	// to preserve the folder structure we can overwrite this with the full path.
	header.Name = filename

	// Change to deflate to gain better compression
	// see http://golang.org/pkg/archive/zip/#pkg-constants
	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, fileToZip)
	return err
}
