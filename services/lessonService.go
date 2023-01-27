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

// lesson list
func GetLessonList(context *gin.Context) {
	connection := dbconnect.DbInit()
	defer dbconnect.CloseDatabase(connection)

	var lessonList []*entities.Lesson
	connection.Find(&lessonList)

	GenericResponse(context, config.SUCCESS, "", lessonList)
}

// create lesson
func PostLesson(context *gin.Context) {
	lang := context.Keys["Lang"]
	body := request.LessonCreateRequest{}
	if err := context.ShouldBindJSON(&body); err != nil {
		utils.CheckError(err, context, utils.TextLanguage("badRequest", lang.(string)))
		return
	}
	connection := dbconnect.DbInit()
	defer dbconnect.CloseDatabase(connection)
	var lesson entities.Lesson
	connection.Where("name = ?", body.Name).First(&lesson)
	if lesson.Name != "" {
		GenericResponse(context, config.ERROR, utils.TextLanguage("lessonAlreadyExist", lang.(string)), nil)
		return
	}

	newLesson := entities.Lesson{Name: body.Name}
	connection.Create(&newLesson)
	GenericResponse(context, config.SUCCESS, "", nil)
}

// update lesson
func UpdateLesson(context *gin.Context) {
	lang := context.Keys["Lang"]
	body := request.LessonUpdateRequest{}
	if err := context.ShouldBindJSON(&body); err != nil {
		utils.CheckError(err, context, utils.TextLanguage("badRequest", lang.(string)))
		return
	}
	connection := dbconnect.DbInit()
	defer dbconnect.CloseDatabase(connection)
	var lesson entities.Lesson
	connection.Where("id = ?", body.Id).First(&lesson)
	if lesson.Name == "" {
		GenericResponse(context, config.ERROR, utils.TextLanguage("notFound", lang.(string)), nil)
		return
	}

	var existLesson entities.Lesson
	connection.Where("name = ?", body.Name).Not("id = ?", lesson.ID).First(&existLesson)

	if existLesson.Name != "" {
		GenericResponse(context, config.ERROR, utils.TextLanguage("lessonAlreadyExist", lang.(string)), nil)
		return
	}

	lesson.Name = body.Name
	connection.Save(&lesson)
	GenericResponse(context, config.SUCCESS, "", nil)
}

// getById lesson
func GetLessonById(context *gin.Context) {
	lang := context.Keys["Lang"]
	uri := request.IdRequest{}
	if err := context.BindUri(&uri); err != nil {
		utils.CheckError(err, context, utils.TextLanguage("badRequest", lang.(string)))
		return
	}
	connection := dbconnect.DbInit()
	defer dbconnect.CloseDatabase(connection)
	var lesson entities.Lesson
	connection.Where("id = ?", uri.Id).First(&lesson)
	if lesson.Name == "" {
		GenericResponse(context, config.ERROR, utils.TextLanguage("notFound", lang.(string)), nil)
		return
	}

	GenericResponse(context, config.SUCCESS, "", lesson)
}

// delete lesson
func DeleteLessonById(context *gin.Context) {
	lang := context.Keys["Lang"]
	uri := request.IdRequest{}
	if err := context.BindUri(&uri); err != nil {
		utils.CheckError(err, context, utils.TextLanguage("badRequest", lang.(string)))
		return
	}
	connection := dbconnect.DbInit()
	defer dbconnect.CloseDatabase(connection)
	connection.Delete(&entities.Lesson{}, uri.Id)
	GenericResponse(context, config.SUCCESS, "", nil)
}
