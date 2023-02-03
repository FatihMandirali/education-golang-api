package services

import (
	. "education.api/config"
	"education.api/dbconnect"
	. "education.api/dto/request"
	. "education.api/entities"
	"education.api/enum"
	. "education.api/generic"
	"education.api/utils"
	"github.com/gin-gonic/gin"
	gormbulk "github.com/t-tiger/gorm-bulk-insert/v2"
	"time"
)

// record  apply Student
func PostRecordApplyStudent(context *gin.Context) {
	lang := context.Keys["Lang"]
	body := ApplyStudentRecordRequest{}
	if err := context.ShouldBindJSON(&body); err != nil {
		utils.CheckError(err, context, utils.TextLanguage("badRequest", lang.(string)))
		return
	}
	connection := dbconnect.DbInit()
	defer dbconnect.CloseDatabase(connection)
	var user User
	connection.Where("id = ?", body.StudentId).First(&user)
	if user.Email == "" {
		GenericResponse(context, ERROR, utils.TextLanguage("notFound", lang.(string)), nil)
		return
	}

	timeNow := time.Now()

	user.IsRecord = true
	user.TotalAmount = body.TotalAmount
	user.FirstAmount = body.FirstAmount
	user.Installment = body.InstallmentCount
	user.RecordDate = timeNow
	user.PaymentType = body.PaymentType
	connection.Save(&user)

	var studentPayment []interface{}
	if body.PaymentType == enum.Advance {
		newStudentPayment := StudentPayment{
			StudentID:        body.StudentId,
			PaymentType:      enum.Advance,
			PaymentDate:      timeNow,
			IsPayment:        true,
			Amount:           body.TotalAmount,
			PaymentApplyDate: timeNow,
		}
		studentPayment = append(studentPayment, newStudentPayment)
	} else {
		if body.FirstAmount > 0 {
			newStudentPayment := StudentPayment{
				StudentID:        body.StudentId,
				PaymentType:      enum.FirstAmount,
				PaymentDate:      timeNow,
				IsPayment:        true,
				Amount:           body.FirstAmount,
				PaymentApplyDate: timeNow,
			}
			studentPayment = append(studentPayment, newStudentPayment)
		}

		for i := 1; i <= body.InstallmentCount; i++ {
			paymentType := enum.Installment
			paymentDate := body.StartDate
			if i > 1 {
				paymentDate = body.StartDate.AddDate(0, i-1, 0)
			}
			newStudentPayment := StudentPayment{
				StudentID:   body.StudentId,
				PaymentType: enum.StudentApplyPaymentEnum(paymentType),
				PaymentDate: paymentDate,
				IsPayment:   false,
				Amount:      (body.TotalAmount - body.FirstAmount) / float32(body.InstallmentCount),
			}
			studentPayment = append(studentPayment, newStudentPayment)

		}
	}

	err := gormbulk.BulkInsert(connection, studentPayment, 3000)
	if err != nil {
		GenericResponse(context, ERROR, utils.TextLanguage("error", lang.(string)), nil)
		return
	}
	GenericResponse(context, SUCCESS, "", nil)
}

// student payment list
func GetStudentPaymentList(context *gin.Context) {
	lang := context.Keys["Lang"]

	connection := dbconnect.DbInit()
	defer dbconnect.CloseDatabase(connection)

	uri := IdRequest{}
	if err := context.BindUri(&uri); err != nil {
		utils.CheckError(err, context, utils.TextLanguage("badRequest", lang.(string)))
		return
	}

	var userPayments []StudentPayment
	connection.Preload("Student").Where("student_id = ?", uri.Id).Find(&userPayments)

	GenericResponse(context, SUCCESS, "", userPayments)
}

func PostPaymentInstallment(context *gin.Context) {
	lang := context.Keys["Lang"]
	body := StudentInstallmentPaymentRequest{}
	if err := context.ShouldBindJSON(&body); err != nil {
		utils.CheckError(err, context, utils.TextLanguage("badRequest", lang.(string)))
		return
	}
	connection := dbconnect.DbInit()
	defer dbconnect.CloseDatabase(connection)
	var studentPayment StudentPayment
	connection.Where("student_id = ?", body.StudentId).Where("id = ?", body.PaymentId).First(&studentPayment)
	if studentPayment.ID == 0 {
		GenericResponse(context, ERROR, utils.TextLanguage("notFound", lang.(string)), nil)
		return
	}

	studentPayment.IsPayment = true
	studentPayment.PaymentApplyDate = time.Now()
	connection.Save(&studentPayment)
	GenericResponse(context, SUCCESS, "", nil)
}
