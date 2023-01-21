package services

import (
	. "education.api/config"
	"education.api/dbconnect"
	. "education.api/dto/response"
	. "education.api/utils"
	"encoding/json"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
)

// admin list
func GetAdmins(w http.ResponseWriter, r *http.Request) {
	lang := r.Header.Get("Accept-Language")

	c := r.Header.Get("Authorization")
	claims := &TokenResponse{}
	tkn, error := jwt.ParseWithClaims(c, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(JwtSecret), nil
	})

	if !tkn.Valid || error != nil {
		res, err := json.Marshal(BaseResponse{StatusCode: 500, Message: TextLanguage("tokenError", lang)})
		CheckError(err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
		return
	}

	connection := dbconnect.DbInit()
	defer dbconnect.CloseDatabase(connection)

	res, err := json.Marshal(BaseResponse{StatusCode: 200})
	CheckError(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	return
}
