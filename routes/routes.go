package routes

import (
	"github.com/Ansalps/BrotoStack/pkg/di"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	handler := di.DependencyInjection()
	router.GET("/", handler.SignUp)
}
