package services

import (
	. "education.api/config"
	"education.api/dbconnect"
	. "education.api/dto/request"
	. "education.api/entities"
	"education.api/enum"
	. "education.api/generic"
	"education.api/utils"
	"github.com/gin-gonic/gin"
)

// teacher list
func GetTeacher(context *gin.Context) {
	connection := dbconnect.DbInit()
	defer dbconnect.CloseDatabase(connection)

	var user []*User
	connection.Where("role = ?", enum.Teacher).Find(&user)

	GenericResponse(context, SUCCESS, "", user)
}

// create teacher
func PostTeacher(context *gin.Context) {
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
		Role:        enum.Teacher,
		Password:    hashPassword,
		PhoneNumber: body.PhoneNumber,
		LessonID:    body.LessonId,
		ClassID:     body.ClassId,
		BranchID:    body.BranchId,
	}
	connection.Create(&newUser)
	GenericResponse(context, SUCCESS, "", nil)
}

// update teacher
func UpdateTeacher(context *gin.Context) {
	lang := context.Keys["Lang"]
	body := UpdateUserRequest{}
	if err := context.ShouldBindJSON(&body); err != nil {
		utils.CheckError(err, context, utils.TextLanguage("badRequest", lang.(string)))
		return
	}
	connection := dbconnect.DbInit()
	defer dbconnect.CloseDatabase(connection)
	var user User
	connection.Where("id = ?", body.Id).Where("role = ?", enum.Teacher).First(&user)
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

	hashPassword, err := utils.HashPassword(body.Password)
	if err != nil {
		GenericResponse(context, ERROR, utils.TextLanguage("error", lang.(string)), nil)
		return
	}
	user.Password = hashPassword
	user.Email = body.Email
	user.Name = body.Name
	user.Surname = body.Surname
	user.PhoneNumber = body.PhoneNumber
	user.ClassID = body.ClassId
	user.BranchID = body.BranchId
	user.LessonID = body.LessonId
	connection.Save(&user)
	GenericResponse(context, SUCCESS, "", nil)
}

// getById teacher
func GetTeacherById(context *gin.Context) {
	lang := context.Keys["Lang"]
	uri := IdRequest{}
	if err := context.BindUri(&uri); err != nil {
		utils.CheckError(err, context, utils.TextLanguage("badRequest", lang.(string)))
		return
	}
	connection := dbconnect.DbInit()
	defer dbconnect.CloseDatabase(connection)
	var user User
	connection.Where("id = ?", uri.Id).Where("role = ?", enum.Teacher).First(&user)
	if user.Email == "" {
		GenericResponse(context, ERROR, utils.TextLanguage("notFound", lang.(string)), nil)
		return
	}

	GenericResponse(context, SUCCESS, "", user)
}

// delete teacher
func DeleteTeacherById(context *gin.Context) {
	lang := context.Keys["Lang"]
	uri := IdRequest{}
	if err := context.BindUri(&uri); err != nil {
		utils.CheckError(err, context, utils.TextLanguage("badRequest", lang.(string)))
		return
	}
	connection := dbconnect.DbInit()
	defer dbconnect.CloseDatabase(connection)
	connection.Where("role = ?", enum.Teacher).Delete(&User{}, uri.Id)
	GenericResponse(context, SUCCESS, "", nil)
}
