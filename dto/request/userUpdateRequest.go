package request

type UpdateUserRequest struct {
	Id          int    `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Surname     string `json:"surname" binding:"required"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	BranchId    int    `json:"branch"`
	ClassId     int    `json:"class"`
	CoverId     int    `json:"cover"`
	IsActive    bool   `json:"isActive"`
}
