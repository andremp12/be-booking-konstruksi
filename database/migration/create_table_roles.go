package database

import (
	"time"
)

type Role struct {
	ID        int       `gorm:"type:int(1);primary_key;auto_increment"`
	Name      string    `gorm:"type:varchar(10);not null"`
	CreatedAt time.Time `gorm:"type:timestamp;not null"`
	UpdatedAt time.Time `gorm:"type:timestamp;not null"`
}
