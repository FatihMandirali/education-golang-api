package services

import (
	"education.api/config"
	"education.api/dbconnect"
	"education.api/dto/request"
	. "education.api/entities"
	. "education.api/generic"
	"education.api/utils"
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

// branch list
func GetBranchList(context *gin.Context) {
	queryPage, _ := strconv.Atoi(context.Query("page"))
	queryLimit, _ := strconv.Atoi(context.Query("limit"))
	search := context.Query("search")
	startDate, errorStartDate := time.Parse(time.RFC3339, context.Query("startDate"))
	endDate, errorEndDate := time.Parse(time.RFC3339, context.Query("endDate"))
	//isActive, _ := strconv.ParseBool(context.Query("isActive"))

	connection := dbconnect.DbInit()
	defer dbconnect.CloseDatabase(connection)

	var branchList []*Branch
	db := connection

	if errorStartDate == nil || !startDate.IsZero() {
		db = db.Where("created_at >= ?", startDate)
	}

	if errorEndDate == nil || !endDate.IsZero() {
		db = db.Where("created_at <= ?", endDate)
	}
	if search != "" {
		db = db.Where("name LIKE ?", "%"+search+"%").Or("surname LIKE ? ", "%"+search+"%").Or("email LIKE ? ", "%"+search+"%").Or("phone_number LIKE ? ", "%"+search+"%")
	}
	//db := connection.Where("email = ?","fatih@gmail.com")
	//https://github.com/hellokaton/gorm-paginator
	pagination := pagination.Paging(&pagination.Param{
		DB:      db,
		Page:    queryPage,
		Limit:   queryLimit,
		OrderBy: []string{"id desc"},
	}, &branchList)

	GenericResponse(context, config.SUCCESS, "", pagination)
}

func GetAllBranchList(context *gin.Context) {
	connection := dbconnect.DbInit()
	defer dbconnect.CloseDatabase(connection)
	var branchList []*Branch
	connection.Where("is_active = ?", true).Find(&branchList)
	GenericResponse(context, config.SUCCESS, "", branchList)
}

// create branch
func PostBranch(context *gin.Context) {
	lang := context.Keys["Lang"]
	body := request.BranchCreateRequest{}
	if err := context.ShouldBindJSON(&body); err != nil {
		utils.CheckError(err, context, utils.TextLanguage("badRequest", lang.(string)))
		return
	}
	connection := dbconnect.DbInit()
	defer dbconnect.CloseDatabase(connection)
	var branch Branch
	connection.Where("name = ?", body.Name).First(&branch)
	if branch.Name != "" {
		GenericResponse(context, config.ERROR, utils.TextLanguage("branchAlreadyExist", lang.(string)), nil)
		return
	}

	newUser := Branch{Name: body.Name, PhoneNumber: body.PhoneNumber, Address: body.Address, IsActive: body.IsActive}
	connection.Create(&newUser)
	GenericResponse(context, config.SUCCESS, "", nil)
}

// update branch
func UpdateBranch(context *gin.Context) {
	lang := context.Keys["Lang"]
	body := request.BranchUpdateRequest{}
	if err := context.ShouldBindJSON(&body); err != nil {
		utils.CheckError(err, context, utils.TextLanguage("badRequest", lang.(string)))
		return
	}
	connection := dbconnect.DbInit()
	defer dbconnect.CloseDatabase(connection)
	var branch Branch
	connection.Where("id = ?", body.Id).First(&branch)
	if branch.Name == "" {
		GenericResponse(context, config.ERROR, utils.TextLanguage("notFound", lang.(string)), nil)
		return
	}

	var existBranch Branch
	connection.Where("name = ?", body.Name).Not("id = ?", branch.ID).First(&existBranch)

	if existBranch.Name != "" {
		GenericResponse(context, config.ERROR, utils.TextLanguage("branchAlreadyExist", lang.(string)), nil)
		return
	}

	branch.Name = body.Name
	branch.Address = body.Address
	branch.PhoneNumber = body.PhoneNumber
	branch.IsActive = body.IsActive
	connection.Save(&branch)
	GenericResponse(context, config.SUCCESS, "", nil)
}

// getById branch
func GetBranchById(context *gin.Context) {
	lang := context.Keys["Lang"]
	uri := request.IdRequest{}
	if err := context.BindUri(&uri); err != nil {
		utils.CheckError(err, context, utils.TextLanguage("badRequest", lang.(string)))
		return
	}
	connection := dbconnect.DbInit()
	defer dbconnect.CloseDatabase(connection)
	var branch Branch
	connection.Where("id = ?", uri.Id).First(&branch)
	if branch.Name == "" {
		GenericResponse(context, config.ERROR, utils.TextLanguage("notFound", lang.(string)), nil)
		return
	}

	GenericResponse(context, config.SUCCESS, "", branch)
}

// delete branch
func DeleteBranchById(context *gin.Context) {
	lang := context.Keys["Lang"]
	uri := request.IdRequest{}
	if err := context.BindUri(&uri); err != nil {
		utils.CheckError(err, context, utils.TextLanguage("badRequest", lang.(string)))
		return
	}
	connection := dbconnect.DbInit()
	defer dbconnect.CloseDatabase(connection)
	connection.Delete(&Branch{}, uri.Id)
	GenericResponse(context, config.SUCCESS, "", nil)
}
