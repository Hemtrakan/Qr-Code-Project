package control

import (
	"archive/zip"
	"io"
	"os"
	"qrcode/present/structure"
)

//func (ctrl *APIControl) GenQrCodeByName(files structure.GenQr) (file string ,Error error){
//	output := "zipfile/" + files.FileZip + ".zip"
//	if err := ZipFiles(output, files.Filename); err != nil {
//		panic(err)
//	}
//	file = output
//	return
//}

func (ctrl *APIControl) GenQrCode(files structure.GenQr) (file string ,Error error){
	output := "zipfile/" + files.FileZip + ".zip"
	if err := ZipFiles(output, files.Filename); err != nil {
		panic(err)
	}
	file = output
	return
}

func ZipFiles(filename string, files []string) error {
	newZipFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer newZipFile.Close()

	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()

	// Add files to zip
	for _, file := range files {
		pathFile := "fileqrcode/" + file + ".PNG"
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