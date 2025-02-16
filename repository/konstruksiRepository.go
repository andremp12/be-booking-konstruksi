package repository

import (
	database "booking-konstruksi/database/migration"
	"booking-konstruksi/response"
	"gorm.io/gorm"
)

type Konstruksi interface {
	GetCountStatus() ([]response.CountStatusKonstruksi, error)
	GetRiwayatKonstruksi() ([]database.Konstruksi, error)
	GetRiwayatKonstruksiMandor(clientId string) ([]database.Konstruksi, error)
	GetRiwayatKonstruksiClient(mandorId string) ([]database.Konstruksi, error)
	GetAllData() ([]database.Konstruksi, error)
	GetData(id string) (database.Konstruksi, error)
	GetKonstruksiUser(clientId string) ([]database.Konstruksi, error)
	GetKonstruksiMandor(mandorId string) ([]database.Konstruksi, error)
	GetLastData() (database.Konstruksi, error)
	Booking(Konstruksi database.Konstruksi) (database.Konstruksi, error)
	UpdateStatus(id string, status string) error
	KonfirmasiBooking(konstruksi database.Konstruksi, id string) (database.Konstruksi, error)
	Delete(id string) error
}

type repoKonstruksi struct {
	db *gorm.DB
}

func NewRepoKonstruksi(db *gorm.DB) *repoKonstruksi {
	return &repoKonstruksi{db: db}
}

func (r *repoKonstruksi) GetCountStatus() ([]response.CountStatusKonstruksi, error) {
	var response []response.CountStatusKonstruksi
	var konstruksi []database.Konstruksi

	err := r.db.Model(&konstruksi).Select("status,COUNT(*) as total").Group("status").Find(&response).Error

	return response, err
}

func (r *repoKonstruksi) GetRiwayatKonstruksi() ([]database.Konstruksi, error) {
	var konstruksis []database.Konstruksi

	err := r.db.Debug().Preload("Client").Preload("Mandor").Preload("TipeKonstruksi").Order("id desc").Find(&konstruksis, "status = 'Closed'").Error

	return konstruksis, err
}

func (r *repoKonstruksi) GetRiwayatKonstruksiMandor(mandorId string) ([]database.Konstruksi, error) {
	var konstruksis []database.Konstruksi

	err := r.db.Debug().Preload("Client").Preload("Mandor").Preload("TipeKonstruksi").Order("id desc").Find(&konstruksis, "status = 'Closed' AND mandor_id=?", mandorId).Error

	return konstruksis, err
}

func (r *repoKonstruksi) GetRiwayatKonstruksiClient(clientId string) ([]database.Konstruksi, error) {
	var konstruksis []database.Konstruksi

	err := r.db.Debug().Preload("Client").Preload("Mandor").Preload("TipeKonstruksi").Order("id desc").Find(&konstruksis, "status = 'Finished' AND client_id=?", clientId).Error

	return konstruksis, err
}

func (r *repoKonstruksi) GetAllData() ([]database.Konstruksi, error) {
	var Konstruksis []database.Konstruksi

	err := r.db.Preload("Client").Preload("Mandor").Preload("TipeKonstruksi").Order("id desc").Find(&Konstruksis).Error

	return Konstruksis, err
}

func (r *repoKonstruksi) GetData(id string) (database.Konstruksi, error) {
	var Konstruksi database.Konstruksi

	err := r.db.Preload("Client").Preload("Mandor").Preload("TipeKonstruksi").First(&Konstruksi, id).Error

	return Konstruksi, err
}

func (r *repoKonstruksi) GetKonstruksiUser(clientId string) ([]database.Konstruksi, error) {
	var Konstruksi []database.Konstruksi

	err := r.db.Debug().Preload("Client").Preload("Mandor").Preload("TipeKonstruksi").Order("id desc").Find(&Konstruksi, "client_id = ?", clientId).Error

	return Konstruksi, err
}

func (r *repoKonstruksi) GetKonstruksiMandor(mandorId string) ([]database.Konstruksi, error) {
	var Konstruksi []database.Konstruksi

	err := r.db.Debug().Preload("Client").Preload("Mandor").Preload("TipeKonstruksi").Order("id desc").Find(&Konstruksi, "mandor_id = ?", mandorId).Error

	return Konstruksi, err
}

func (r *repoKonstruksi) GetLastData() (database.Konstruksi, error) {
	var konstruksi database.Konstruksi

	err := r.db.Last(&konstruksi).Error

	return konstruksi, err
}

func (r *repoKonstruksi) Booking(Konstruksi database.Konstruksi) (database.Konstruksi, error) {
	err := r.db.Debug().Create(&Konstruksi).Error

	return Konstruksi, err
}

func (r *repoKonstruksi) UpdateStatus(id string, status string) error {
	var konstruksi database.Konstruksi

	err := r.db.Model(&konstruksi).Where("id = ?", id).Update("status", status).Error

	return err
}

func (r *repoKonstruksi) KonfirmasiBooking(konstruksiUpdate database.Konstruksi, id string) (database.Konstruksi, error) {
	var konstruksi database.Konstruksi

	err := r.db.First(&konstruksi, id).Error

	err = r.db.Debug().Model(&konstruksi).Select("mandor_id", "status", "start_date", "end_date").Updates(map[string]interface{}{"mandor_id": konstruksiUpdate.MandorID, "status": konstruksiUpdate.Status, "start_date": konstruksiUpdate.StartDate, "end_date": konstruksiUpdate.EndDate}).Error
	//err = r.db.Debug().Model(&konstruksi).Updates(konstruksiUpdate).Error

	return konstruksi, err
}

func (r *repoKonstruksi) Delete(id string) error {
	var konstruksi database.Konstruksi

	err := r.db.First(&konstruksi, id).Error

	err = r.db.Delete(&konstruksi).Error

	return err
}
