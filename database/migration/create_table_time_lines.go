package database

import "time"

type TimeLine struct {
	ID           int        `gorm:"type:int(1);primary_key;auto_increment"`
	KonstruksiID int        `gorm:"type:int(11);not null"`
	Konstruksi   Konstruksi `gorm:"references:ID;foreignKey:KonstruksiID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name         string     `gorm:"type:varchar(50);not null"`
	Date         *time.Time `gorm:"type:datetime;not null"`
	Description  *string    `gorm:"type:text;not null"`
	CreatedAt    time.Time  `gorm:"type:timestamp;not null"`
	UpdatedAt    time.Time  `gorm:"type:timestamp;not null"`
}
