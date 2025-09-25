package models
type UserSignUpRequest struct{
	Username string	`json:"username" validate:"required,min=1,max=30,usernameRegex,startsnotwith=.,endsnotwith=.,noRepeatedPeriods"`
	Email string	`json:"email" validate:"required,email"`
	Password string	`json:"password" validate:"required,min=6"`
	Confirmpassword string	`json:"confirm_password" validate:"required,min=6,eqfield=Password"`
}
type OtpVerification struct{
	Data string `json:"otp"`
}