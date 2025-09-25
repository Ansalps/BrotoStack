package service

import (
	"fmt"

	"github.com/Ansalps/BrotoStack/pkg/models"
	ServiceInterface "github.com/Ansalps/BrotoStack/pkg/service/interface"
	"github.com/Ansalps/BrotoStack/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userService ServiceInterface.User_Service_Interface
}

func NewUserService(repo ServiceInterface.User_Service_Interface) *UserService {
	return &UserService{
		userService: repo,
	}
}
func (u UserService) ValidateSignUpRequest(signuprequest models.UserSignUpRequest) error {
	err := utils.Validate(signuprequest)
	if err != nil {
		return err
	}
	return nil
}
func (u UserService) StoreUnverifiedUserInDb(signuprequest models.UserSignUpRequest) error {
	v, err := HashPassword(signuprequest.Password)
	if err != nil {
		return fmt.Errorf("Failed to hash user signup password, error : %v", err)
	}
	signuprequest.Confirmpassword = v
	err = u.userService.Store_Unverified_User(signuprequest)
	if err != nil {
		return fmt.Errorf("failed to insert the unverified user into database and database error is: %v:", err)
	}
	return nil
}
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
