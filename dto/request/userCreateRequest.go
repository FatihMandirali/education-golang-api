package request

type UserCreateRequest struct {
	Name        string `json:"name" binding:"required"`
	Surname     string `json:"surname" binding:"required"`
	Password    string `json:"password" binding:"required"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	BranchId    int    `json:"branch"`
	ClassId     int    `json:"class"`
	LessonId    int    `json:"lesson"`
	CoverId     int    `json:"cover"`
	IsActive    bool   `json:"isActive"`
}
