package database

import "time"

type Pembayaran struct {
	ID           int        `gorm:"type:int(11);primary_key;auto_increment"`
	KonstruksiID int        `gorm:"type:int(11);null"`
	Konstruksi   Konstruksi `gorm:"references:ID;foreignKey:KonstruksiID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ClientID     int        `gorm:"type:int(11);not null"`
	Client       User       `gorm:"references:ID;foreignKey:ClientID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Metode       *string    `gorm:"type:varchar(100);null"`
	Name         string     `gorm:"type:varchar(10);not null"`
	Kode         string     `gorm:"type:varchar(50);not null;unique"`
	Status       string     `gorm:"type:ENUM('Pending', 'Paid', 'Failed','Unpaid');DEFAULT('Waiting');NOT NULL"`
	Total        int        `gorm:"type:bigInt;not null"`
	Token        string     `gorm:"type:varchar(255);null"`
	DueDate      *time.Time `gorm:"type:datetime;not null"`
	PaymentDate  *time.Time `gorm:"type:datetime;null"`
	CreatedAt    time.Time  `gorm:"type:timestamp;not null"`
	UpdatedAt    time.Time  `gorm:"type:timestamp;not null"`
}
