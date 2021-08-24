package control

import (
	"archive/zip"
	"encoding/json"
	"errors"
	uuid2 "github.com/gofrs/uuid"
	"github.com/yeqown/go-qrcode"
	"gorm.io/datatypes"
	"image"
	"image/png"
	"io"
	"os"
	"qrcode/access/constant"
	rdbmsstructure "qrcode/access/rdbms/structure"
	"qrcode/present/structure"
	"qrcode/utility"
	"strconv"
	"strings"
)

func (ctrl *APIControl) GetAllQrCode() (response []structure.GetQrCode, Error error) {
	var getQrCodeArray []structure.GetQrCode

	data, err := ctrl.access.RDBMS.GetAllQrCode()
	if err != nil {
		Error = err
		return
	}
	for _, res := range data {
		dataOwner,err:= ctrl.GetAccount(int(res.OwnerId))
		if err != nil {
			Error = err
			return
		}
		resGetQrCode := structure.GetQrCode{
			OwnerId:      res.OwnerId,
			OwnerName:    dataOwner.FirstName + " " + dataOwner.LastName,
			CreatedAt:    res.CreatedAt,
			UpdatedAt:    res.UpdatedAt,
			TemplateName: res.TemplateName,
			QrCodeId:     res.QrCodeUUID.String(),
			CodeName:     res.Code + "-" + res.Count,
		}
		getQrCodeArray = append(getQrCodeArray, resGetQrCode)
	}
	response = getQrCodeArray
	return
}


func (ctrl *APIControl) GetQrCodeById(OwnerId int) (response []structure.GetQrCode, Error error) {
	var getQrCodeArray []structure.GetQrCode
	check, err := ctrl.access.RDBMS.CheckAccountId(uint(OwnerId))
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
			OwnerId:      res.OwnerId,
			CreatedAt:    res.CreatedAt,
			UpdatedAt:    res.UpdatedAt,
			TemplateName: res.TemplateName,
			QrCodeId:     res.QrCodeUUID.String(),
			CodeName:     res.Code + "-" + res.Count,
		}
		getQrCodeArray = append(getQrCodeArray, resGetQrCode)
	}
	response = getQrCodeArray
	return
}

func (ctrl *APIControl) GetDataQrCode(QrCodeId string) (response structure.GetDataQrCode, Error error) {
	data, err := ctrl.access.RDBMS.GetDataQrCode(QrCodeId)
	if err != nil {
		Error = errors.New("ไม่มี QrCode นี้อยู่ในระบบ")
		return
	}
	var HistoryArray []structure.GetHistory
	History, err := ctrl.access.RDBMS.GetHistory(QrCodeId)

	for _, h := range History {
		dataHistory := structure.GetHistory{
			HistoryInfo: h.HistoryInfo,
			UserId:      h.UserId,
			UpdatedAt:   h.UpdatedAt,
		}
		HistoryArray = append(HistoryArray, dataHistory)
	}

	response = structure.GetDataQrCode{
		QrCodeId:     data.QrCodeUUID.String(),
		Info:         data.Info,
		Ops:          data.Ops,
		OwnerId:      int(data.OwnerId),
		TemplateName: data.TemplateName,
		CodeName:     data.Code + "-" + data.Count,
		HistoryInfo:  HistoryArray,
	}

	return
}

func (ctrl *APIControl) DeleteQrCode(req structure.DelQrCode) (Error error) {
	if len(req.QrCodeId) == 0 {
		Error = errors.New("ไม่มี Qr-Code ถูกส่งมา")
	}
	for _, del := range req.QrCodeId {
		_, err := ctrl.access.RDBMS.GetDataQrCode(del)
		if err != nil {
			Error = errors.New("Qr-Code ที่จะลบไม่มีอยู่ในระบบ")
			return
		}
		err = ctrl.access.RDBMS.DeleteQrCode(del)
		if err != nil {
			Error = err
			return
		}
	}
	return
}

