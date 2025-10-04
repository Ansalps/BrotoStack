package service

import (
	"github.com/Ansalps/BrotoStack/pkg/models"
	ServiceInterface "github.com/Ansalps/BrotoStack/pkg/service/interface"
	"github.com/Ansalps/BrotoStack/pkg/utils"
)

type adminService struct {
	adminService ServiceInterface.Admin_Service_Interface
}

func NewAdminService(repo ServiceInterface.Admin_Service_Interface) *adminService {
	return &adminService{
		adminService: repo,
	}
}

func (u adminService) ValidateAdminSignUpRequest(signuprequest models.UserSignUpRequest) error {
	err := utils.Validate(signuprequest)
	if err != nil {
		return err
	}
	return nil
}
