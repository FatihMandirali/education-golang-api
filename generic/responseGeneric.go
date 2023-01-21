package generic

import (
	"education.api/dto/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GenericResponse(context *gin.Context, status string, message string, data interface{}) {
	context.AbortWithStatusJSON(http.StatusOK, response.BaseResponse{
		Data:       data,
		StatusCode: status,
		Message:    message,
	})
}
