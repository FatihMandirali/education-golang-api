package middleware

import (
	. "education.api/config"
	. "education.api/dto/response"
	. "education.api/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
)

func ValidateToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		lang := context.Request.Header.Get("Accept-Language")
		token := context.Request.Header.Get("Authorization")
		claims := &TokenResponse{}
		tkn, error := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(JwtSecret), nil
		})

		if tkn == nil || !tkn.Valid || error != nil {
			context.AbortWithStatusJSON(http.StatusOK, BaseResponse{StatusCode: ERROR, Message: TextLanguage("tokenError", lang)})
			return
		}
		context.Set("Role", string(claims.Role))
		context.Set("Lang", lang)
	}
}

func AuthorizationToken(validRoles []string) gin.HandlerFunc {
	return func(context *gin.Context) {
		if len(context.Keys) == 0 {
			context.AbortWithStatusJSON(http.StatusOK, BaseResponse{StatusCode: ERROR, Message: TextLanguage("tokenError", "")})
		}
		roleVal := context.Keys["Role"]
		langVal := context.Keys["Lang"]
		if roleVal == nil || langVal == nil {
			context.AbortWithStatusJSON(http.StatusOK, BaseResponse{StatusCode: ERROR, Message: TextLanguage("tokenError", "")})
		}
		lang := langVal.(string)
		for _, val := range validRoles {
			if val != roleVal.(string) {
				context.AbortWithStatusJSON(http.StatusOK, BaseResponse{StatusCode: ERROR, Message: TextLanguage("tokenError", lang)})
			}
		}
	}
}
