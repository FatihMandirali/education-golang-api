package response

import (
	"education.api/entities"
	"time"
)

type AdminResponse struct {
	Id          int       `json:"id"`
	Email       string    `json:"email"`
	Name        string    `json:"name"`
	Surname     string    `json:"surname"`
	PhoneNumber string    `json:"phoneNumber"`
	CreateDate  time.Time `json:"createDate"`
	IsActive    bool      `json:"isActive"`
}

func CreateAdminResponse(user entities.User) AdminResponse {
	return AdminResponse{
		Id:          int(user.ID),
		Name:        user.Name,
		Email:       user.Email,
		Surname:     user.Surname,
		PhoneNumber: user.PhoneNumber,
		CreateDate:  user.CreatedAt,
		IsActive:    user.DeletedAt == nil,
	}
}

type AdminListResponse []*AdminResponse

func CreateAdminListResponse(users []entities.User) AdminListResponse {
	adminList := AdminListResponse{}
	for _, p := range users {
		admin := CreateAdminResponse(p)
		adminList = append(adminList, &admin)
	}
	return adminList
}
