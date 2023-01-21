package services

import (
	. "education.api/config"
	"education.api/dbconnect"
	. "education.api/dto/request"
	. "education.api/dto/response"
	. "education.api/entities"
	. "education.api/utils"
	"encoding/json"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

func PostLogin(W http.ResponseWriter, r *http.Request) {
	lang := r.Header.Get("Accept-Language")
	connection := dbconnect.DbInit()
	defer dbconnect.CloseDatabase(connection)

	var loginRequest LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	CheckError(err)

	var user User
	connection.First(&user, "email = ?", loginRequest.Email)
	if user.Email == "" {
		CheckError(err)
	}
	match := CheckPasswordHash(loginRequest.Password, user.Password)
	if match == false {
		res, err := json.Marshal(BaseResponse{StatusCode: 500, Message: TextLanguage("loginInfoError", lang)})
		CheckError(err)

		W.Header().Set("Content-Type", "application/json")
		W.WriteHeader(http.StatusOK)
		W.Write(res)
		return
	}
	expirationTime := time.Now().Add(1 * time.Minute)
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
		W.WriteHeader(http.StatusInternalServerError)
		return
	}
	//signedToken := []byte(tokenString)
	res, err := json.Marshal(BaseResponse{Data: tokenString, StatusCode: 200})
	CheckError(err)

	W.Header().Set("Content-Type", "application/json")
	W.WriteHeader(http.StatusOK)
	W.Write(res)
}
