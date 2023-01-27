package services

import (
	"education.api/config"
	"education.api/dbconnect"
	"education.api/dto/request"
	"education.api/entities"
	. "education.api/generic"
	"education.api/utils"
	"github.com/gin-gonic/gin"
)

// announcement list
func GetAnnouncementList(context *gin.Context) {
	connection := dbconnect.DbInit()
	defer dbconnect.CloseDatabase(connection)

	var branchList []*entities.Announcement
	connection.Preload("User").Find(&branchList)

	GenericResponse(context, config.SUCCESS, "", branchList)
}

// create announcement
func PostAnnouncement(context *gin.Context) {
	lang := context.Keys["Lang"]
	body := request.AnnouncementCreateRequest{}
	if err := context.ShouldBindJSON(&body); err != nil {
		utils.CheckError(err, context, utils.TextLanguage("badRequest", lang.(string)))
		return
	}
	connection := dbconnect.DbInit()
	defer dbconnect.CloseDatabase(connection)

	newAnnouncement := entities.Announcement{Title: body.Title, Description: body.Description, Type: body.Type, UserID: body.UserId, StartDate: body.StartDate, EndDate: body.EndDate}
	connection.Create(&newAnnouncement)
	GenericResponse(context, config.SUCCESS, "", nil)
}

// update announcement
func UpdateAnnouncement(context *gin.Context) {
	lang := context.Keys["Lang"]
	body := request.AnnouncementUpdateRequest{}
	if err := context.ShouldBindJSON(&body); err != nil {
		utils.CheckError(err, context, utils.TextLanguage("badRequest", lang.(string)))
		return
	}
	connection := dbconnect.DbInit()
	defer dbconnect.CloseDatabase(connection)
	var announcement entities.Announcement
	connection.Where("id = ?", body.Id).First(&announcement)
	if announcement.Title == "" {
		GenericResponse(context, config.ERROR, utils.TextLanguage("notFound", lang.(string)), nil)
		return
	}

	announcement.Title = body.Title
	announcement.Description = body.Description
	announcement.Type = body.Type
	announcement.UserID = body.UserId
	announcement.StartDate = body.StartDate
	announcement.EndDate = body.EndDate
	connection.Save(&announcement)
	GenericResponse(context, config.SUCCESS, "", nil)
}

// getById announcement
func GetAnnouncementById(context *gin.Context) {
	lang := context.Keys["Lang"]
	uri := request.IdRequest{}
	if err := context.BindUri(&uri); err != nil {
		utils.CheckError(err, context, utils.TextLanguage("badRequest", lang.(string)))
		return
	}
	connection := dbconnect.DbInit()
	defer dbconnect.CloseDatabase(connection)
	var announcement entities.Announcement
	connection.Where("id = ?", uri.Id).First(&announcement)
	if announcement.Title == "" {
		GenericResponse(context, config.ERROR, utils.TextLanguage("notFound", lang.(string)), nil)
		return
	}

	GenericResponse(context, config.SUCCESS, "", announcement)
}

// delete announcement
func DeleteAnnouncementById(context *gin.Context) {
	lang := context.Keys["Lang"]
	uri := request.IdRequest{}
	if err := context.BindUri(&uri); err != nil {
		utils.CheckError(err, context, utils.TextLanguage("badRequest", lang.(string)))
		return
	}
	connection := dbconnect.DbInit()
	defer dbconnect.CloseDatabase(connection)
	connection.Delete(&entities.Announcement{}, uri.Id)
	GenericResponse(context, config.SUCCESS, "", nil)
}
