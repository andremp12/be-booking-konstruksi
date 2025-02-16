package service

import (
	database "booking-konstruksi/database/migration"
	"booking-konstruksi/repository"
	"booking-konstruksi/request"
	"booking-konstruksi/response"
	"time"
)

type TimeLineKonstruksi interface {
	GetTimeLineKonstruksi(konstruksiId string) ([]response.TimeLine, error)
	Create(request request.TimeLine) error
	Update(request request.TimeLine) error
	Delete(id string) error
}

type timeLineKonstruksiService struct {
	repoTimeLine repository.TimeLineKonstruksi
}

func NewServiceTimelineKonstruksi(repoTimeLine repository.TimeLineKonstruksi) *timeLineKonstruksiService {
	return &timeLineKonstruksiService{repoTimeLine: repoTimeLine}
}

func (s *timeLineKonstruksiService) GetTimeLineKonstruksi(kosntruksiId string) ([]response.TimeLine, error) {
	var responseTimeLines []response.TimeLine

	timeLines, err := s.repoTimeLine.GetTimeLineKonstruksi(kosntruksiId)

	for _, timeLine := range timeLines {
		responseTimeLine_ := response.TimeLineResponse(&timeLine)
		responseTimeLines = append(responseTimeLines, responseTimeLine_)
	}

	return responseTimeLines, err
}

func (s *timeLineKonstruksiService) Create(request request.TimeLine) error {
	date, err := time.Parse("2006-01-02", request.Date)

	timeLine := database.TimeLine{
		KonstruksiID: request.KonstruksiID,
		Name:         request.Name,
		Description:  request.Description,
		Date:         &date,
	}

	err = s.repoTimeLine.Create(timeLine)
	return err
}

func (s *timeLineKonstruksiService) Update(request request.TimeLine) error {
	date, err := time.Parse("2006-01-02", request.Date)

	timeLine := database.TimeLine{
		KonstruksiID: request.KonstruksiID,
		Name:         request.Name,
		Description:  request.Description,
		Date:         &date,
	}

	err = s.repoTimeLine.Update(timeLine)
	return err
}

func (s *timeLineKonstruksiService) Delete(id string) error {
	err := s.repoTimeLine.Delete(id)

	return err
}
