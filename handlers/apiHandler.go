package handlers

import (
	_ "education.api/docs" //swagger sayfasının hata vermemesi için eklememiz lazım
	"education.api/enum"
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
	r.Use(middleware.CORSMiddleware())

	//Login
	r.POST("login", PostLogin)

	api := r.Group("/api")
	api.Use(middleware.ValidateToken())

	admins := api.Group("/admin")
	admins.Use(middleware.AuthorizationToken([]string{string(enum.Admin)}))
	admins.GET("list", GetAdmins)
	admins.POST("create", PostAdmin)
	admins.PUT("update", UpdateAdmin)
	admins.GET("/:id", GetAdminById)
	admins.DELETE("delete/:id", DeleteAdminById)

	covers := api.Group("/cover")
	covers.Use(middleware.AuthorizationToken([]string{string(enum.Admin)}))
	covers.GET("/", GetCover)
	covers.GET("/allList", GetCoverAllList)
	covers.POST("/", PostCover)
	covers.PUT("/", UpdateCover)
	covers.GET("/:id", GetCoverById)
	covers.DELETE("/:id", DeleteCoverById)

	teachers := api.Group("/teacher")
	teachers.Use(middleware.AuthorizationToken([]string{string(enum.Admin)}))
	teachers.GET("/", GetTeacher)
	teachers.POST("/", PostTeacher)
	teachers.PUT("/", UpdateTeacher)
	teachers.GET("/:id", GetTeacherById)
	teachers.DELETE("delete/:id", DeleteTeacherById)

	studentNotRecord := api.Group("/student")
	studentNotRecord.Use(middleware.AuthorizationToken([]string{string(enum.Admin)}))
	studentNotRecord.GET("/list", GetNotRecordStudent)
	studentNotRecord.POST("/create", PostNotRecordStudent)
	studentNotRecord.PUT("/update", UpdateNotRecordStudent)
	studentNotRecord.GET("/:id", GetNotRecordStudentById)
	studentNotRecord.DELETE("/:id", DeleteNotRecordStudentById)

	branchs := api.Group("/branch")
	branchs.Use(middleware.AuthorizationToken([]string{string(enum.Admin)}))
	branchs.GET("/list", GetBranchList)
	branchs.GET("/allList", GetAllBranchList)
	branchs.POST("/create", PostBranch)
	branchs.PUT("/update", UpdateBranch)
	branchs.GET("/:id", GetBranchById)
	branchs.DELETE("delete/:id", DeleteBranchById)

	announcements := api.Group("/announcement")
	announcements.Use(middleware.AuthorizationToken([]string{string(enum.Admin)}))
	announcements.GET("/", GetAnnouncementList)
	announcements.POST("/", PostAnnouncement)
	announcements.PUT("/", UpdateAnnouncement)
	announcements.GET("/:id", GetAnnouncementById)
	announcements.DELETE("/:id", DeleteAnnouncementById)

	classes := api.Group("/class")
	classes.Use(middleware.AuthorizationToken([]string{string(enum.Admin)}))
	classes.GET("/list", GetClassList)
	classes.GET("/allList", GetClassAllList)
	classes.POST("/create", PostClass)
	classes.PUT("/update", UpdateClass)
	classes.GET("/:id", GetClassById)
	classes.DELETE("delete/:id", DeleteClassById)

	studentApply := api.Group("/studentPayment")
	studentApply.Use(middleware.AuthorizationToken([]string{string(enum.Admin)}))
	studentApply.POST("/apply", PostRecordApplyStudent)
	studentApply.GET("/list/:id", GetStudentPaymentList)
	studentApply.POST("/installment", PostPaymentInstallment)

	lessons := api.Group("/lesson")
	lessons.Use(middleware.AuthorizationToken([]string{string(enum.Admin)}))
	lessons.GET("/", GetLessonList)
	lessons.POST("/", PostLesson)
	lessons.PUT("/", UpdateLesson)
	lessons.GET("/:id", GetLessonById)
	lessons.DELETE("/:id", DeleteLessonById)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
