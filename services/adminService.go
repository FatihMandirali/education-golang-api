package services

import (
	. "education.api/config"
	"education.api/dbconnect"
	. "education.api/dto/request"
	"education.api/dto/response"
	. "education.api/entities"
	"education.api/enum"
	. "education.api/generic"
	"education.api/utils"
	"encoding/json"
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/swag/example/celler/httputil"
	"strconv"
)

// GetAdmins is the handler of list user endpoint
// @Summary List users
// @Description list all the users based on filter given
// @Tags user
// @Produce  json
// @
// @Router /api/admin [get]
func GetAdmins(context *gin.Context) {
	queryPage, _ := strconv.Atoi(context.Query("page"))
	queryLimit, _ := strconv.Atoi(context.Query("limit"))
	search := context.Query("search")

	connection := dbconnect.DbInit()
	defer dbconnect.CloseDatabase(connection)

	var user []*User
	db := connection.Where("role = ?", enum.Admin)

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
	}, &user)
	jsonParse, _ := json.Marshal(pagination.Records)
	userList := []User{}
	json.Unmarshal(jsonParse, &userList)
	pagination.Records = response.CreateAdminListResponse(userList)
	GenericResponse(context, SUCCESS, "", pagination)
}

// create admin
func PostAdmin(context *gin.Context) {
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

	newUser := User{Name: body.Name, Surname: body.Surname, Email: body.Email, Role: enum.Admin, Password: hashPassword, PhoneNumber: body.PhoneNumber}
	connection.Create(&newUser)
	GenericResponse(context, SUCCESS, "", nil)
}

// update admin
func UpdateAdmin(context *gin.Context) {
	lang := context.Keys["Lang"]
	body := UpdateUserRequest{}
	if err := context.ShouldBindJSON(&body); err != nil {
		utils.CheckError(err, context, utils.TextLanguage("badRequest", lang.(string)))
		return
	}
	connection := dbconnect.DbInit()
	defer dbconnect.CloseDatabase(connection)
	var user User
	connection.Where("id = ?", body.Id).Where("role = ?", enum.Admin).First(&user)
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
	connection.Save(&user)
	GenericResponse(context, SUCCESS, "", nil)
}

// getById admin
func GetAdminById(context *gin.Context) {
	lang := context.Keys["Lang"]
	uri := IdRequest{}
	if err := context.BindUri(&uri); err != nil {
		utils.CheckError(err, context, utils.TextLanguage("badRequest", lang.(string)))
		return
	}
	connection := dbconnect.DbInit()
	defer dbconnect.CloseDatabase(connection)
	var user User
	connection.Where("id = ?", uri.Id).Where("role = ?", enum.Admin).First(&user)
	if user.Email == "" {
		GenericResponse(context, ERROR, utils.TextLanguage("notFound", lang.(string)), nil)
		return
	}
	data := response.CreateAdminResponse(user)
	GenericResponse(context, SUCCESS, "", data)
}

// delete admin
func DeleteAdminById(context *gin.Context) {
	lang := context.Keys["Lang"]
	uri := IdRequest{}
	if err := context.BindUri(&uri); err != nil {
		utils.CheckError(err, context, utils.TextLanguage("badRequest", lang.(string)))
		return
	}
	connection := dbconnect.DbInit()
	defer dbconnect.CloseDatabase(connection)
	connection.Where("role = ?", enum.Admin).Delete(&User{}, uri.Id)
	GenericResponse(context, SUCCESS, "", nil)
}
