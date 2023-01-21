package services

import (
	. "education.api/config"
	"education.api/dbconnect"
	. "education.api/dto/response"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// admin list
func GetAdmins(context *gin.Context) {
	lang := context.Keys["Lang"]
	log.Printf(lang.(string))
	connection := dbconnect.DbInit()
	defer dbconnect.CloseDatabase(connection)

	context.AbortWithStatusJSON(http.StatusOK, BaseResponse{StatusCode: SUCCESS})
}
