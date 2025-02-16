package service

import (
	database "booking-konstruksi/database/migration"
	"booking-konstruksi/repository"
	"booking-konstruksi/response"
)

type Mandor interface {
	GetAllData() ([]response.User, error)
}

type serviceMandor struct {
	repoMandor repository.Mandor
}

func NewServiceMandor(repoMandor repository.Mandor) *serviceMandor {
	return &serviceMandor{repoMandor: repoMandor}
}

func (s *serviceMandor) GetAllData() ([]response.User, error) {
	var mandor []database.User
	var responseMandor []response.User

	mandor, err := s.repoMandor.GetAllData()

	for _, m := range mandor {
		responseMandor_ := response.UserResponse(&m)
		responseMandor = append(responseMandor, *responseMandor_)
	}

	return responseMandor, err
}
