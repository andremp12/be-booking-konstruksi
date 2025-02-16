package repository

import (
	database "booking-konstruksi/database/migration"
	"gorm.io/gorm"
)

type Satuan interface {
	GetAllData() ([]database.Satuan, error)
	GetData(id string) (database.Satuan, error)
	Create(Satuan database.Satuan) (database.Satuan, error)
	Update(Satuan database.Satuan) (database.Satuan, error)
	Delete(Satuan database.Satuan) error
}

type repoSatuan struct {
	db *gorm.DB
}

func NewRepoSatuan(db *gorm.DB) *repoSatuan {
	return &repoSatuan{db: db}
}

func (r *repoSatuan) GetAllData() ([]database.Satuan, error) {
	var Satuans []database.Satuan

	err := r.db.Find(&Satuans).Error

	return Satuans, err
}

func (r *repoSatuan) GetData(id string) (database.Satuan, error) {
	var Satuan database.Satuan

	err := r.db.First(&Satuan, id).Error

	return Satuan, err
}

func (r *repoSatuan) Create(Satuan database.Satuan) (database.Satuan, error) {
	err := r.db.Debug().Create(&Satuan).Error

	return Satuan, err
}

func (r *repoSatuan) Delete(Satuan database.Satuan) error {
	err := r.db.Delete(&Satuan).Error

	return err
}

func (r *repoSatuan) Update(Satuan database.Satuan) (database.Satuan, error) {
	err := r.db.Debug().Save(&Satuan).Error

	return Satuan, err
}
