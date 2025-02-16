package repository

import (
	database "booking-konstruksi/database/migration"
	"gorm.io/gorm"
	"time"
)

type AuthenticationRepository interface {
	Login(AuthUser database.AuthUser) error
	Logout(token string) error
	ValidateToken(token string) (database.AuthUser, error)
	Register(User database.User) (database.User, error)
	CheckUser(email string) (database.User, error)
}

type authenticationRepository struct {
	db *gorm.DB
}

func NewAuthenticationRepository(db *gorm.DB) *authenticationRepository {
	return &authenticationRepository{db: db}
}

func (r *authenticationRepository) Login(AuthUser database.AuthUser) error {
	err := r.db.Debug().Create(&AuthUser).Error

	return err
}

func (r *authenticationRepository) Logout(token string) error {
	var authUser database.AuthUser

	err := r.db.Debug().Where("token = ?", token).Delete(&authUser).Error

	return err
}

func (r *authenticationRepository) ValidateToken(Token string) (database.AuthUser, error) {
	var AuthUser database.AuthUser

	err := r.db.Debug().Preload("User").Where("token = ? AND expires_at > ?", Token, time.Now()).First(&AuthUser).Error

	return AuthUser, err
}

func (r *authenticationRepository) Register(User database.User) (database.User, error) {
	err := r.db.Debug().Create(&User).Error

	return User, err
}

func (r *authenticationRepository) CheckUser(email string) (database.User, error) {
	var User database.User

	err := r.db.Debug().Preload("Role").Where("email = ?", email).First(&User).Error

	return User, err
}
