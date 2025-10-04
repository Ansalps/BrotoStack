package ServiceInterface

import (
	"time"

	"github.com/Ansalps/BrotoStack/pkg/models"
)

type User_Service_Interface interface {
	Store_Unverified_User(models.UserSignUpRequest)error
	CheckIfEmailExistsInOtp(string)(bool,error)
	Store_Otp_For_User(Otp,email string)error
	Overwrite_Otp_To_Email(otpo,email string)error
	Does_Username_Exist_In_DB(string)error
	Does_Email_Exist_In_DB(string)(bool,error)
	RemoveUnverifiedUsersOlderThan3Minutes(time.Duration)error
	CheckForUnverifiedUserInDB(string)error
	CheckIfOtpExists(otp,email string)(string,error)
	CheckIfOtpExpired(otp,email string)error
	VerifyUser(string)error
	InvalidateOtp(otp,email string)error
	Delete_Unverified_User_With_Same_Email(string)error
	ResetPassword(email,password string)error
	FetchStoredHashFromExistingUser(string)(string,error)
	FetchDetailsForExistingUser(string)(models.Users,error)
}
