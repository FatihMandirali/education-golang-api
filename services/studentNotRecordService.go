package services

import (
	. "education.api/config"
	"education.api/dbconnect"
	. "education.api/dto/request"
	. "education.api/entities"
	"education.api/enum"
	. "education.api/generic"
	"education.api/utils"
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

// student list
func GetNotRecordStudent(context *gin.Context) {
	queryPage, _ := strconv.Atoi(context.Query("page"))
	queryLimit, _ := strconv.Atoi(context.Query("limit"))
	branch, _ := strconv.Atoi(context.Query("branch"))
	search := context.Query("search")
	startDate, errorStartDate := time.Parse(time.RFC3339, context.Query("startDate"))
	endDate, errorEndDate := time.Parse(time.RFC3339, context.Query("endDate"))
	isRecord, _ := strconv.ParseBool(context.Query("isRecord"))
	isActive, _ := strconv.ParseBool(context.Query("isActive"))

	connection := dbconnect.DbInit()
	defer dbconnect.CloseDatabase(connection)

	var user []*User
	db := connection.Preload("Branch").Preload("Class").Where("role = ?", enum.Student).Where("is_record = ?", isRecord).Where("is_active = ?", isActive).Find(&user)

	if errorStartDate == nil || !startDate.IsZero() {
		db = db.Where("created_at >= ?", startDate)
	}
	if errorEndDate == nil || !endDate.IsZero() {
		db = db.Where("created_at <= ?", endDate)
	}
	if branch > 0 {
		db = db.Where("branch_id <= ?", branch)
	}
	if search != "" {
		db = db.Where("name LIKE ?", "%"+search+"%")
	}
	pagination := pagination.Paging(&pagination.Param{
		DB:      db,
		Page:    queryPage,
		Limit:   queryLimit,
		OrderBy: []string{"id desc"},
	}, &user)
	GenericResponse(context, SUCCESS, "", pagination)
}

// create student
func PostNotRecordStudent(context *gin.Context) {
	lang := context.Keys["Lang"]
	body := UserCreateRequest{}
	if err := context.ShouldBindJSON(&body); err != nil {
		utils.CheckError(err, context, utils.TextLanguage("badRequest", lang.(string)))
		return
	}
	connection := dbconnect.DbInit()
	defer dbconnect.CloseDatabase(connection)
	var user User
	connection.Where("email = ?", body.Email).First(&user)
	if user.Email != "" {
		GenericResponse(context, ERROR, utils.TextLanguage("emailAlreadyExist", lang.(string)), nil)
		return
	}

	hashPassword, err := utils.HashPassword(body.Password)
	if err != nil {
		utils.CheckError(err, context, utils.TextLanguage("error", lang.(string)))
		return
	}

	newUser := User{
		Name:        body.Name,
		Surname:     body.Surname,
		Email:       body.Email,
		Role:        enum.Student,
		Password:    hashPassword,
		PhoneNumber: body.PhoneNumber,
		ClassID:     body.ClassId,
		BranchID:    body.BranchId,
		IsRecord:    false,
		IsActive:    body.IsActive,
		CoverID:     body.CoverId,
	}
	connection.Create(&newUser)
	GenericResponse(context, SUCCESS, "", nil)
}

// update student
func UpdateNotRecordStudent(context *gin.Context) {
	lang := context.Keys["Lang"]
	body := UpdateUserRequest{}
	if err := context.ShouldBindJSON(&body); err != nil {
		utils.CheckError(err, context, utils.TextLanguage("badRequest", lang.(string)))
		return
	}
	connection := dbconnect.DbInit()
	defer dbconnect.CloseDatabase(connection)
	var user User

	connection.Where("id = ?", body.Id).Where("role = ?", enum.Student).Where("is_record = ?", false).First(&user)
	if user.Email == "" {
		GenericResponse(context, ERROR, utils.TextLanguage("notFound", lang.(string)), nil)
		return
	}

	var existUser User
	connection.Where("email = ?", body.Email).Not("id = ?", user.ID).First(&existUser)

	if existUser.Email != "" {
		GenericResponse(context, ERROR, utils.TextLanguage("emailAlreadyExist", lang.(string)), nil)
		return
	}
	if body.Password != "" {
		hashPassword, err := utils.HashPassword(body.Password)
		if err != nil {
			GenericResponse(context, ERROR, utils.TextLanguage("error", lang.(string)), nil)
			return
		}
		user.Password = hashPassword
	}

	user.Email = body.Email
	user.Name = body.Name
	user.Surname = body.Surname
	user.PhoneNumber = body.PhoneNumber
	user.ClassID = body.ClassId
	user.BranchID = body.BranchId
	user.CoverID = body.CoverId
	user.IsActive = body.IsActive
	connection.Save(&user)
	GenericResponse(context, SUCCESS, "", nil)
}

// getById student
func GetNotRecordStudentById(context *gin.Context) {
	lang := context.Keys["Lang"]
	uri := IdRequest{}
	if err := context.BindUri(&uri); err != nil {
		utils.CheckError(err, context, utils.TextLanguage("badRequest", lang.(string)))
		return
	}
	connection := dbconnect.DbInit()
	defer dbconnect.CloseDatabase(connection)
	var user User
	connection.Where("id = ?", uri.Id).Where("role = ?", enum.Student).First(&user)
	if user.Email == "" {
		GenericResponse(context, ERROR, utils.TextLanguage("notFound", lang.(string)), nil)
		return
	}

	GenericResponse(context, SUCCESS, "", user)
}

// delete student
func DeleteNotRecordStudentById(context *gin.Context) {
	lang := context.Keys["Lang"]
	uri := IdRequest{}
	if err := context.BindUri(&uri); err != nil {
		utils.CheckError(err, context, utils.TextLanguage("badRequest", lang.(string)))
		return
	}
	connection := dbconnect.DbInit()
	defer dbconnect.CloseDatabase(connection)
	connection.Where("role = ?", enum.Student).Where("is_record = ?", false).Delete(&User{}, uri.Id)
	GenericResponse(context, SUCCESS, "", nil)
}
