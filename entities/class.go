package entities

import (
	"github.com/jinzhu/gorm"
)

type Class struct {
	gorm.Model `json:"model"`
	Name       string `json:"name"`
	BranchID   int    `json:"branchId"`
	Branch     Branch `json:"branch"`
	IsActive   bool   `json:"isActive"`
	//TODO: SINIF ÖĞRETMENİ EKLE
}
