package HandlerInterface

import "github.com/Ansalps/BrotoStack/pkg/models"

type Admin_Handler_Interface interface {
	ValidateAdminSignUpRequest(models.UserSignUpRequest) error
}
