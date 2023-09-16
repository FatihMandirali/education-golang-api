package services

import (
	"education.api/config"
	"education.api/dbconnect"
	"education.api/dto/request"
	"education.api/entities"
	. "education.api/generic"
	"education.api/utils"
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

// class list
func GetClassList(context *gin.Context) {
	queryPage, _ := strconv.Atoi(context.Query("page"))
	queryLimit, _ := strconv.Atoi(context.Query("limit"))
	search := context.Query("search")
	startDate, errorStartDate := time.Parse(time.RFC3339, context.Query("startDate"))
	endDate, errorEndDate := time.Parse(time.RFC3339, context.Query("endDate"))
	//isActive, _ := strconv.ParseBool(context.Query("isActive"))

	connection := dbconnect.DbInit()
	defer dbconnect.CloseDatabase(connection)

	var classList []*entities.Class
	db := connection.Preload("Branch")
	//connection.Preload("Branch").Find(&classList)
	if errorStartDate == nil || !startDate.IsZero() {
		db = db.Where("created_at >= ?", startDate)
	}
	if errorEndDate == nil || !endDate.IsZero() {
		db = db.Where("created_at <= ?", endDate)
	}
	if search != "" {
		db = db.Where("name LIKE ?", "%"+search+"%")
	}
	pagination := pagination.Paging(&pagination.Param{
		DB:      db,
		Page:    queryPage,
		Limit:   queryLimit,
		OrderBy: []string{"id desc"},
	}, &classList)
	GenericResponse(context, config.SUCCESS, "", pagination)
}

// class list
func GetClassAllList(context *gin.Context) {
	connection := dbconnect.DbInit()
	defer dbconnect.CloseDatabase(connection)

	var classList []*entities.Class
	connection.Preload("Branch").Where("is_active = ?", true).Find(&classList)
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

	newClass := entities.Class{Name: body.Name, BranchID: body.BranchId, IsActive: body.IsActive}
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
	class.IsActive = body.IsActive
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
