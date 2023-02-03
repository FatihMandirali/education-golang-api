package entities

import (
	"education.api/enum"
	"github.com/jinzhu/gorm"
	"time"
)

type StudentPayment struct {
	gorm.Model       `json:"model"`
	PaymentType      enum.StudentApplyPaymentEnum `json:"paymentType"`
	PaymentDate      time.Time                    `json:"paymentDate"`
	PaymentApplyDate time.Time                    `json:"paymentApplyDate"`
	StudentID        int                          `json:"studentID"`
	Student          User                         `json:"student"`
	Amount           float32                      `json:"amount"`
	IsPayment        bool                         `json:"isPayment"`
}
