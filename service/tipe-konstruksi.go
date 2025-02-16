package service

import (
	database "booking-konstruksi/database/migration"
	"booking-konstruksi/repository"
	"booking-konstruksi/request"
	"booking-konstruksi/response"
	"fmt"
	"os"
)

type TipeKonstruksi interface {
	GetTipeLanding() ([]response.TipeKonstruksi, error)
	GetAllData() ([]response.TipeKonstruksi, error)
	GetData(id string) (response.TipeKonstruksi, error)
	Create(request request.TipeKonstruksi) (response.TipeKonstruksi, error)
	Update(request request.TipeKonstruksi, id string) (response.TipeKonstruksi, error)
	Delete(id string) error
}

type serviceTipeKonstruksi struct {
	repoTipeKonstruksi repository.TipeKonstruksi
}

func NewServiceTipeKonstruksi(repoTipeKonstruksi repository.TipeKonstruksi) *serviceTipeKonstruksi {
	return &serviceTipeKonstruksi{repoTipeKonstruksi: repoTipeKonstruksi}
}

func (s *serviceTipeKonstruksi) GetTipeLanding() ([]response.TipeKonstruksi, error) {
	var tipeKonstruksis []response.TipeKonstruksi

	dataTipe, err := s.repoTipeKonstruksi.GetTipeLanding()

	for _, data := range dataTipe {
		responseTipe := response.TipeKonstruksiResponse(&data)
		tipeKonstruksis = append(tipeKonstruksis, responseTipe)
	}
	fmt.Println(tipeKonstruksis)

	return tipeKonstruksis, err
}

func (s *serviceTipeKonstruksi) GetAllData() ([]response.TipeKonstruksi, error) {
	var tipeKonstruksis []response.TipeKonstruksi

	dataTipe, err := s.repoTipeKonstruksi.GetAllData()

	for _, data := range dataTipe {
		responseTipe := response.TipeKonstruksiResponse(&data)
		tipeKonstruksis = append(tipeKonstruksis, responseTipe)
	}
	fmt.Println(tipeKonstruksis)

	return tipeKonstruksis, err
}

func (s *serviceTipeKonstruksi) GetData(id string) (response.TipeKonstruksi, error) {
	dataTipe, err := s.repoTipeKonstruksi.GetData(id)

	tipeKonstruksi := response.TipeKonstruksiResponse(&dataTipe)

	return tipeKonstruksi, err
}

func (s *serviceTipeKonstruksi) Create(request request.TipeKonstruksi) (response.TipeKonstruksi, error) {
	hargaFull, _ := request.HargaFull.Int64()
	hargaJasa, _ := request.HargaJasa.Int64()

	tipeKonstruksi := database.TipeKonstruksi{
		Name:      request.Name,
		HargaFull: int(hargaFull),
		HargaJasa: int(hargaJasa),
		Image:     request.Image,
	}

	tipekonstruksi_, err := s.repoTipeKonstruksi.Create(tipeKonstruksi)
	responseTipe := response.TipeKonstruksiResponse(&tipekonstruksi_)

	return responseTipe, err
}

func (s *serviceTipeKonstruksi) Update(request request.TipeKonstruksi, id string) (response.TipeKonstruksi, error) {
	tipeKonstruksi, err := s.repoTipeKonstruksi.GetData(id)

	//Hapus file yang lama
	_, err = os.Stat(request.Image)
	if !os.IsNotExist(err) {
		os.Remove(tipeKonstruksi.Image)
	}
	fmt.Println(tipeKonstruksi.Image)
	fmt.Println(err)

	hargaFull, _ := request.HargaFull.Int64()
	hargaJasa, _ := request.HargaJasa.Int64()

	tipeKonstruksi.Name = request.Name
	tipeKonstruksi.HargaFull = int(hargaFull)
	tipeKonstruksi.HargaJasa = int(hargaJasa)
	if request.FileImage != nil {
		tipeKonstruksi.Image = request.Image
	}
	tipekonstruksi_, err := s.repoTipeKonstruksi.Update(tipeKonstruksi)
	responseTipe := response.TipeKonstruksiResponse(&tipekonstruksi_)

	return responseTipe, err
}

func (s *serviceTipeKonstruksi) Delete(id string) error {
	tipeKonstruksi, err := s.repoTipeKonstruksi.GetData(id)

	//Memeriksa apakah file ada di dalam folder images
	_, err = os.Stat(tipeKonstruksi.Image)
	if os.IsNotExist(err) {
		err = s.repoTipeKonstruksi.Delete(tipeKonstruksi)
		return err
	}

	//tipeKonstruksi.Image : ./images/nameFile_dalam_database
	os.Remove(tipeKonstruksi.Image)
	err = s.repoTipeKonstruksi.Delete(tipeKonstruksi)

	return err
}
