package service

import (
	database "booking-konstruksi/database/migration"
	"booking-konstruksi/repository"
	"booking-konstruksi/request"
	"booking-konstruksi/response"
	"fmt"
	"strconv"
	"time"
)

type Pembayaran interface {
	GetRiwayatPembayaran() ([]response.Pembayaran, error)
	GetRiwayatPembayaranClient(clientId string) ([]response.Pembayaran, error)
	GetAllData() ([]response.Pembayaran, error)
	GetData(id string) (response.Pembayaran, error)
	GetPembayaranClient(konstruksiId string) ([]response.Pembayaran, error)
	Create(request request.Pembayaran) (response.Pembayaran, error)
	Update(id string, Pembayaran request.Pembayaran) (response.Pembayaran, error)
	Delete(id string) error
	SuccessPayment(status string, id string) error
	GetTotalPaid() ([]response.TotalPaid, error)
	GetTotalPaidKonstruksi(konstruksi_id string) (response.TotalPaid, error)
}

type servicePembayaran struct {
	repoPembayaran repository.Pembayaran
	repoKonstruksi repository.Konstruksi
}

func NewServicePembayaran(repoPembayaran repository.Pembayaran) *servicePembayaran {
	return &servicePembayaran{repoPembayaran: repoPembayaran}
}

func (s *servicePembayaran) GetRiwayatPembayaran() ([]response.Pembayaran, error) {
	var pembayarans []database.Pembayaran

	pembayarans, err := s.repoPembayaran.GetRiwayatPembayaran()

	responsePembayarans := []response.Pembayaran{}
	for _, pembayaran := range pembayarans {
		responsePembayaran_ := response.ResponsePembayaran(&pembayaran)
		responsePembayarans = append(responsePembayarans, responsePembayaran_)
	}

	return responsePembayarans, err
}

func (s *servicePembayaran) GetRiwayatPembayaranClient(clientId string) ([]response.Pembayaran, error) {
	var pembayarans []database.Pembayaran

	pembayarans, err := s.repoPembayaran.GetRiwayatPembayaranClient(clientId)

	responsePembayarans := []response.Pembayaran{}
	for _, pembayaran := range pembayarans {
		responsePembayaran_ := response.ResponsePembayaran(&pembayaran)
		responsePembayarans = append(responsePembayarans, responsePembayaran_)
	}

	return responsePembayarans, err
}

func (s *servicePembayaran) GetAllData() ([]response.Pembayaran, error) {
	var pembayarans []database.Pembayaran

	pembayarans, err := s.repoPembayaran.GetAllData()

	responsePembayarans := []response.Pembayaran{}
	for _, pembayaran := range pembayarans {
		responsePembayaran_ := response.ResponsePembayaran(&pembayaran)
		responsePembayarans = append(responsePembayarans, responsePembayaran_)
	}

	return responsePembayarans, err
}

func (s *servicePembayaran) GetData(id string) (response.Pembayaran, error) {
	var pembayaran database.Pembayaran

	pembayaran, err := s.repoPembayaran.GetData(id)
	responsePembayaran := response.ResponsePembayaran(&pembayaran)

	return responsePembayaran, err
}

func (s *servicePembayaran) GetPembayaranClient(konstruksiId string) ([]response.Pembayaran, error) {
	pembayarans, err := s.repoPembayaran.GetPembayaranClient(konstruksiId)

	responsePembayarans := []response.Pembayaran{}
	for _, pembayaran := range pembayarans {
		responsePembayaran_ := response.ResponsePembayaran(&pembayaran)
		responsePembayarans = append(responsePembayarans, responsePembayaran_)
	}

	return responsePembayarans, err
}

func (s *servicePembayaran) Create(request request.Pembayaran) (response.Pembayaran, error) {
	var kode string

	konstruksiID, _ := strconv.Atoi(request.KonstruksiID)
	clientID, _ := strconv.Atoi(request.ClientID)
	dueDate, _ := time.Parse("2006-01-02", *request.DueDate)
	total, _ := request.Total.Int64()
	//paymentDate, _ := time.Parse("2006-01-02", *request.PaymentDate)

	// Create payment code
	lastPembayaran, _ := s.repoPembayaran.GetLastData()

	if lastPembayaran.ID == 0 {
		kode = fmt.Sprintf("TRX%d%d%04d", konstruksiID, clientID, 1)
	} else {
		kode = fmt.Sprintf("TRX%d%d%04d", konstruksiID, clientID, lastPembayaran.ID+1)
	}

	pembayaran := database.Pembayaran{
		KonstruksiID: konstruksiID,
		ClientID:     clientID,
		Name:         request.Name,
		Status:       "Unpaid",
		Kode:         kode,
		DueDate:      &dueDate,
		Total:        int(total),
		//PaymentDate:  &paymentDate,
	}

	pembayaran, err := s.repoPembayaran.Create(pembayaran)
	pembayaranResponse := response.ResponsePembayaran(&pembayaran)

	pembayaranResponse.Konstruksi = nil
	pembayaranResponse.Client = nil

	return pembayaranResponse, err
}

func (s *servicePembayaran) Delete(id string) error {
	err := s.repoPembayaran.Delete(id)

	return err
}

func (s *servicePembayaran) Update(id string, request request.Pembayaran) (response.Pembayaran, error) {
	dueDate, _ := time.Parse("2006-01-02", *request.DueDate)
	total, _ := request.Total.Int64()

	pembayaran, err := s.repoPembayaran.GetData(id)
	pembayaran.DueDate = &dueDate
	pembayaran.Total = int(total)
	pembayaran.Name = request.Name

	pembayaran, err = s.repoPembayaran.Update(pembayaran)
	responsePembayaran := response.ResponsePembayaran(&pembayaran)

	return responsePembayaran, err
}

func (s *servicePembayaran) SuccessPayment(status string, id string) error {

	err := s.repoPembayaran.SuccessPayment(status, id)

	//konstruksi_id := strconv.Itoa(pembayaran.KonstruksiID)
	//err = s.repoPembayaran.UpdatePaidKonstruksi(konstruksi_id, pembayaran.Total)

	return err
}

func (s *servicePembayaran) GetTotalPaid() ([]response.TotalPaid, error) {
	totalPaid, err := s.repoPembayaran.GetTotalPaid()

	return totalPaid, err
}

func (s *servicePembayaran) GetTotalPaidKonstruksi(konstruksi_id string) (response.TotalPaid, error) {
	totalPaid, err := s.repoPembayaran.GetTotalPaidKonstruksi(konstruksi_id)

	return totalPaid, err
}
