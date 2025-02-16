package database

import "time"

type Satuan struct {
	ID        int       `gorm:"type:int(1);primary_key;auto_increment"`
	Name      string    `gorm:"type:varchar(10);not null"`
	Keterangan string	`gorm:"type:text;null"`
	CreatedAt time.Time `gorm:"type:timestamp;not null"`
	UpdatedAt time.Time `gorm:"type:timestamp;not null"`
}
