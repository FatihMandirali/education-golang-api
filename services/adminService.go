package services

import (
	. "education.api/config"
	"education.api/dbconnect"
	. "education.api/entities"
	. "education.api/generic"
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/gin-gonic/gin"
	"strconv"
)

// admin list
func GetAdmins(context *gin.Context) {
	query := context.Request.URL.Query()
	queryPage, _ := strconv.Atoi(query.Get("page"))
	queryLimit, _ := strconv.Atoi(query.Get("limit"))

	connection := dbconnect.DbInit()
	defer dbconnect.CloseDatabase(connection)

	var user []*User
	db := connection
	//db := connection.Where("email = ?","fatih@gmail.com")
	//https://github.com/hellokaton/gorm-paginator
	response := pagination.Paging(&pagination.Param{
		DB:      db,
		Page:    queryPage,
		Limit:   queryLimit,
		OrderBy: []string{"id desc"},
	}, &user)
	GenericResponse(context, SUCCESS, "", response)
}
