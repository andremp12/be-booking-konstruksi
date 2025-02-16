package repository

import (
	database "booking-konstruksi/database/migration"
	"gorm.io/gorm"
)

type LaporanKonstruksi interface {
	GetLaporanKonstruksi(konstruksiId string) ([]database.Laporan, error)
	GetActivityKonstruksi(mandorId string) ([]database.Laporan, error)
	Create(konstruksi database.Laporan, status string) error
	//Update(konstruksi database.Laporan) error
	Delete(id string) error
}

type laporanKonstruksiRepository struct {
	db *gorm.DB
}

func NewRepositoryLaporanKonstruksi(db *gorm.DB) *laporanKonstruksiRepository {
	return &laporanKonstruksiRepository{db: db}
}

func (r *laporanKonstruksiRepository) GetLaporanKonstruksi(kosntruksiId string) ([]database.Laporan, error) {
	var laporan []database.Laporan

	err := r.db.Debug().Find(&laporan, "konstruksi_id = ?", kosntruksiId).Error

	return laporan, err
}

func (r *laporanKonstruksiRepository) GetActivityKonstruksi(mandorId string) ([]database.Laporan, error) {
	var Laporan []database.Laporan

	err := r.db.Debug().Order("id desc").Find(&Laporan, "mandor_id = ?", mandorId).Error

	return Laporan, err
}

func (r *laporanKonstruksiRepository) Create(laporan database.Laporan, status string) error {
	var konstruksi database.Konstruksi

	err := r.db.Debug().Create(&laporan).Error
	err = r.db.Model(&konstruksi).Where("id = ?", laporan.KonstruksiID).Update("status", status).Error
	return err
}

func (r *laporanKonstruksiRepository) Delete(id string) error {
	var laporan database.Laporan

	err := r.db.Where("id = ?", id).Delete(&laporan).Error

	return err
}
