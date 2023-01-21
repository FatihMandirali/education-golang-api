package response

import (
	"education.api/enum"
	"github.com/golang-jwt/jwt/v4"
)

type TokenResponse struct {
	Email string        `json:"email"`
	Role  enum.RoleEnum `json:"role"`
	jwt.RegisteredClaims
}
