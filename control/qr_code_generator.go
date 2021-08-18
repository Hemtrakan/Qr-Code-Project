package control

import (
	"archive/zip"
	"encoding/json"
	"errors"
	uuid2 "github.com/gofrs/uuid"
	"github.com/yeqown/go-qrcode"
	"gorm.io/datatypes"
	"io"
	"os"
	"qrcode/access/constant"
	rdbmsstructure "qrcode/access/rdbms/structure"
	"qrcode/present/structure"
	"qrcode/utility"
	"strconv"
)

func (ctrl *APIControl) GetQrCodeById(OwnerId int) (response []structure.GetQrCode, Error error) {
	var getQrCodeArray []structure.GetQrCode
	data, err := ctrl.access.RDBMS.GetQrCodeByOwnerId(OwnerId)
	if err != nil {
		Error = err
		return
	}
	if len(data) == 0 {
		Error = errors.New("this data does not exist")
		return
	}
	for _, res := range data {
		resGetQrCode := structure.GetQrCode{
			OwnerId:      res.OwnerId,
			TemplateName: res.TemplateName,
			QrCodeId:     res.QrCodeUUID.String(),
			CodeName:     res.Code,
		}
		getQrCodeArray = append(getQrCodeArray, resGetQrCode)
	}
	response = getQrCodeArray
	return
}

func (ctrl *APIControl) GetDataQrCode(QrCodeId string) (response structure.GetDataQrCode, Error error) {
	data, err := ctrl.access.RDBMS.GetDataQrCode(QrCodeId)
	if err != nil {
		Error = err
		return
	}
	if data.ID == 0 {
		Error = errors.New("this data does not exist")
		return
	}
	response = structure.GetDataQrCode{
		QrCodeId:     data.QrCodeUUID.String(),
		Info:         data.Info,
		Ops:          data.Ops,
		HistoryInfo:  data.HistoryInfo,
		OwnerId:      int(data.OwnerId),
		TemplateName: data.TemplateName,
	}
	return
}

func (ctrl *APIControl) DeleteQrCode(req structure.DelQrCode) (Error error) {
	for _, del := range req.QrCodeId {
		data, err := ctrl.access.RDBMS.GetDataQrCode(del)
		if err != nil {
			Error = err
			return
		}
		if data.QrCodeUUID.String() != del {
			Error = errors.New("don't have this qr code")
		}
		err = ctrl.access.RDBMS.DeleteQrCode(del)
		if err != nil {
			Error = err
			return
		}
	}
	return
}

func (ctrl *APIControl) CreateQrCode(req structure.GenQrCode) (Error error) {
	ownerId := int(req.OwnerId)
	data, err := ctrl.access.RDBMS.GetAccount(ownerId)
	if err != nil {
		Error = err
		return
	}
	if data.ID == 0 {
		Error = errors.New("there is no owner of this id in the system")
		return
	}
	if data.Role != string(constant.Owner) {
		Error = errors.New("invalid user rights")
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
	err = json.Unmarshal(byteInfo, &structureInfo)
	if err != nil {
		Error = err
		return
	}
	// สร้าง QR-Code
	count, err := ctrl.access.RDBMS.GetQrCode(req.OwnerId, req.TemplateName)
	if err != nil {
		Error = err
		return
	}
	counts := len(count)
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
			Ops:          datatypes.JSON(""),
			HistoryInfo:  datatypes.JSON(""),
			QrCodeUUID:   uuid,
			Code:         req.CodeName + "-" + number,
		}

		err = ctrl.access.RDBMS.CreateQrCode(save)
		if err != nil {
			Error = err
			return
		}
	}
	return
}

func (ctrl *APIControl) AddFileZipById(req structure.FileZip) (file string, Error error) {
	var arrayFileName []structure.ArrayFileName
	data, _ := ctrl.access.RDBMS.GetAccount(req.OwnerId)
	if data.ID == 0 {
		Error = errors.New("there is no owner of this id in the system")
		return
	}
	if data.Role != string(constant.Owner) {
		Error = errors.New("invalid user rights")
		return
	}
	for _, QrCodeId := range req.QrCodeId {
		data, err := ctrl.access.RDBMS.GetQrCodeByQrCodeId(req.OwnerId, QrCodeId)
		if err != nil {
			Error = err
			return
		}
		if data.ID == 0 {
			Error = errors.New("this qr-code does not exist")
			for _, QrCodeId := range req.QrCodeId {
				data, err := ctrl.access.RDBMS.GetQrCodeByQrCodeId(req.OwnerId, QrCodeId)
				if err != nil {
					Error = errors.New("this qr-code does not exist")
					return
				}
				filename := data.Code
				path := string(constant.SaveFileLocationQrCode) + "/" + filename + ".PNG"
				err = os.Remove(path)
				if err != nil {
					Error = errors.New("this qr-code does not exist")
					return
				}
			}
			return
		}
		filename := data.Code
		qrc, err := qrcode.New(constant.Http + "/" + data.QrCodeUUID.String())
		if err != nil {
			Error = err
			return
		}
		path := string(constant.SaveFileLocationQrCode) + "/" + filename + ".PNG"
		// save file
		if err = qrc.Save(path); err != nil {
			Error = err
			return
		}
		files := structure.ArrayFileName{
			FileName: filename,
		}
		arrayFileName = append(arrayFileName, files)
	}
	output := "zipfile/" + req.FileZip + ".zip"
	if err := ZipFilesById(output, arrayFileName); err != nil {
		Error = err
		return
	}
	for _, QrCodeId := range req.QrCodeId {
		data, err := ctrl.access.RDBMS.GetQrCodeByQrCodeId(req.OwnerId, QrCodeId)
		if err != nil {
			Error = err
			return
		}
		filename := data.Code
		path := string(constant.SaveFileLocationQrCode) + "/" + filename + ".PNG"
		err = os.Remove(path)
		if err != nil {
			Error = err
			return
		}
	}
	file = output
	return
}

