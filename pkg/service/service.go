package service

import "github.com/Ansalps/BrotoStack/pkg/repo"

type UserService struct{
	Service *repo.Repo
}
func NewUserService(repo *repo.Repo)*UserService{
	return &UserService{
		Service: repo,
	}
}