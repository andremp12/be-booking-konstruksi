package service

import (
	database "booking-konstruksi/database/migration"
	"booking-konstruksi/repository"
	"booking-konstruksi/request"
	"booking-konstruksi/response"
	"strconv"
)

type LaporanKonstruksi interface {
	GetLaporanKonstruksi(konstruksiId string) ([]response.Laporan, error)
	//GetActivityKonstruksi(mandorId string) ([]response.Laporan, error)
	Create(request request.Laporan) error
	//Update(konstruksi database.Laporan) error
	Delete(id string) error
}

type serviceLaporanKonstruksi struct {
	repoLaporan repository.LaporanKonstruksi
}

func NewServiceLaporanKonstruksi(repoLaporan repository.LaporanKonstruksi) *serviceLaporanKonstruksi {
	return &serviceLaporanKonstruksi{repoLaporan: repoLaporan}
}

func (s *serviceLaporanKonstruksi) GetLaporanKonstruksi(konstruksiId string) ([]response.Laporan, error) {
	var responseLaporans []response.Laporan
	//var bahanKonstruksi map[string]interface{}

	laporans, err := s.repoLaporan.GetLaporanKonstruksi(konstruksiId)
	for _, laporan := range laporans {
		responseLaporan_ := response.LaporanResponse(&laporan)
		//err = json.Unmarshal([]byte(*laporan.BahanKonstruksi), &bahanKonstruksi)
		//responseLaporan_.BahanKonstruksi = bahanKonstruksi
		responseLaporans = append(responseLaporans, responseLaporan_)
	}

	return responseLaporans, err
}

func (s *serviceLaporanKonstruksi) Create(request request.Laporan) error {
	mandorId, err := strconv.Atoi(request.MandrorID)
	//bahanKonstruksi, err := json.Marshal(request.BahanKonstruksi)
	//bahanKonstruksi_ := string(bahanKonstruksi)

	laporan := database.Laporan{
		KonstruksiID:    request.KonstruksiID,
		MandorID:        mandorId,
		Title:           request.Title,
		Tipe:            request.Tipe,
		Client:          request.Client,
		MandorName:      request.MandorName,
		BahanKonstruksi: &request.BahanKonstruksi,
		Image:           request.Image,
	}

	err = s.repoLaporan.Create(laporan, request.Status)

	return err
}

func (s *serviceLaporanKonstruksi) Delete(id string) error {
	err := s.repoLaporan.Delete(id)

	return err
}