func (ctrl *APIControl) CreateQrCode(req structure.GenQrCode) (file string, Error error) {
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
			Ops:          datatypes.JSON(""),
			QrCodeUUID:   uuid,
			Code:         req.CodeName,
			Count:        number,
		}

		err = ctrl.access.RDBMS.CreateQrCode(save)
		if err != nil {
			Error = err
			return
		}
		arrayString = append(arrayString, uuid.String())
	}

	fileZip := structure.FileZip{
		OwnerId:  ownerId,
		FileZip:  "zip",
		QrCodeId: arrayString,
	}

	zip, err := ctrl.AddFileZipById(fileZip)
	file = zip
	return
}

func (ctrl *APIControl) AddFileZipById(req structure.FileZip) (file string, Error error) {
	var arrayFileName []structure.ArrayFileName

	req.FileZip = strings.Trim(req.FileZip, "\t \n")
	if req.FileZip == "" {
		Error = errors.New("FileZip ต้องไม่ว่าง")
		return
	}
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

	err = os.RemoveAll(string(constant.SaveFileLocationZipFile))
	if err != nil {
		Error = err
		return
	}
	err = os.Mkdir(string(constant.SaveFileLocationQrCode), 0755)
	err = os.Mkdir(string(constant.SaveFileLocationZipFile), 0755)
	err = os.Mkdir(string(constant.SaveFileTextLocation), 0755)
	if err != nil {
		Error = err
		return
	}
	for _, QrCodeId := range req.QrCodeId {
		data, err := ctrl.access.RDBMS.GetQrCodeByQrCodeId(req.OwnerId, QrCodeId)
		if err != nil {
			Error = errors.New("ไม่มี Qr-Code ที่นี้อยู่ในระบบ")
			err = os.RemoveAll(string(constant.SaveFileLocationZipFile))
			if err != nil {
				Error = err
				return
			}
			err = os.RemoveAll(string(constant.SaveFileLocationQrCode))
			if err != nil {
				Error = err
				return
			}
			err = os.RemoveAll(string(constant.SaveFileLocationZipFile))
			if err != nil {
				Error = err
				return
			}
			return
		}

		filename := data.Code + "-" + data.Count
		text := string(constant.SaveFileTextLocation) + "/" + filename + ".PNG"
		f, err := os.Create(text)
		if err != nil {
			Error = err
			return
		}
		defer f.Close()
		img := image.NewRGBA(image.Rect(0, 0, 180, 180))
		utility.AddLabel(img, 50, 50, filename)
		if err := png.Encode(f, img); err != nil {
			panic(err)
		}
		// gen QrCode

		//logo := "./logo.jpeg"
		//col := color.RGBA{150, 230, 148, 255}
		qrc, err := qrcode.New(constant.Http+"/"+data.QrCodeUUID.String(),
			qrcode.WithLogoImageFilePNG(text),
			//qrcode.WithBgColor(col),
		)
		if err != nil {
			Error = err
			return
		}
		// save file
		path := string(constant.SaveFileLocationQrCode) + "/" + filename + ".PNG"
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
	if err := ZipFilesById(output, arrayFileName); err != nil {
		Error = err
		return
	}
	err = os.RemoveAll(string(constant.SaveFileLocationQrCode))
	if err != nil {
		Error = err
		return
	}
	err = os.RemoveAll(string(constant.SaveFileTextLocation))
	if err != nil {
		Error = err
		return
	}
	file = output
	return
}

func (ctrl *APIControl) AddFileZipByTemplateName(req structure.FileZipByTemplateName) (file string, Error error) {
	req.TemplateName = strings.Trim(req.TemplateName, "\t \n")
	if req.TemplateName == "" {
		Error = errors.New("TemplateName ต้องไม่ว่าง")
		return
	}
	req.FileZip = strings.Trim(req.FileZip, "\t \n")
	if req.FileZip == "" {
		Error = errors.New("FileZip ต้องไม่ว่าง")
		return
	}
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
		filename := res.Code + "-" + res.Count
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
