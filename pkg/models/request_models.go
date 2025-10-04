package models
type UserSignUpRequest struct{
	Username string	`json:"username" validate:"required,min=1,max=30,usernameRegex,startsnotwith=.,endsnotwith=.,noRepeatedPeriods"`
	Email string	`json:"email" validate:"required,email"`
	Password string	`json:"password" validate:"required,min=6"`
	Confirmpassword string	`json:"confirm_password" validate:"required,min=6,eqfield=Password"`
}
type OtpVerification struct{
	Email string `json:"email" validate:"required,email"`
	Data string `json:"otp"`
	Action string `json:"action" validate:"oneof=signup forget-password"`
}
type ResendOtpToEmail struct{
	Email string `json:"email" validate:"required,email"`
	Action string `json:"action" validate:"oneof=signup forget-password"`
}
type ForgetPassword struct{
	Email string `json:"email" validate:"required,email"`
}
type ResetPassword struct{
	Password string `json:"password" validate:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" validate:"required,min=6,eqfield=Password"`
}

type UserLogin struct{
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}