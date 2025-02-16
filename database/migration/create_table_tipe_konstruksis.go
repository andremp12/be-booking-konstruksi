package database

import "time"

type TipeKonstruksi struct {
	ID          int `gorm:"type:int(11);primary_key;auto_increment"`
	Konstruksis []Konstruksi
	Name        string    `gorm:"type:varchar(50);not null"`
	HargaFull   int       `gorm:"type:bigInt;not null"`
	HargaJasa   int       `gorm:"type:bigInt;not null"`
	Image       string    `gorm:"type:varchar(255);null"`
	CreatedAt   time.Time `gorm:"type:timestamp;not null"`
	UpdatedAt   time.Time `gorm:"type:timestamp;not null"`
}
