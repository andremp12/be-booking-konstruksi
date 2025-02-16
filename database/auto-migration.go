package database

import (
	"booking-konstruksi/database/migration"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&database.TipeKonstruksi{})
	err = db.AutoMigrate(&database.Konstruksi{})
	err = db.AutoMigrate(&database.Pembayaran{})
	err = db.AutoMigrate(&database.Laporan{})
	err = db.AutoMigrate(&database.Role{})
	err = db.AutoMigrate(&database.User{})
	err = db.AutoMigrate(&database.Satuan{})
	err = db.AutoMigrate(&database.AuthUser{})
	err = db.AutoMigrate(&database.TimeLine{})

	return err
}
