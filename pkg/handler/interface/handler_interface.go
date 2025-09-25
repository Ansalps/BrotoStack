package HandlerInterface

import "github.com/Ansalps/BrotoStack/pkg/models"

type User_Handler_Interface interface{
	ValidateSignUpRequest(models.UserSignUpRequest) error
	StoreUnverifiedUserInDb(models.UserSignUpRequest) error
}