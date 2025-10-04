package routes

import (
	"log"
	"time"

	"github.com/Ansalps/BrotoStack/pkg/di"
	"github.com/Ansalps/BrotoStack/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	adminHandler, userHandler := di.DependencyInjection()
	err := userHandler.RemoveUnverifiedUsersOlderThan3Minutes(3 * time.Minute)
	if err != nil {
		log.Println("failed remove unverified users older than 3 minutes error is :", err)
	}

	router.POST("/api/v1/admin/signup", adminHandler.AdmnSignUp)

	v1:=router.Group("/api/v1/user")

	v1.POST("/signup", userHandler.UserSignUp)
	v1.POST("/verify-otp",userHandler.VerifyOtp)
	v1.POST("/resend-otp",userHandler.ResendOtp)
	v1.POST("/forget-password",userHandler.ForgetPassword)
	v1.POST("/reset-password",middleware.OtpAuthMiddleware,userHandler.ResetPassword)
	v1.POST("/refresh",middleware.AccessRegenerator)
	v1.POST("/login",userHandler.Login)
}
