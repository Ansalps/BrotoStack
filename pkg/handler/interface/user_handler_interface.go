package HandlerInterface

import (
	"time"

	"github.com/Ansalps/BrotoStack/pkg/models"
)

type User_Handler_Interface interface{
	ValidateUserSignUpRequest(models.UserSignUpRequest) error
	StoreUnverifiedUserInDb(models.UserSignUpRequest) error
	GenerateAndSendOtpToEmail(string) error
	RemoveUnverifiedUsersOlderThan3Minutes(time.Duration)error
	VerifyOtp(otp,email,action string)(string,error)
	CheckIfUserExists(string)error
	CheckForUnverifiedUserInDB(string)error
	ResetPassword(email,password string)error
	PasswordMatchForExistingUser(email,password string)(models.Users,error)
}