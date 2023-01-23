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

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