func (ctrl *APIControl) AddFileZipByTemplateName(req structure.FileZipByTemplateName) (file string, Error error) {
	ownerId := int(req.OwnerId)
	dataOwnerId, err := ctrl.access.RDBMS.GetAccount(ownerId)
	if err != nil {
		Error = err
		return
	}
	if dataOwnerId.ID == 0 {
		Error = errors.New("there is no owner of this id in the system")
		return
	}
	if dataOwnerId.Role != string(constant.Owner) {
		Error = errors.New("invalid user rights")
		return
	}
	data, err := ctrl.access.RDBMS.GetQrCode(req.OwnerId, req.TemplateName)
	if err != nil {
		Error = err
		return
	}
	if len(data) == 0 {
		Error = errors.New("no new qrcode")
		return
	}
	err = os.RemoveAll(string(constant.SaveFileLocationZipFile))
	if err != nil {
		Error = err
		return
	}
	err = os.Mkdir(string(constant.SaveFileLocationQrCode), 0755)
	err = os.Mkdir(string(constant.SaveFileLocationZipFile), 0755)
	if err != nil {
		Error = err
		return
	}
	var path string
	var arrayFileName []structure.ArrayFileName
	for _, res := range data {
		qrc, err := qrcode.New(constant.Http + "/" + res.QrCodeUUID.String())
		if err != nil {
			Error = err
			return
		}
		filename := res.Code
		path = string(constant.SaveFileLocationQrCode) + "/" + filename + ".PNG"
		// save file
		if err = qrc.Save(path); err != nil {
			Error = err
			return
		}
		files := structure.ArrayFileName{
			FileName: filename,
		}
		arrayFileName = append(arrayFileName, files)
	}
	output := string(constant.SaveFileLocationZipFile) + "/" + req.FileZip + ".zip"
	if err := ZipFilesByTemplateName(output, arrayFileName); err != nil {
		Error = err
		return
	}

	file = output
	err = os.RemoveAll(string(constant.SaveFileLocationQrCode))
	if err != nil {
		Error = err
		return
	}
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
		pathFile := string(constant.SaveFileLocationQrCode) + "/" + file.FileName + ".PNG"
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
		pathFile := string(constant.SaveFileLocationQrCode) + "/" + file.FileName + ".PNG"
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

//func (ctrl *APIControl) GenerateQrCode(req *structure.GenerateQrCode) (Error error) {
//	QrCode := rdbmsstructure.Qrcode{
//		QrCode:     req.QrCode,
//		OwnersId:   req.OwnersId,
//		TeamPageId: req.TeamPageId,
//		Location:   rdbmsstructure.Location{
//			Country:     req.Location.Country,
//			Address:     req.Location.Address,
//			SubDistrict: req.Location.SubDistrict,
//			District:    req.Location.District,
//			Province:    req.Location.Province,
//			Zipcode:     req.Location.Zipcode,
//			XCoordinate: req.Location.XCoordinate,
//			YCoordinate: req.Location.YCoordinate,
//		},
//	}
//	err := ctrl.access.RDBMS.GenerateQrCode(QrCode)
//	if err != nil {
//		Error = err
//		return
//	}
//	data, err := ctrl.access.RDBMS.GetIdTeamPage(req.TeamPageId)
//	if err != nil {
//		Error = err
//		return
//	}
//	//qrc, err := qrcode.New(constant.Http+ "/api/" +"getByIdTeamPage/" + strconv.FormatUint(uint64(req.TeamPageId),10))
//	qrc, err := qrcode.New(constant.Http+ "/"+ data.UUID.String())
//	if err != nil {
//		fmt.Printf("could not generate QRCode: %v", err)
//	}
//
//	path := string(constant.SaveFileLocationQrCode) + "/" + req.QrCode + strconv.FormatUint(uint64(req.TeamPageId),10) + ".PNG"
//	// save file
//	if err = qrc.Save(path); err != nil {
//		fmt.Printf("could not save image: %v", err)
//	}
//	return err
//}
