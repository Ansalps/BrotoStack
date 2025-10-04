package di

import (
	"github.com/Ansalps/BrotoStack/pkg/db"
	"github.com/Ansalps/BrotoStack/pkg/handler"
	"github.com/Ansalps/BrotoStack/pkg/repo"
	"github.com/Ansalps/BrotoStack/pkg/service"
)

func DependencyInjection() (*handler.AdminHandler,*handler.UserHandler) {
	db := db.ConnectToDb()
	userRepo := repo.NewUserRepo(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	adminRepo:=repo.NewAdminReop(db)
	adminService:=service.NewAdminService(adminRepo)
	adminHandler:=handler.NewAdminHanlder(adminService)
	return adminHandler,userHandler
}
