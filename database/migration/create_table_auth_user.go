package database

import "time"

type TypeRole string

const (
	Admin  TypeRole = "admin"
	Mandor TypeRole = "mandor"
	Client TypeRole = "client"
)

type AuthUser struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	UserID    int
	User      User     `gorm:"references:ID;foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Role      TypeRole `gorm:"type:ENUM('admin', 'mandor', 'client')"`
	Token     string   `gorm:"unique;not null"`
	CreatedAt time.Time
	ExpiresAt time.Time `gorm:"not null"`
}
