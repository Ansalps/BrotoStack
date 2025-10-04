package service

import (
	"errors"
	"fmt"
	"log"
	"net/smtp"
	"time"

	"github.com/Ansalps/BrotoStack/pkg/middleware"
	"github.com/Ansalps/BrotoStack/pkg/models"
	ServiceInterface "github.com/Ansalps/BrotoStack/pkg/service/interface"
	"github.com/Ansalps/BrotoStack/pkg/utils"
	"github.com/robfig/cron/v3"
)

type userService struct {
	userService ServiceInterface.User_Service_Interface
}

func NewUserService(repo ServiceInterface.User_Service_Interface) *userService {
	return &userService{
		userService: repo,
	}
}
func (u userService) ValidateUserSignUpRequest(signuprequest models.UserSignUpRequest) error {
	fmt.Println("ValidateUserSignUpReques")
	err := utils.Validate(signuprequest)
	if err != nil {
		return err
	}
	fmt.Println("ValidateUserSignUpReques 222")
	err = u.userService.Delete_Unverified_User_With_Same_Email(signuprequest.Email)
	exists, err := u.userService.Does_Email_Exist_In_DB(signuprequest.Email)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("email '%s' already exists choose another email", signuprequest.Email)
	}
	err = u.userService.Does_Username_Exist_In_DB(signuprequest.Username)
	if err != nil {
		return err
	}
	return nil
}
func (u userService) StoreUnverifiedUserInDb(signuprequest models.UserSignUpRequest) error {
	v, err := utils.HashPassword(signuprequest.Password)
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
func (u userService) GenerateAndSendOtpToEmail(email string) error {

	otp := utils.GenerateOtp()
	fmt.Println("generated otp in backend is", otp)
	auth := smtp.PlainAuth(
		"",
		"brotostack@gmail.com",
		"issy nngo mcal yzrj",
		"smtp.gmail.com",
	)
	msg := "Subject: Otp verification\nThis is and otp verification request for your email from BrotoStack. Yout otp is "+otp
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"brotostack@gmail.com",
		[]string{email},
		[]byte(msg),
	)
	if err != nil {
		return fmt.Errorf("Error in sending otp to email,erro is: %w", err)
	}
	exists,err:=u.userService.CheckIfEmailExistsInOtp(email)
	if err!=nil{
		return err
	}
	if !exists{
		err = u.userService.Store_Otp_For_User(otp, email)
		if err != nil {
			return fmt.Errorf("Error in storing in otp information in otp table, error is: %w", err)
		}
	} else{
		err:=u.userService.Overwrite_Otp_To_Email(otp,email)
		if err!=nil{
			return err
		}
	}
	return nil
}

func (u userService) RemoveUnverifiedUsersOlderThan3Minutes(age time.Duration) error {
	c := cron.New()
	_, err := c.AddFunc("3 * * * *", func() {
		err := u.userService.RemoveUnverifiedUsersOlderThan3Minutes(age)
		if err != nil {
			log.Println(err)
		}
	})
	if err != nil {
		log.Println(err)
	}
	c.Start()
	return nil
}
func (u userService) VerifyOtp(otp, email, action string) (string,error) {
	if action == "signup" {
		err := u.userService.CheckForUnverifiedUserInDB(email)
		if err != nil {
			return "",err
		}
	}
	if action == "forget-password" {
		exists, err := u.userService.Does_Email_Exist_In_DB(email)
		if err != nil {
			return "",err
		}
		if !exists {
			return "",fmt.Errorf("user with email '%s' doesn't exist", email)
		}
	}
	fmt.Println("inservice",otp,email,action)
	OTP, err := u.userService.CheckIfOtpExists(otp, email)
	if err != nil {
		return "",err
	}
	if OTP == otp {
		err := u.userService.CheckIfOtpExpired(otp, email)
		if err != nil {
			return "",err
		}
		if action == "signup" {
			err = u.userService.VerifyUser(email)
			if err != nil {
				return "",err
			}
		}
		err = u.userService.InvalidateOtp(otp, email)
		if err != nil {
			return "",err
		}
	} else {
		return "",errors.New("otp entered is wrong")
	}
	if action == "forget-password" {
		token, err := middleware.CreateTokenForResetPassword(email)
		if err != nil {
			return "",err
		} else{
			return token,nil
		}
	}
	return "",nil
}

func (u userService) CheckIfUserExists(email string) error {
	exists, err := u.userService.Does_Email_Exist_In_DB(email)
	if err != nil {
		fmt.Println("is ther error in service",err)
		return err
	}
	if !exists {
		return fmt.Errorf("user with email '%s' didn't exist", email)
	}
	return nil
}
func (u userService) CheckForUnverifiedUserInDB(email string) error {
	err := u.userService.CheckForUnverifiedUserInDB(email)
	if err != nil {
		return err
	}
	return nil
}
func (u userService)ResetPassword(email,password string)error{
	Password,err:=utils.HashPassword(password)
	if err != nil {
		return err
	}
	err=u.userService.ResetPassword(email,Password)
	if err!=nil{
		return err
	}
	return nil
}
func (u userService)PasswordMatchForExistingUser(email,password string)(models.Users,error){
	passwordHash,err:=u.userService.FetchStoredHashFromExistingUser(email)
	if err!=nil{
		return models.Users{},err
	}
	fmt.Println("",passwordHash)
	err=utils.CheckPasswordHash(password,passwordHash)
	if err!=nil{
		return models.Users{},err
	}
	user,err:=u.userService.FetchDetailsForExistingUser(email)
	return user,nil
}