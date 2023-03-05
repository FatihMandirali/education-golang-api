package middleware

import (
	"education.api/config"
	. "education.api/dto/response"
	"education.api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatusJSON(http.StatusOK, BaseResponse{StatusCode: config.ERROR, Message: utils.TextLanguage("error", "tr")})
			return
		}

		c.Next()
	}
}
