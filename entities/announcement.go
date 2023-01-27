package entities

import (
	"education.api/enum"
	"github.com/jinzhu/gorm"
	"time"
)

type Announcement struct {
	gorm.Model  `json:"model"`
	Title       string                `json:"title"`
	Description string                `json:"description"`
	StartDate   time.Time             `json:"startDate"`
	EndDate     time.Time             `json:"endDate"`
	Type        enum.AnnouncementEnum `json:"type"`
	UserID      int                   `json:"userId"`
	User        User                  `json:"user"`
}
