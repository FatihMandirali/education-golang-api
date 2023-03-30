package request

type UpdateUserRequest struct {
	Id          int    `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Surname     string `json:"surname" binding:"required"`
	Password    string `json:"password" binding:"required"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	BranchId    int    `json:"branchId"`
	ClassId     int    `json:"classId"`
	LessonId    int    `json:"lessonId"`
	IsActive    bool   `json:"isActive"`
}
