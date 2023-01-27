package entities

import "github.com/jinzhu/gorm"

type Lesson struct {
	gorm.Model `json:"model"`
	Name       string `json:"name"`
}
