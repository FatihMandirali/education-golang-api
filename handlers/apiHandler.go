package handlers

import (
	_ "education.api/cmd/docs" //swagger sayfasının hata vermemesi için eklememiz lazım
	. "education.api/enum"
	"education.api/middleware"
	. "education.api/services"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// swagger kurulumu: https://santoshk.dev/posts/2022/how-to-integrate-swagger-ui-in-go-backend-gin-edition/
// swagger init hatası çözümü: https://github.com/swaggo/swag/issues/197
func Run() {
	r := gin.Default()
	//Login
	r.POST("/login", PostLogin)
	//Admins
	api := r.Group("/api")
	api.Use(middleware.ValidateToken())

	admins := api.Group("/admin")
	admins.Use(middleware.AuthorizationToken([]string{string(Admin)}))
	admins.GET("/", GetAdmins)
	admins.POST("/", PostAdmin)
	admins.PUT("/", UpdateAdmin)
	admins.GET("/:id", GetAdminById)
	admins.DELETE("/:id", DeleteAdminById)

	covers := api.Group("/cover")
	covers.Use(middleware.AuthorizationToken([]string{string(Admin)}))
	covers.GET("/", GetCover)
	covers.POST("/", PostCover)
	covers.PUT("/", UpdateCover)
	covers.GET("/:id", GetCoverById)
	covers.DELETE("/:id", DeleteCoverById)

	teachers := api.Group("/teacher")
	teachers.Use(middleware.AuthorizationToken([]string{string(Admin)}))
	teachers.GET("/", GetTeacher)
	teachers.POST("/", PostTeacher)
	teachers.PUT("/", UpdateTeacher)
	teachers.GET("/:id", GetTeacherById)
	teachers.DELETE("/:id", DeleteTeacherById)

	studentNotRecord := api.Group("/studentNotRecord")
	studentNotRecord.Use(middleware.AuthorizationToken([]string{string(Admin)}))
	studentNotRecord.GET("/", GetNotRecordStudent)
	studentNotRecord.POST("/", PostNotRecordStudent)
	studentNotRecord.PUT("/", UpdateNotRecordStudent)
	studentNotRecord.GET("/:id", GetNotRecordStudentById)
	studentNotRecord.DELETE("/:id", DeleteNotRecordStudentById)

	branchs := api.Group("/branch")
	branchs.Use(middleware.AuthorizationToken([]string{string(Admin)}))
	branchs.GET("/", GetBranchList)
	branchs.POST("/", PostBranch)
	branchs.PUT("/", UpdateBranch)
	branchs.GET("/:id", GetBranchById)
	branchs.DELETE("/:id", DeleteBranchById)

	announcements := api.Group("/announcement")
	announcements.Use(middleware.AuthorizationToken([]string{string(Admin)}))
	announcements.GET("/", GetAnnouncementList)
	announcements.POST("/", PostAnnouncement)
	announcements.PUT("/", UpdateAnnouncement)
	announcements.GET("/:id", GetAnnouncementById)
	announcements.DELETE("/:id", DeleteAnnouncementById)

	classes := api.Group("/class")
	classes.Use(middleware.AuthorizationToken([]string{string(Admin)}))
	classes.GET("/", GetClassList)
	classes.POST("/", PostClass)
	classes.PUT("/", UpdateClass)
	classes.GET("/:id", GetClassById)
	classes.DELETE("/:id", DeleteClassById)

	studentApply := api.Group("/studentPayment")
	studentApply.Use(middleware.AuthorizationToken([]string{string(Admin)}))
	studentApply.POST("/apply", PostRecordApplyStudent)
	studentApply.GET("/list/:id", GetStudentPaymentList)
	studentApply.POST("/installment", PostPaymentInstallment)

	lessons := api.Group("/lesson")
	lessons.Use(middleware.AuthorizationToken([]string{string(Admin)}))
	lessons.GET("/", GetLessonList)
	lessons.POST("/", PostLesson)
	lessons.PUT("/", UpdateLesson)
	lessons.GET("/:id", GetLessonById)
	lessons.DELETE("/:id", DeleteLessonById)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
