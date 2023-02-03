package request

type StudentInstallmentPaymentRequest struct {
	StudentId int `json:"studentId"`
	PaymentId int `json:"paymentId"`
}
