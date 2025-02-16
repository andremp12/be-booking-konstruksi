package repository

import (
	database "booking-konstruksi/database/migration"
	"gorm.io/gorm"
)

type Client interface {
	GetAllData() ([]database.User, error)
}

type repoClient struct {
	db *gorm.DB
}

func NewRepoClient(db *gorm.DB) *repoClient {
	return &repoClient{db: db}
}

func (r *repoClient) GetAllData() ([]database.User, error) {
	var client []database.User

	err := r.db.Preload("Role").Find(&client, "role_id = ?", 16).Error

	return client, err
}
