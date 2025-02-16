package database

import "time"

type Laporan struct {
	ID              int        `gorm:"type:int(11);primary_key;auto_increment"`
	KonstruksiID    int        `gorm:"type:int(11);not null"`
	Konstruksi      Konstruksi `gorm:"references:ID;foreignKey:KonstruksiID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	MandorID        int        `gorm:"type:int(11);not null"`
	Mandor          User       `gorm:"references:ID;foreignKey:MandorID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Title           string     `gorm:"type:varchar(100);not null"`
	Tipe            string     `gorm:"type:varchar(50);not null"`
	Client          string     `gorm:"type:varchar(50);not null"`
	MandorName      string     `gorm:"type:varchar(50);not null"`
	BahanKonstruksi *string    `gorm:"type:text;null"`
	Deskripsi       string     `gorm:"type:text; null"`
	Image           string     `gorm:"type:varchar(255);null"`
	CreatedAt       time.Time  `gorm:"type:timestamp;not null"`
	UpdatedAt       time.Time  `gorm:"type:timestamp;not null"`
}
