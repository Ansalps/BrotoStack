package di

import (
	"github.com/Ansalps/BrotoStack/pkg/db"
	"github.com/Ansalps/BrotoStack/pkg/handler"
	"github.com/Ansalps/BrotoStack/pkg/repo"
	"github.com/Ansalps/BrotoStack/pkg/service"
)

func DependencyInjection() *handler.UserHandler {
	db := db.ConnectToDb()
	repo := repo.NewUserRepo(db)
	service := service.NewUserService(repo)
	handler := handler.NewUserHandler(service)
	return handler
}
