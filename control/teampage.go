package control

import (
	"qrcode/access/constant"
	templates2 "qrcode/present/structure/templates"
)

func (ctrl *APIControl) GetTemplate() templates2.Templatesdata {
	templates := constant.Template
	var arrTemplates templates2.Templatesdata
	for _, item := range templates {
		codeWithFullName := templates2.Template{
			TemplatesName: item,
		}
		arrTemplates.Data = append(arrTemplates.Data, codeWithFullName)
	}
	return arrTemplates
}

//
//func (ctrl *APIControl) GetHtml(id string) (URL string, Error error) {
//	client := resty.New()
//	_, err := client.R().
//		//SetPathParam("id", id).
//		Get("http://localhost:12000/viewdata/" + id)
//	if err != nil {
//		Error = err
//		return
//	}
//	return
//}
//
//func (ctrl *APIControl) GetByIdTeamPage(teamPageId structure.GetByIdTeamPage) (response *structure.ResGetByIdTeamPage, Error error) {
//	data, err := ctrl.access.RDBMS.GetByIdTeamPage(teamPageId.TeamPageId)
//	jsonFile, err := os.Open(string(constant.SaveFileLocation) + "/" + data.TeamPageFile)
//	if err != nil {
//		Error = err
//		return
//	}
//	defer jsonFile.Close() // ปิดไฟล์หลังจากทำงานเสร็จทั้งหมด
//	byteValue, _ := ioutil.ReadAll(jsonFile)
//	var resFile interface{}
//	err = json.Unmarshal(byteValue, &resFile)
//	if err != nil {
//		Error = errors.New("record not found")
//		return
//	}
//	TeamPageStructure := structure.ResGetByIdTeamPage{
//		TeamPageId: data.ID,
//		Data:       resFile,
//	}
//	response = &TeamPageStructure
//	return
//}

//func (ctrl *APIControl) GetAllTeamPage(id int) (response []*structure.GetAllTeamPage, Error error) {
//	//var TeamPageArray []*structure.GetAllTeamPage
//	//res, err := ctrl.access.RDBMS.GetAllTeamPage(id)
//	//if err != nil {
//	//	Error = err
//	//	return
//	//}
//	//for _, TeamPage := range res {
//	//	var TeamPageStructure = structure.GetAllTeamPage{
//	//		Id:           TeamPage.ID,
//	//		TeamPageId:   TeamPage.UUID,
//	//		TeamPageName: TeamPage.TeamPageName,
//	//		TeamPageFile: TeamPage.TeamPageFile,
//	//		QrCodeType:   TeamPage.QrCodeType,
//	//	}
//	//	TeamPageArray = append(TeamPageArray, &TeamPageStructure)
//	//}
//	//response = TeamPageArray
//	return
//}

//func (ctrl *APIControl) InsertTeamPage(reqTeamPage *[]structure.Template) (Error error) {
//	//for index, req := range *reqTeamPage {
//	//	uuid, err := uuid.NewV4()
//	//	if err != nil {
//	//		Error = err
//	//		return
//	//	}
//	//	loc, _ := time.LoadLocation("Asia/Bangkok")
//	//	fileNamePrefix := time.Now().In(loc).Format("20060102_150405")
//	//	TeamPageFile := strconv.Itoa(index+1) + "-" + fileNamePrefix + ".json"
//	//	TeamPage := rdbmsstructure.TeamPage{
//	//		TeamPageName: req.TeamPageName,
//	//		TeamPageFile: TeamPageFile,
//	//		UUID:         uuid,
//	//		OwnersId:     uint(req.OwnerId),
//	//		QrCodeType:   req.QrCodeTypeId,
//	//	}
//	//	data, err := ctrl.access.RDBMS.InsertTeamPage(TeamPage)
//	//	//err = ctrl.insertTeamPage(TeamPage)
//	//	if err != nil {
//	//		Error = err
//	//		return
//	//	}
//	//	// บันทึก TeamPage
//	//	FileJson := structure.TeamPage{
//	//		OwnerId:      req.OwnerId,
//	//		TeamPageName: req.TeamPageName,
//	//		QrCodeTypeId:   req.QrCodeTypeId,
//	//		Info:         req.Info,
//	//		Ops:          req.Ops,
//	//	}
//	//
//	//	file, _ := json.MarshalIndent(FileJson, "", " ")
//	//	err = ioutil.WriteFile(string(constant.SaveFileLocation)+"/"+TeamPageFile, file, 0644) //todo ต้องทำเป็น Env
//	//	if err != nil {
//	//		Error = err
//	//		return
//	//	}
//	//
//	//	// สร้าง QR-Code
//	//	qrc, err := qrcode.New(constant.Http + "/" + uuid.String())
//	//	if err != nil {
//	//		fmt.Printf("could not generate QRCode: %v", err)
//	//	}
//	//
//	//	path := string(constant.SaveFileLocationQrCode) + "/" + req.TeamPageName + "_" + strconv.FormatUint(uint64(data.ID), 10) + ".PNG"
//	//	// save file
//	//	if err = qrc.Save(path); err != nil {
//	//		fmt.Printf("could not save image: %v", err)
//	//	}
//	//}
//	return
//}

