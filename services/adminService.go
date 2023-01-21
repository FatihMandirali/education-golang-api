package services

import (
	. "education.api/config"
	"education.api/dbconnect"
	. "education.api/generic"
	"github.com/gin-gonic/gin"
	"log"
)

// admin list
func GetAdmins(context *gin.Context) {
	lang := context.Keys["Lang"]
	log.Printf(lang.(string))
	connection := dbconnect.DbInit()
	defer dbconnect.CloseDatabase(connection)

	GenericResponse(context, SUCCESS, "", nil)
}
