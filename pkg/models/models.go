package models

import (
	"gorm.io/gorm"
)

type Admins struct{
	gorm.Model
	Email string	`gorm:"type:varchar(30);unique"`
	PasswordHash string `gorm:"type:varchar"`
}
type Otps struct{
	gorm.Model
	Email string `gorm:"type:varchar(30);unique"`
	Data string
	Is_Valid bool `gorm:"type:bool;default:true"`
}
type Users struct{
	gorm.Model		
	Username string `gorm:"type:varchar(30);unique"`
	Email string	`gorm:"type:varchar(30);unique"`
	PasswordHash string `gorm:"type:varchar"`
	IsVerified bool	`gorm:"type:bool;default:false"`
}
