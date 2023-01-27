package request

import (
	"education.api/enum"
	"time"
)

type AnnouncementUpdateRequest struct {
	Id          int                   `json:"id" binding:"required"`
	Title       string                `json:"title" binding:"required"`
	Description string                `json:"description" binding:"required"`
	Type        enum.AnnouncementEnum `json:"type" binding:"required"`
	UserId      int                   `json:"userId"`
	StartDate   time.Time             `json:"startDate" binding:"required"`
	EndDate     time.Time             `json:"endDate" binding:"required"`
}
