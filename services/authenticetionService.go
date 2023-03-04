package services

import (
	. "education.api/config"
	"education.api/dbconnect"
	. "education.api/dto/request"
	. "education.api/dto/response"
	. "education.api/entities"
	. "education.api/generic"
	. "education.api/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// PostLogin godoc
// @Summary Login servisi
// @Produce json
// @Accept json
// @Success 200
// @Failure 400
// @Failure 500
// @Router /login/ [post]
func PostLogin(context *gin.Context) {
	lang := context.Request.Header.Get("Accept-Language")
	connection := dbconnect.DbInit()
	defer dbconnect.CloseDatabase(connection)

	var loginRequest LoginRequest
	err := json.NewDecoder(context.Request.Body).Decode(&loginRequest)
	CheckError(err, context, TextLanguage("loginInfoError", lang))

	var user User
	connection.First(&user, "email = ?", loginRequest.Email)
	if user.Email == "" {
		CheckError(err, context, TextLanguage("loginInfoError", lang))
	}
	match := CheckPasswordHash(loginRequest.Password, user.Password)
	if match == false {
		GenericResponse(context, ERROR, TextLanguage("loginInfoError", lang), nil)
		return
	}
	expirationTime := time.Now().Add(1000000 * time.Minute)
	claims := &TokenResponse{
		Email: user.Email,
		Role:  user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	var jwtKey = []byte(JwtSecret)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		GenericResponse(context, ERROR, TextLanguage("loginInfoError", lang), nil)
		return
	}
	GenericResponse(context, SUCCESS, "", tokenString)
}
