package repository

import (
	database "booking-konstruksi/database/migration"
	"gorm.io/gorm"
)

type Mandor interface {
	GetAllData() ([]database.User, error)
}

type repoMandor struct {
	db *gorm.DB
}

func NewRepoMandor(db *gorm.DB) *repoMandor {
	return &repoMandor{db: db}
}

func (r *repoMandor) GetAllData() ([]database.User, error) {
	var mandor []database.User

	err := r.db.Preload("Role").Find(&mandor, "role_id = ?", 15).Error

	return mandor, err
}
