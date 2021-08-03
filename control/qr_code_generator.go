package control

import (
	"fmt"
	"github.com/yeqown/go-qrcode"
	"qrcode/access/constant"
	rdbmsstructure "qrcode/access/rdbms/structure"
	"qrcode/present/structure"
	"strconv"
)

func (ctrl *APIControl) GenerateQrCode(req *structure.GenerateQrCode) (Error error) {

	QrCode := rdbmsstructure.Qrcode{
		QrCode:     req.QrCode,
		OwnersId:   req.OwnersId,
		TeamPageId: req.TeamPageId,
		Location:   rdbmsstructure.Location{
			Country:     req.Location.Country,
			Address:     req.Location.Address,
			SubDistrict: req.Location.SubDistrict,
			District:    req.Location.District,
			Province:    req.Location.Province,
			Zipcode:     req.Location.Zipcode,
			XCoordinate: req.Location.XCoordinate,
			YCoordinate: req.Location.YCoordinate,
		},
	}
	err := ctrl.access.RDBMS.GenerateQrCode(QrCode)
	if err != nil {
		Error = err
		return
	}
	data, err := ctrl.access.RDBMS.GetIdTeamPage(req.TeamPageId)
	if err != nil {
		Error = err
		return
	}
	//qrc, err := qrcode.New(constant.Http+ "/api/" +"getByIdTeamPage/" + strconv.FormatUint(uint64(req.TeamPageId),10))
	qrc, err := qrcode.New(constant.Http+ "/"+ data.UUID.String())
	if err != nil {
		fmt.Printf("could not generate QRCode: %v", err)
	}

	path := string(constant.SaveFileLocationQrCode) + "/" + req.QrCode + strconv.FormatUint(uint64(req.TeamPageId),10) + ".PNG"
	// save file
	if err = qrc.Save(path); err != nil {
		fmt.Printf("could not save image: %v", err)
	}
	return err
}