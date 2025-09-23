package di

import (
	"github.com/Ansalps/BrotoStack/db"
	"github.com/Ansalps/BrotoStack/handler"
	"github.com/Ansalps/BrotoStack/repo"
	"github.com/Ansalps/BrotoStack/service"
)

func DependencyInjection() *handler.UserHandler {
	db := db.ConnectToDb()
	repo := repo.NewUserRepo(db)
	service := service.NewUserService(repo)
	handler := handler.NewUserHandler(service)
	return handler
}
