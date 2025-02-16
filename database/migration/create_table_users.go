package database

import "time"

type User struct {
	ID     int  `gorm:"type:int(11);primary_key;auto_increment"`
	RoleID int  `gorm:"type:int(1);not null"`
	Role   Role `gorm:"references:ID;foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	//Scopes    string    `gorm:"type:enum('admin','mandor','client');not null"`
	Name      string    `gorm:"type:varchar(50);not null"`
	Email     string    `gorm:"type:varchar(50);not null;unique"`
	NoWA      string    `gorm:"type:varchar(15);not null"`
	Password  string    `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"type:timestamp;not null"`
	UpdatedAt time.Time `gorm:"type:timestamp;not null"`
}
