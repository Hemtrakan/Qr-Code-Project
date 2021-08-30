package control
//
//import (
//	"fmt"
//	"gorm.io/gorm"
//	rdbmsstructure "qrcode/access/rdbms/structure"
//)
//
//func (ctrl *APIControl) TestQrCode() (response rdbmsstructure.TestQrCode, Error error) {
//	Qrs := rdbmsstructure.TestQrCode{
//		QrCodeName: "computerData",
//		QrCodeId:   1,
//	}
//	err := ctrl.access.RDBMS.TestInsertQR(Qrs)
//	if err != nil {
//		Error = err
//		return
//	}
//
//	Ops := rdbmsstructure.TestOps{
//		OpsData:     "OpsData1",
//		QrCodeRefer: 1,
//	}
//	err = ctrl.access.RDBMS.TestInsertOps(Ops)
//	if err != nil {
//		Error = err
//		return
//	}
//	History := rdbmsstructure.TestHistory{
//		HistoryData: "HistoryData1",
//		QrCodeRefer: 1,
//	}
//	err = ctrl.access.RDBMS.TestInsertHistory(History)
//	if err != nil {
//		Error = err
//		return
//	}
//	res, err := ctrl.access.RDBMS.TestGetQrData()
//	if err != nil {
//		Error = err
//		return
//	}
//
//	Qr := rdbmsstructure.TestQrCode{}
//	for i, data := range res {
//		fmt.Println(i)
//		Qr = rdbmsstructure.TestQrCode{
//			Model: gorm.Model{
//				ID:        data.ID,
//				CreatedAt: data.CreatedAt,
//				UpdatedAt: data.UpdatedAt,
//				DeletedAt: data.DeletedAt,
//			},
//			QrCodeName:  data.QrCodeName,
//			QrCodeId:    data.QrCodeId,
//			DataHistory: data.DataHistory,
//			DataOps:     data.DataOps,
//		}
//	}
//	response = Qr
//
//	return
//}
