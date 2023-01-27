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

// class list
func GetClassList(context *gin.Context) {
	connection := dbconnect.DbInit()
	defer dbconnect.CloseDatabase(connection)

	var classList []*entities.Class
	connection.Preload("Branch").Find(&classList)

	GenericResponse(context, config.SUCCESS, "", classList)
}

// create class
func PostClass(context *gin.Context) {
	lang := context.Keys["Lang"]
	body := request.ClassCreateRequest{}
	if err := context.ShouldBindJSON(&body); err != nil {
		utils.CheckError(err, context, utils.TextLanguage("badRequest", lang.(string)))
		return
	}
	connection := dbconnect.DbInit()
	defer dbconnect.CloseDatabase(connection)
	var class entities.Class
	connection.Where("name = ?", body.Name).First(&class)
	if class.Name != "" {
		GenericResponse(context, config.ERROR, utils.TextLanguage("classAlreadyExist", lang.(string)), nil)
		return
	}

	newClass := entities.Class{Name: body.Name, BranchID: body.BranchId}
	connection.Create(&newClass)
	GenericResponse(context, config.SUCCESS, "", nil)
}

// update class
func UpdateClass(context *gin.Context) {
	lang := context.Keys["Lang"]
	body := request.ClassUpdateRequest{}
	if err := context.ShouldBindJSON(&body); err != nil {
		utils.CheckError(err, context, utils.TextLanguage("badRequest", lang.(string)))
		return
	}
	connection := dbconnect.DbInit()
	defer dbconnect.CloseDatabase(connection)
	var class entities.Class
	connection.Where("id = ?", body.Id).First(&class)
	if class.Name == "" {
		GenericResponse(context, config.ERROR, utils.TextLanguage("notFound", lang.(string)), nil)
		return
	}

	var existClass entities.Class
	connection.Where("name = ?", body.Name).Not("id = ?", class.ID).First(&existClass)

	if existClass.Name != "" {
		GenericResponse(context, config.ERROR, utils.TextLanguage("classAlreadyExist", lang.(string)), nil)
		return
	}

	class.Name = body.Name
	class.BranchID = body.BranchId
	connection.Save(&class)
	GenericResponse(context, config.SUCCESS, "", nil)
}

// getById class
func GetClassById(context *gin.Context) {
	lang := context.Keys["Lang"]
	uri := request.IdRequest{}
	if err := context.BindUri(&uri); err != nil {
		utils.CheckError(err, context, utils.TextLanguage("badRequest", lang.(string)))
		return
	}
	connection := dbconnect.DbInit()
	defer dbconnect.CloseDatabase(connection)
	var class entities.Class
	connection.Preload("Branch").Where("id = ?", uri.Id).First(&class)
	if class.Name == "" {
		GenericResponse(context, config.ERROR, utils.TextLanguage("notFound", lang.(string)), nil)
		return
	}

	GenericResponse(context, config.SUCCESS, "", class)
}

// delete class
func DeleteClassById(context *gin.Context) {
	lang := context.Keys["Lang"]
	uri := request.IdRequest{}
	if err := context.BindUri(&uri); err != nil {
		utils.CheckError(err, context, utils.TextLanguage("badRequest", lang.(string)))
		return
	}
	connection := dbconnect.DbInit()
	defer dbconnect.CloseDatabase(connection)
	connection.Delete(&entities.Class{}, uri.Id)
	GenericResponse(context, config.SUCCESS, "", nil)
}
