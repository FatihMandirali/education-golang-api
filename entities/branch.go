package entities

import (
	"github.com/jinzhu/gorm"
)

type Branch struct {
	gorm.Model  `json:"model"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
}
