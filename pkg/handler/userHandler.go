package handler

import (
	"fmt"
	"time"

	HandlerInterface "github.com/Ansalps/BrotoStack/pkg/handler/interface"
	"github.com/Ansalps/BrotoStack/pkg/middleware"
	"github.com/Ansalps/BrotoStack/pkg/models"
	"github.com/Ansalps/BrotoStack/pkg/utils"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userhandler HandlerInterface.User_Handler_Interface
}

func NewUserHandler(service HandlerInterface.User_Handler_Interface) *UserHandler {
	return &UserHandler{
		userhandler: service,
	}
}

func (h UserHandler) UserSignUp(c *gin.Context) {
	var signuprequest models.UserSignUpRequest
	if err := c.BindJSON(&signuprequest); err != nil {
		utils.ErrorResponse(c, 400, "error in binding json", err)
	}

	err := h.userhandler.ValidateUserSignUpRequest(signuprequest)
	if err != nil {
		fmt.Println("check in handler if error is nil or not", err)
		utils.ErrorResponse(c, 400, "failed to validate user signup request", err)
		return
	}
	err = h.userhandler.StoreUnverifiedUserInDb(signuprequest)
	if err != nil {
		fmt.Println("checking in handler middle")
		utils.ErrorResponse(c, 500, "failed to store unverified user in database", err)
		return
	}
	err = h.userhandler.GenerateAndSendOtpToEmail(signuprequest.Email)
	if err != nil {
		utils.ErrorResponse(c, 500, "Error in Sending otp to email or storing otp in database", err)
		return
	}
	fmt.Println("checking in handler last")
	utils.SuccessResponse(c, 200, "unverified user successfully inserted into database. otp sent to email", signuprequest.Email)
}

func (u UserHandler) RemoveUnverifiedUsersOlderThan3Minutes(age time.Duration) error {
	err := u.userhandler.RemoveUnverifiedUsersOlderThan3Minutes(age)
	if err != nil {
		return err
	}
	return nil
}

func (u UserHandler) VerifyOtp(c *gin.Context) {
	var otp models.OtpVerification
	if err := c.BindJSON(&otp); err != nil {
		utils.ErrorResponse(c, 400, "error in binding json", err)
		return
	}
	fmt.Println("inhandler ",otp.Data,otp.Email,otp.Action)
	if err:=utils.Validate(otp); err!=nil{
		utils.ErrorResponse(c,400,"failed to validate otp verifecation request body",err)
		return
	}
	token,err := u.userhandler.VerifyOtp(otp.Data, otp.Email, otp.Action)
	if err != nil {
		utils.ErrorResponse(c, 400, "failed to verify otp", err)
		return
	}
	str:="otp verified successfully, user can login now "
	data:=make(map[string]string)
	if token!=""{
		data["token"]=token
		utils.SuccessResponse(c,200,"otp verifed succesfully, user can reset password now",data)
		return
	}
	utils.SuccessResponse(c, 200, str, nil)
}
func (u UserHandler) ResendOtp(c *gin.Context) {
	var resend models.ResendOtpToEmail
	if err := c.BindJSON(&resend); err != nil {
		utils.ErrorResponse(c, 400, "error in binding json", err)
		return
	}
	if err:=utils.Validate(resend); err!=nil{
		utils.ErrorResponse(c,400,"error in validateing resend otp request body",err)
		return
	}
	if resend.Action == "signup" {
		err := u.userhandler.CheckForUnverifiedUserInDB(resend.Email)
		if err != nil {
			utils.ErrorResponse(c, 400, "signup before trying to send otp", err)
			return
		}
	}
	if resend.Action == "forget-password" {
		err := u.userhandler.CheckIfUserExists(resend.Email)
		if err != nil {
			str := "user with email " + resend.Email + " didn't exist"
			utils.ErrorResponse(c, 400, str, err)
			return
		}
	}
	err := u.userhandler.GenerateAndSendOtpToEmail(resend.Email)
	if err != nil {
		utils.ErrorResponse(c, 500, "Error in sending email or storing otp in database", err)
		return
	}
	utils.SuccessResponse(c, 200, "Succesffully resend otp to email", nil)
}

