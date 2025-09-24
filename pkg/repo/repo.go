package repo

import (

	"gorm.io/gorm"
)
type Repo struct{
	db *gorm.DB
}
func NewUserRepo(db *gorm.DB) *Repo{
	return &Repo{
		db: db,
	}
}