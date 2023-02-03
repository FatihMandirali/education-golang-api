package request

import (
	"education.api/enum"
	"time"
)

type ApplyStudentRecordRequest struct {
	StudentId        int                          `json:"studentId" binding:"required"`
	TotalAmount      float32                      `json:"totalAmount" binding:"required"`
	FirstAmount      float32                      `json:"firstAmount"`
	PaymentType      enum.StudentApplyPaymentEnum `json:"paymentType"`
	InstallmentCount int                          `json:"installmentCount"`
	StartDate        time.Time                    `json:"startDate"`
}
