package request

import (
	"education.api/enum"
	"time"
)

type AnnouncementCreateRequest struct {
	Title       string                `json:"title" binding:"required"`
	Description string                `json:"description" binding:"required"`
	Type        enum.AnnouncementEnum `json:"type" binding:"required"`
	UserId      int                   `json:"userId"`
	StartDate   time.Time             `json:"startDate" binding:"required"`
	EndDate     time.Time             `json:"endDate" binding:"required"`
}
