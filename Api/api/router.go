package api

import (
	"api/api/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Auth service
// @version 1.0
// @description Auth service
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewGin(h *handler.Handler) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	auth := r.Group("/auth")
	{
		auth.POST("/signup", h.SignUp)
		auth.POST("/login", h.LogIn)
		auth.POST("/enter-email", h.EnterEmail)
		auth.POST("/change-password", h.ChangePassword)
		auth.POST("/forget-password", h.ForgetPassword)
		auth.POST("/reset-password", h.ResetPassword)
		auth.POST("/change-email", h.ChangeEmail)
		auth.POST("/verify-email", h.VerifyEmail)
		auth.POST("/validateToken", h.ValidateToken)
		auth.POST("/refreshToken", h.RefreshToken)
		auth.GET("/getProfile", h.GetProfile)
	}

	swaggerURL := ginSwagger.URL("http://localhost:8036/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler, swaggerURL))

	return r
}
