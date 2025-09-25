package models

import (
	"gorm.io/gorm"
)

// type Admins struct{
// 	id uint
// 	username string
// 	email string
// 	passwordHash string
// 	created_at string
// 	updated_at string
// 	deleted_at string
// }
// type Otps struct{
// 	id uint
// 	role string
// 	data string
// 	created_at string
// 	updated_at string
// 	deleted_at string
// 	is_used bool
// }
type Users struct{
	gorm.Model		
	Username string `gorm:"type:varchar(30)"`
	Email string	`gorm:"type:varchar(30)"`
	PasswordHash string `gorm:"type:varchar"`
	IsVerified bool	`gorm:"type:bool;default:false"`
}
