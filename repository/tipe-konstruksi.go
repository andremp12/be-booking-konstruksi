package repository

import (
	database "booking-konstruksi/database/migration"
	"gorm.io/gorm"
)

type TipeKonstruksi interface {
	GetTipeLanding() ([]database.TipeKonstruksi, error)
	GetAllData() ([]database.TipeKonstruksi, error)
	GetData(id string) (database.TipeKonstruksi, error)
	Create(tipeKonstruksi database.TipeKonstruksi) (database.TipeKonstruksi, error)
	Update(tipeKonstruksi database.TipeKonstruksi) (database.TipeKonstruksi, error)
	Delete(tipeKonstruksi database.TipeKonstruksi) error
}

type repoTipeKonstruksi struct {
	db *gorm.DB
}

func NewRepoTipeKonstruksi(db *gorm.DB) *repoTipeKonstruksi {
	return &repoTipeKonstruksi{db: db}
}

func (r *repoTipeKonstruksi) GetTipeLanding() ([]database.TipeKonstruksi, error) {
	var tipeKonstruksis []database.TipeKonstruksi

	err := r.db.Order("id desc").Limit(4).Find(&tipeKonstruksis).Error

	return tipeKonstruksis, err
}

func (r *repoTipeKonstruksi) GetAllData() ([]database.TipeKonstruksi, error) {
	var tipeKonstruksis []database.TipeKonstruksi

	err := r.db.Find(&tipeKonstruksis).Error

	return tipeKonstruksis, err
}

func (r *repoTipeKonstruksi) GetData(id string) (database.TipeKonstruksi, error) {
	var tipeKonstruksi database.TipeKonstruksi

	err := r.db.First(&tipeKonstruksi, id).Error

	return tipeKonstruksi, err
}

func (r *repoTipeKonstruksi) Create(tipeKonstruksi database.TipeKonstruksi) (database.TipeKonstruksi, error) {
	err := r.db.Debug().Create(&tipeKonstruksi).Error

	return tipeKonstruksi, err
}

func (r *repoTipeKonstruksi) Delete(tipeKonstruksi database.TipeKonstruksi) error {
	err := r.db.Debug().Delete(&tipeKonstruksi).Error

	return err
}

func (r *repoTipeKonstruksi) Update(tipeKonstruksi database.TipeKonstruksi) (database.TipeKonstruksi, error) {
	err := r.db.Debug().Save(&tipeKonstruksi).Error

	return tipeKonstruksi, err
}
