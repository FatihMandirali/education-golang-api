package utils

import (
	. "education.api/config"
	. "education.api/generic"
	"github.com/gin-gonic/gin"
	"log"
)

func CheckError(err error, context *gin.Context, message string) {
	if err != nil {
		log.Fatalln(err.Error())
		GenericResponse(context, ERROR, message, nil)
		return
	}
}