//func (ctrl *APIControl) UpdateTeamPage(teamPageId structure.GetByIdTeamPage, reqTeamPage *structure.Template) (Error error) {
//	//data, err := ctrl.access.RDBMS.GetByIdTeamPage(teamPageId.TeamPageId)
//	//file, _ := json.MarshalIndent(reqTeamPage, "", " ")
//	//dataLog, err := ctrl.access.RDBMS.GetByIdLogTeamPage(teamPageId.Id)
//	//
//	//// logfile
//	//TeamPageFile := ""
//	//i := len(dataLog) + 1
//	//TeamPageFile = strconv.Itoa(i) + "-" + data.TeamPageFile
//	//LogTeamPage := rdbmsstructure.LogTeamPage{
//	//	LogTeamPageName: reqTeamPage.TeamPageName,
//	//	LogTeamPageFile: TeamPageFile,
//	//	TeamPageId:      data.ID,
//	//	OwnersId:        data.OwnersId,
//	//}
//	//err = os.Rename(string(constant.SaveFileLocation)+"/"+data.TeamPageFile, string(constant.LogLocation)+"/"+TeamPageFile)
//	//if err != nil {
//	//	Error = err
//	//	return
//	//}
//	//err = ctrl.access.RDBMS.InsertLogTeamPage(LogTeamPage)
//	//if err != nil {
//	//	Error = err
//	//	return
//	//}
//	//
//	//// เขียนไฟล์
//	//err = ioutil.WriteFile(string(constant.SaveFileLocation)+"/"+data.TeamPageFile, file, 0644) //todo ต้องทำเป็น Env
//	//if err != nil {
//	//	Error = err
//	//	return
//	//}
//	//TeamPage := rdbmsstructure.TeamPage{
//	//	Model: gorm.Model{
//	//		ID: teamPageId.Id,
//	//	},
//	//	TeamPageName: reqTeamPage.TeamPageName,
//	//}
//	//err = ctrl.updateTeamPage(TeamPage)
//	//if err != nil {
//	//	Error = err
//	//	return
//	//}
//	return
//}

//func (ctrl *APIControl) DeleteTeamPage(teamPageId structure.GetByIdTeamPage) (Error error) {
//	//Data, err := ctrl.access.RDBMS.GetByIdTeamPage(teamPageId.ID)
//	//if err != nil {
//	//	Error = err
//	//	return
//	//}
//	//err = os.Remove("fileteampage/" + Data.TeamPageFile)
//	//if err != nil {
//	//	log.Fatal(err)
//	//}
//	err := ctrl.access.RDBMS.DeleteTeamPage(teamPageId.Id)
//	if err != nil {
//		Error = err
//		return
//	}
//	return
//}

//func (ctrl *APIControl) insertTeamPage(TeamPage rdbmsstructure.TeamPage) error {
//	err := ctrl.access.RDBMS.InsertTeamPage(TeamPage)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func (ctrl *APIControl) updateTeamPage(TeamPage rdbmsstructure.Template) error {
//	err := ctrl.access.RDBMS.UpdateTeamPage(TeamPage)
//	if err != nil {
//		return err
//	}
//	return nil
//}

//func (ctrl *APIControl) GetAllLogTeamPage(teamPageId structure.GetAllLogTeamPage) (response []structure.GetAllLogTeamPage, Error error) {
//	var TeamPageArray []structure.GetAllLogTeamPage
//	res, err := ctrl.access.RDBMS.GetAllDataListLogTeamPage(teamPageId.ID)
//	if err != nil {
//		Error = err
//		return
//	}
//
//	for _, data := range res {
//		jsonFile, err := os.Open(string(constant.LogLocation) + "/" + data.LogTeamPageFile)
//		if err != nil {
//			Error = err
//			return
//		}
//		byteValue, _ := ioutil.ReadAll(jsonFile)
//		var resLog interface{}
//		err = json.Unmarshal(byteValue, &resLog)
//		if err != nil {
//			Error = err
//			return
//		}
//		defer jsonFile.Close() // ปิดไฟล์หลังจากทำงานเสร็จทั้งหมด
//		TeamPageStructure := structure.GetAllLogTeamPage{
//			ID:        data.ID,
//			UpdatedAt: data.UpdatedAt,
//			LogData:   &resLog,
//		}
//		TeamPageArray = append(TeamPageArray, TeamPageStructure)
//	}
//	response = TeamPageArray
//	return
//}
