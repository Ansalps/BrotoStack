package repo

import (
	"gorm.io/gorm"
)

type adminRepo struct {
	db *gorm.DB
}

func NewAdminReop(db *gorm.DB) *adminRepo {
	return &adminRepo{
		db: db,
	}
}