func (u UserHandler) ForgetPassword(c *gin.Context) {
	var forgetPassword models.ForgetPassword
	if err := c.BindJSON(&forgetPassword); err != nil {
		utils.ErrorResponse(c, 400, "error in binding json", err)
		return
	}
	if err:=utils.Validate(forgetPassword); err!=nil{
		utils.ErrorResponse(c,400,"error in validating forget password request body",err)
		return
	}
	err := u.userhandler.CheckIfUserExists(forgetPassword.Email)
	if err != nil {
		utils.ErrorResponse(c, 400, "email does not exist", err)
		return
	}
	err = u.userhandler.GenerateAndSendOtpToEmail(forgetPassword.Email)
	if err != nil {
		utils.ErrorResponse(c, 500, "Error in sending email or storing otp in database", err)
		return
	}
	utils.SuccessResponse(c, 200, "succesfully send otp to email verify your email", nil)
}
func (u UserHandler)ResetPassword(c *gin.Context){
	email,ok:=c.Get("email")
	if !ok{
		utils.ErrorResponse(c,400,"email not set in context",nil)
		return
	}
	v,ok1:=email.(string)
	if !ok1{
		utils.ErrorResponse(c,400,"error in asserting email",nil)
		return
	}
	var resetpassword models.ResetPassword
	if err:=c.BindJSON(&resetpassword); err!=nil{
		utils.ErrorResponse(c,400,"error in binding json",nil)
		return
	}
	if err:=utils.Validate(resetpassword); err!=nil{
		utils.ErrorResponse(c,400,"error in validating reset password request",err)
		return
	}
	err:=u.userhandler.CheckIfUserExists(v)
	if err!=nil{
		utils.ErrorResponse(c,400,"user with this email "+v+" doesn't exist",nil)
		return
	}
	err=u.userhandler.ResetPassword(v,resetpassword.Password)
	if err!=nil{
		utils.ErrorResponse(c,500,"error resetting password",nil)
		return
	}
	utils.SuccessResponse(c,200,"successfully reset password for user with email "+v,nil)
}
func (u UserHandler) Login(c *gin.Context){
	var userlogin models.UserLogin
	fmt.Println("is it in on login")
	if err:=c.BindJSON(&userlogin); err!=nil{
		utils.ErrorResponse(c,400,"error in binding userlognin",err)
		return
	}
	if err:=utils.Validate(userlogin);err!=nil{
		utils.ErrorResponse(c,400,"error in validating user login request body",err)
		return
	}
	fmt.Println("is it reaching before check if user exists")
	err:=u.userhandler.CheckIfUserExists(userlogin.Email)
	if err!=nil{
		utils.ErrorResponse(c,400,"email or password incorrect",err)
		fmt.Println("hope here is the error")
		return
	}
	fmt.Println("make sure it is not here, is it reaching before check if user exists")
	user,err:=u.userhandler.PasswordMatchForExistingUser(userlogin.Email,userlogin.Password)
	if err!=nil{
		fmt.Println("is it in first error")
		utils.ErrorResponse(c,500,"error in querying for password match",err)
		return
	}
	fmt.Println(user)
	if user == (models.Users{}) {
		fmt.Println("is it in second error")
		utils.ErrorResponse(c,400,"email or password incorrect",err)
		return
	}
	fmt.Println("make sure it is not here")
	id:=utils.ConvertUintToId(user.ID)
	AccessTokenString,err:=middleware.GenerateAccessToken(id)
	if err!=nil{
		utils.ErrorResponse(c,500,"failed to generate access token",nil)
		return
	}
	RefreshTokenString,err:=middleware.GenerateRefreshToken(id)
	if err!=nil{
		utils.ErrorResponse(c,500,"failed to generate refresh token",nil)
		return
	}
	data:=map[string]string{
		"Access_Token":AccessTokenString,
		"Refresh_Token":RefreshTokenString,
	}
	utils.SuccessResponse(c,200,"user logged in successfully",data)
}