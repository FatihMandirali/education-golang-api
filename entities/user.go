package entities

import (
	"education.api/enum"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model `json:"model"`
	Role       enum.RoleEnum `json:"role"`
	Email      string        `json:"email"`
	Name       string        `json:"name"`
	Surname    string        `json:"surname"`
	Password   string        `json:"password"`
}
