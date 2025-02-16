package service

import (
	database "booking-konstruksi/database/migration"
	"booking-konstruksi/repository"
	"booking-konstruksi/request"
	"booking-konstruksi/response"
	"fmt"
	"time"
)

type Konstruksi interface {
	GetCountStatus() ([]response.CountStatusKonstruksi, error)
	GetRiwayatKonstruksi() ([]response.Konstruksi, error)
	GetRiwayatKonstruksiMandor(clientId string) ([]response.Konstruksi, error)
	GetRiwayatKonstruksiClient(mandorId string) ([]response.Konstruksi, error)
	GetAllData() ([]response.Konstruksi, error)
	GetData(id string) (response.Konstruksi, error)
	GetKonstruksiUser(clientId string) ([]response.Konstruksi, error)
	GetKonstruksiMandor(mandorId string) ([]response.Konstruksi, error)
	Booking(request request.Konstruksi) (response.Konstruksi, error)
	UpdateStatus(id string, status string) error
	KonfirmasiBooking(request request.Confirmation, id string) (response.Konstruksi, error)
	Delete(id string) error
}

type serviceKonstruksi struct {
	repoKonstruksi repository.Konstruksi
}

func NewServiceKonstruksi(repoKonstruksi repository.Konstruksi) *serviceKonstruksi {
	return &serviceKonstruksi{repoKonstruksi: repoKonstruksi}
}

func (s *serviceKonstruksi) GetCountStatus() ([]response.CountStatusKonstruksi, error) {
	response, err := s.repoKonstruksi.GetCountStatus()

	return response, err
}

func (s *serviceKonstruksi) GetRiwayatKonstruksiMandor(mandorId string) ([]response.Konstruksi, error) {
	var responseKonstruksis []response.Konstruksi
	konstruksis, err := s.repoKonstruksi.GetRiwayatKonstruksiMandor(mandorId)

	for _, k := range konstruksis {
		responseKonstruksi_ := response.ResponseKonstruksi(&k)
		responseKonstruksis = append(responseKonstruksis, *responseKonstruksi_)
	}

	return responseKonstruksis, err
}

func (s *serviceKonstruksi) GetRiwayatKonstruksiClient(clientId string) ([]response.Konstruksi, error) {
	var responseKonstruksis []response.Konstruksi
	konstruksis, err := s.repoKonstruksi.GetRiwayatKonstruksiClient(clientId)

	for _, k := range konstruksis {
		responseKonstruksi_ := response.ResponseKonstruksi(&k)
		responseKonstruksis = append(responseKonstruksis, *responseKonstruksi_)
	}

	return responseKonstruksis, err
}

func (s *serviceKonstruksi) GetRiwayatKonstruksi() ([]response.Konstruksi, error) {
	var responseKonstruksis []response.Konstruksi
	konstruksis, err := s.repoKonstruksi.GetRiwayatKonstruksi()

	for _, k := range konstruksis {
		responseKonstruksi_ := response.ResponseKonstruksi(&k)
		responseKonstruksis = append(responseKonstruksis, *responseKonstruksi_)
	}

	return responseKonstruksis, err
}

func (s *serviceKonstruksi) GetAllData() ([]response.Konstruksi, error) {
	var konstruksis []database.Konstruksi
	var responseKonstruksis []response.Konstruksi

	konstruksis, err := s.repoKonstruksi.GetAllData()

	for _, k := range konstruksis {
		responseKonstruksi_ := response.ResponseKonstruksi(&k)
		responseKonstruksis = append(responseKonstruksis, *responseKonstruksi_)
	}

	return responseKonstruksis, err
}

func (s *serviceKonstruksi) GetData(id string) (response.Konstruksi, error) {
	var konstruksi database.Konstruksi

	konstruksi, err := s.repoKonstruksi.GetData(id)
	responseKonstruksi := response.ResponseKonstruksi(&konstruksi)

	return *responseKonstruksi, err
}

func (s *serviceKonstruksi) GetKonstruksiUser(clientId string) ([]response.Konstruksi, error) {
	var konstruksis []database.Konstruksi
	var responseKonstruksis []response.Konstruksi

	konstruksis, err := s.repoKonstruksi.GetKonstruksiUser(clientId)

	for _, k := range konstruksis {
		responseKonstruksi_ := response.ResponseKonstruksi(&k)
		responseKonstruksis = append(responseKonstruksis, *responseKonstruksi_)
	}

	return responseKonstruksis, err
}

func (s *serviceKonstruksi) GetKonstruksiMandor(mandorId string) ([]response.Konstruksi, error) {
	var konstruksis []database.Konstruksi
	var responseKonstruksis []response.Konstruksi

	konstruksis, err := s.repoKonstruksi.GetKonstruksiMandor(mandorId)

	for _, k := range konstruksis {
		responseKonstruksi_ := response.ResponseKonstruksi(&k)
		responseKonstruksis = append(responseKonstruksis, *responseKonstruksi_)
	}

	return responseKonstruksis, err
}

func (s *serviceKonstruksi) Booking(request request.Konstruksi) (response.Konstruksi, error) {
	var noBooking string
	total, _ := request.Total.Int64()
	remainingPayment := int(total)
	//start_date, _ := time.Parse("2006-01-02", *request.StartDate)
	//end_date, _ := time.Parse("2006-01-02", *request.EndDate)

	// create format no booking
	lastKonstruksi, err := s.repoKonstruksi.GetLastData()
	if lastKonstruksi.ID == 0 {
		noBooking = fmt.Sprintf("BK%06d", 1)
	} else {
		noBooking = fmt.Sprintf("BK%06d", lastKonstruksi.ID+1)
	}

	konstruksi := database.Konstruksi{
		ClientID:         request.ClientID,
		NoBooking:        noBooking,
		TipeKonstruksiID: request.TipeKonstruksiID,
		Tipe:             database.Tipe(request.Tipe),
		Total:            int(total),
		RemainingPayment: &remainingPayment,
		Paid:             0,
		Address:          request.Address,
		Keterangan:       request.Keterangan,
		Status:           database.Status(request.Status),
		//StartDate:        &start_date,
		//EndDate:          &end_date,
	}

	konstruksi, err = s.repoKonstruksi.Booking(konstruksi)
	responseKonstruksi := response.ResponseKonstruksi(&konstruksi)

	return *responseKonstruksi, err
}

func (s *serviceKonstruksi) UpdateStatus(id string, status string) error {
	err := s.repoKonstruksi.UpdateStatus(id, status)

	return err
}

func (s *serviceKonstruksi) KonfirmasiBooking(request request.Confirmation, id string) (response.Konstruksi, error) {
	//var err error
	//
	start_date, _ := time.Parse("2006-01-02", *request.StartDate)
	end_date, _ := time.Parse("2006-01-02", *request.EndDate)

	konstruksi := database.Konstruksi{
		MandorID:  request.MandorID,
		Status:    "Payment",
		StartDate: &start_date,
		EndDate:   &end_date,
	}

	konstruksi, err := s.repoKonstruksi.KonfirmasiBooking(konstruksi, id)

	// Response konstruksi
	konstruksiResponse := response.ResponseKonstruksi(&konstruksi)
	konstruksiResponse.MandorID = konstruksi.MandorID

	return *konstruksiResponse, err
}

func (s *serviceKonstruksi) Delete(id string) error {
	err := s.repoKonstruksi.Delete(id)

	return err
}
