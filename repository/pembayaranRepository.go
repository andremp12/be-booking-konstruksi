package repository

import (
	database "booking-konstruksi/database/migration"
	"booking-konstruksi/response"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type Pembayaran interface {
	GetRiwayatPembayaran() ([]database.Pembayaran, error)
	GetRiwayatPembayaranClient(clientId string) ([]database.Pembayaran, error)
	GetAllData() ([]database.Pembayaran, error)
	GetData(id string) (database.Pembayaran, error)
	GetLastData() (database.Pembayaran, error)
	GetPembayaranClient(konstruksiId string) ([]database.Pembayaran, error)
	Create(Pembayaran database.Pembayaran) (database.Pembayaran, error)
	Update(Pembayaran database.Pembayaran) (database.Pembayaran, error)
	Delete(id string) error
	SuccessPayment(status string, id string) error
	ConfirmationAdminPayment(id string) error
	GetTotalPaid() ([]response.TotalPaid, error)
	GetTotalPaidKonstruksi(konstruksi_id string) (response.TotalPaid, error)
	//UpdatePaidKonstruksi(id string, paid int) error
}

type repoPembayaran struct {
	db *gorm.DB
}

func NewRepoPembayaran(db *gorm.DB) *repoPembayaran {
	return &repoPembayaran{db: db}
}

func (r *repoPembayaran) GetRiwayatPembayaran() ([]database.Pembayaran, error) {
	var pembayarans []database.Pembayaran

	err := r.db.Preload("Konstruksi").Preload("Client").Order("id desc").Find(&pembayarans, "status = 'Paid'").Error

	return pembayarans, err
}

func (r *repoPembayaran) GetRiwayatPembayaranClient(clientId string) ([]database.Pembayaran, error) {
	var pembayarans []database.Pembayaran

	err := r.db.Preload("Konstruksi").Preload("Client").Order("id desc").Find(&pembayarans, "status = 'Paid' AND client_id = ?", clientId).Error

	return pembayarans, err
}

func (r *repoPembayaran) GetAllData() ([]database.Pembayaran, error) {
	var pembayarans []database.Pembayaran

	err := r.db.Preload("Konstruksi").Preload("Client").Order("id desc").Find(&pembayarans).Error

	return pembayarans, err
}

func (r *repoPembayaran) GetData(id string) (database.Pembayaran, error) {
	var pembayaran database.Pembayaran

	err := r.db.Preload("Konstruksi").Preload("Client").First(&pembayaran, id).Error

	return pembayaran, err
}

func (r *repoPembayaran) GetLastData() (database.Pembayaran, error) {
	var pembayaran database.Pembayaran

	err := r.db.Last(&pembayaran).Error

	return pembayaran, err
}

func (r *repoPembayaran) GetPembayaranClient(konstruksiId string) ([]database.Pembayaran, error) {
	var pembayaran []database.Pembayaran

	err := r.db.Debug().Preload("Konstruksi.TipeKonstruksi").Preload(clause.Associations).Order("id desc").Find(&pembayaran, "konstruksi_id = ?", konstruksiId).Error

	return pembayaran, err
}

func (r *repoPembayaran) Create(pembayaran database.Pembayaran) (database.Pembayaran, error) {
	err := r.db.Debug().Create(&pembayaran).Error

	return pembayaran, err
}

func (r *repoPembayaran) Delete(id string) error {
	var pembayaran database.Pembayaran

	err := r.db.Delete(&pembayaran, id).Error

	return err
}

func (r *repoPembayaran) Update(pembayaran database.Pembayaran) (database.Pembayaran, error) {
	err := r.db.Debug().Save(&pembayaran).Error

	return pembayaran, err
}

func (r *repoPembayaran) SuccessPayment(status string, id string) error {
	var pembayaran database.Pembayaran
	var konstruksi database.Konstruksi

	err := r.db.Debug().First(&pembayaran, id).Error
	err = r.db.Debug().Model(&pembayaran).Updates(map[string]interface{}{"status": status, "payment_date": time.Now()}).Error

	err = r.db.First(&konstruksi, pembayaran.KonstruksiID).Error
	err = r.db.Debug().Model(&konstruksi).Updates(map[string]interface{}{"paid": gorm.Expr("paid + ?", pembayaran.Total), "remaining_payment": gorm.Expr("remaining_payment - ?", pembayaran.Total)}).Error

	return err
}

func (r *repoPembayaran) GetTotalPaid() ([]response.TotalPaid, error) {
	var pembayarans []database.Pembayaran
	var totalPaid []response.TotalPaid

	err := r.db.Model(&pembayarans).Select("konstruksi_id, sum(total) as total").Group("konstruksi_id").Find(&totalPaid).Error

	return totalPaid, err
}

func (r *repoPembayaran) GetTotalPaidKonstruksi(konstruksi_id string) (response.TotalPaid, error) {
	var pembayarans database.Pembayaran
	var totalPaid response.TotalPaid

	err := r.db.Model(&pembayarans).Select("konstruksi_id, sum(total) as total").Group("konstruksi_id").Having("konstruksi_id = ?", konstruksi_id).First(&totalPaid).Error

	return totalPaid, err
}

func (r *repoPembayaran) ConfirmationAdminPayment(id string) error {
	var pembayaran database.Pembayaran

	err := r.db.Debug().Model(&pembayaran).Where("id = ?", id).Update("status", "Paid").Error

	return err
}

//func (r *repoPembayaran) UpdatePaidKonstruksi(id string, paid int) error {
//	var konstruksi database.Konstruksi
//
//	err := r.db.First(&konstruksi, id).Error
//	err = r.db.Debug().Model(&konstruksi).Updates(map[string]interface{}{"paid": gorm.Expr("paid + ?", paid), "remaining_payment": gorm.Expr("remaining_payment - ?", paid)}).Error
//
//	return err
//}
