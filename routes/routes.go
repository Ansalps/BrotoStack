package routes

import (
	"github.com/Ansalps/BrotoStack/di"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	handler := di.DependencyInjection()
	router.GET("/", handler.SignUp)
}
