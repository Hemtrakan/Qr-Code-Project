package control

import (
	"encoding/json"
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"io/ioutil"
	"os"
	"qrcode/access/constant"
	rdbmsstructure "qrcode/access/rdbms/structure"
	"qrcode/present/structure"
	"strconv"
	"time"
)

func (ctrl *APIControl) GetHtml(id string) (URL string ,Error error) {
	client := resty.New()
	_, err := client.R().
		//SetPathParam("id", id).
		Get("http://localhost:12000/viewdata/"+id)
	if err != nil {
		Error = err
		return
	}
	return
}

func (ctrl *APIControl) GetByIdTeamPage(teamPageId structure.GetByIdTeamPage) (response *structure.ResGetByIdTeamPage, Error error) {
	data, err := ctrl.access.RDBMS.GetByIdTeamPage(teamPageId.TeamPageId)
	jsonFile, err := os.Open(string(constant.SaveFileLocation) + "/" + data.TeamPageFile)
	if err != nil {
		Error = err
		return
	}
	defer jsonFile.Close() // ปิดไฟล์หลังจากทำงานเสร็จทั้งหมด
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var resFile interface{}
	err = json.Unmarshal(byteValue, &resFile)
	if err != nil {
		Error = errors.New("record not found")
		return
	}
	TeamPageStructure := structure.ResGetByIdTeamPage{
		TeamPageId: data.ID,
		Data:       resFile,
	}
	response = &TeamPageStructure
	return
}

func (ctrl *APIControl) GetAllTeamPage(ownersId int) (response []*structure.GetAllTeamPage, Error error) {
	var TeamPageArray []*structure.GetAllTeamPage
	res, err := ctrl.access.RDBMS.GetAllTeamPage(ownersId)
	if err != nil {
		Error = err
		return
	}
	for _, TeamPage := range res {
		var TeamPageStructure = structure.GetAllTeamPage{
			Id:           TeamPage.ID,
			TeamPageId:   TeamPage.UUID,
			TeamPageName: TeamPage.TeamPageName,
			TeamPageFile: TeamPage.TeamPageFile,
		}
		TeamPageArray = append(TeamPageArray, &TeamPageStructure)
	}
	response = TeamPageArray
	return
}

func (ctrl *APIControl) InsertTeamPage(reqTeamPage *structure.TeamPage) (Error error) {
	file, _ := json.MarshalIndent(reqTeamPage, "", " ")
	loc, _ := time.LoadLocation("Asia/Bangkok")
	fileNamePrefix := time.Now().In(loc).Format("20060102_150405")
	TeamPageFile := fileNamePrefix + ".json"
	uuid, err := uuid.NewV4()
	if err != nil {
		Error = err
		return
	}
	TeamPage := rdbmsstructure.TeamPage{
		TeamPageName: reqTeamPage.TeamPageName,
		TeamPageFile: TeamPageFile,
		OwnersId:     uint(reqTeamPage.OwnerId),
		UUID:         uuid,
	}
	err = ctrl.insertTeamPage(TeamPage)
	if err != nil {
		Error = err
		return
	}

	err = ioutil.WriteFile(string(constant.SaveFileLocation)+"/"+TeamPageFile, file, 0644) //todo ต้องทำเป็น Env
	if err != nil {
		Error = err
		return
	}
	return
}

func (ctrl *APIControl) UpdateTeamPage(teamPageId structure.GetByIdTeamPage, reqTeamPage *structure.TeamPage) (Error error) {
	data, err := ctrl.access.RDBMS.GetByIdTeamPage(teamPageId.TeamPageId)
	file, _ := json.MarshalIndent(reqTeamPage, "", " ")
	dataLog, err := ctrl.access.RDBMS.GetByIdLogTeamPage(teamPageId.Id)

	// logfile
	TeamPageFile := ""
	i := len(dataLog) + 1
	TeamPageFile = strconv.Itoa(i) + "-" + data.TeamPageFile
	LogTeamPage := rdbmsstructure.LogTeamPage{
		LogTeamPageName: reqTeamPage.TeamPageName,
		LogTeamPageFile: TeamPageFile,
		TeamPageId:      data.ID,
		OwnersId:        data.OwnersId,
	}
	err = os.Rename(string(constant.SaveFileLocation)+"/"+data.TeamPageFile, string(constant.LogLocation)+"/"+TeamPageFile)
	if err != nil {
		Error = err
		return
	}
	err = ctrl.access.RDBMS.InsertLogTeamPage(LogTeamPage)
	if err != nil {
		Error = err
		return
	}

	// เขียนไฟล์
	err = ioutil.WriteFile(string(constant.SaveFileLocation)+"/"+data.TeamPageFile, file, 0644) //todo ต้องทำเป็น Env
	if err != nil {
		Error = err
		return
	}
	TeamPage := rdbmsstructure.TeamPage{
		Model: gorm.Model{
			ID: teamPageId.Id,
		},
		TeamPageName: reqTeamPage.TeamPageName,
	}
	err = ctrl.updateTeamPage(TeamPage)
	if err != nil {
		Error = err
		return
	}
	return
}

func (ctrl *APIControl) DeleteTeamPage(teamPageId structure.GetByIdTeamPage) (Error error) {
	//Data, err := ctrl.access.RDBMS.GetByIdTeamPage(teamPageId.ID)
	//if err != nil {
	//	Error = err
	//	return
	//}
	//err = os.Remove("fileteampage/" + Data.TeamPageFile)
	//if err != nil {
	//	log.Fatal(err)
	//}
	err := ctrl.access.RDBMS.DeleteTeamPage(teamPageId.Id)
	if err != nil {
		Error = err
		return
	}
	return
}

func (ctrl *APIControl) insertTeamPage(TeamPage rdbmsstructure.TeamPage) error {
	err := ctrl.access.RDBMS.InsertTeamPage(TeamPage)
	if err != nil {
		return err
	}
	return nil
}
func (ctrl *APIControl) updateTeamPage(TeamPage rdbmsstructure.TeamPage) error {
	err := ctrl.access.RDBMS.UpdateTeamPage(TeamPage)
	if err != nil {
		return err
	}
	return nil
}

func (ctrl *APIControl) GetAllLogTeamPage(teamPageId structure.GetAllLogTeamPage) (response []structure.GetAllLogTeamPage, Error error) {
	var TeamPageArray []structure.GetAllLogTeamPage
	res, err := ctrl.access.RDBMS.GetAllDataListLogTeamPage(teamPageId.ID)
	if err != nil {
		Error = err
		return
	}

	for _, data := range res {
		jsonFile, err := os.Open(string(constant.LogLocation) + "/" + data.LogTeamPageFile)
		if err != nil {
			Error = err
			return
		}
		byteValue, _ := ioutil.ReadAll(jsonFile)
		var resLog interface{}
		err = json.Unmarshal(byteValue, &resLog)
		if err != nil {
			Error = err
			return
		}
		defer jsonFile.Close() // ปิดไฟล์หลังจากทำงานเสร็จทั้งหมด
		TeamPageStructure := structure.GetAllLogTeamPage{
			ID:        data.ID,
			UpdatedAt: data.UpdatedAt,
			LogData:   &resLog,
		}
		TeamPageArray = append(TeamPageArray, TeamPageStructure)
	}
	response = TeamPageArray
	return
}
