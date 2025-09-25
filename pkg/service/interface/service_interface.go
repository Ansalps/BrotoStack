package ServiceInterface

import (
	"github.com/Ansalps/BrotoStack/pkg/models"
	
)

type User_Service_Interface interface {
	Store_Unverified_User(models.UserSignUpRequest)error
}
