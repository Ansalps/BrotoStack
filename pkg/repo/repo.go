package repo

import (
	"github.com/Ansalps/BrotoStack/pkg/models"
	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *Repo {
	return &Repo{
		db: db,
	}
}
func (r *Repo) Store_Unverified_User(signuprequest models.UserSignUpRequest) error {
	user:=models.Users{
		Username: signuprequest.Username,
		Email: signuprequest.Email,
		PasswordHash: signuprequest.Confirmpassword,
	}
	if err:= r.db.Create(&user).Error; err!=nil{
		return err
	}
	return nil
}
