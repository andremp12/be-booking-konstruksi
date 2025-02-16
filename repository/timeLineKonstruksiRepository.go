package repository

import (
	database "booking-konstruksi/database/migration"
	"gorm.io/gorm"
)

type TimeLineKonstruksi interface {
	GetTimeLineKonstruksi(konstruksiId string) ([]database.TimeLine, error)
	Create(konstruksi database.TimeLine) error
	Update(konstruksi database.TimeLine) error
	Delete(id string) error
}

type timeLineKonstruksiRepository struct {
	db *gorm.DB
}

func NewRepositoryTimelineKonstruksi(db *gorm.DB) *timeLineKonstruksiRepository {
	return &timeLineKonstruksiRepository{db: db}
}

func (r *timeLineKonstruksiRepository) GetTimeLineKonstruksi(kosntruksiId string) ([]database.TimeLine, error) {
	var timeLine []database.TimeLine

	err := r.db.Debug().Order("id desc").Limit(5).Find(&timeLine, "konstruksi_id = ?", kosntruksiId).Error

	return timeLine, err
}

func (r *timeLineKonstruksiRepository) Create(timeLine database.TimeLine) error {

	err := r.db.Create(&timeLine).Error
	return err
}

func (r *timeLineKonstruksiRepository) Update(timeLine database.TimeLine) error {

	err := r.db.Save(&timeLine).Error
	return err
}

func (r *timeLineKonstruksiRepository) Delete(id string) error {
	var timeLine database.TimeLine

	err := r.db.Where("id = ?", id).Delete(&timeLine).Error

	return err
}
