package entities

import (
	"education.api/enum"
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model  `json:"model"`
	Role        enum.RoleEnum                `json:"role"`
	Email       string                       `json:"email"`
	Name        string                       `json:"name"`
	Surname     string                       `json:"surname"`
	Password    string                       `json:"password"`
	PhoneNumber string                       `json:"phoneNumber"`
	BranchID    int                          `json:"branchId"`
	Branch      Branch                       `json:"branch"`
	LessonID    int                          `json:"lessonId"`
	Lesson      Lesson                       `json:"lesson"`
	ClassID     int                          `json:"classId"`
	Class       Class                        `json:"class"`
	IsRecord    bool                         `json:"isRecord"`
	TotalAmount float32                      `json:"totalAmount"`
	FirstAmount float32                      `json:"firstAmount"`
	PaymentType enum.StudentApplyPaymentEnum `json:"paymentType"`
	Installment int                          `json:"installment"`
	RecordDate  time.Time                    `json:"recordDate"`
}
