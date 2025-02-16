package service

import (
	database "booking-konstruksi/database/migration"
	"booking-konstruksi/repository"
	"booking-konstruksi/request"
	"booking-konstruksi/response"
)

type Satuan interface {
	GetAllData() ([]response.Satuan, error)
	GetData(id string) (response.Satuan, error)
	Create(request request.Satuan) (response.Satuan, error)
	Update(request request.Satuan, id string) (response.Satuan, error)
	Delete(id string) error
}

type serviceSatuan struct {
	repoSatuan repository.Satuan
}

func NewServiceSatuan(repoSatuan repository.Satuan) *serviceSatuan {
	return &serviceSatuan{repoSatuan: repoSatuan}
}

func (s *serviceSatuan) GetAllData() ([]response.Satuan, error) {
	var Satuans []response.Satuan

	dataSatuan, err := s.repoSatuan.GetAllData()

	for _, data := range dataSatuan {
		responseTipe := response.SatuanResponse(&data)
		Satuans = append(Satuans, responseTipe)
	}

	return Satuans, err
}

func (s *serviceSatuan) GetData(id string) (response.Satuan, error) {
	dataSatuan, err := s.repoSatuan.GetData(id)

	Satuan := response.SatuanResponse(&dataSatuan)

	return Satuan, err
}

func (s *serviceSatuan) Create(request request.Satuan) (response.Satuan, error) {

	Satuan := database.Satuan{
		Name: request.Name,
		Keterangan: request.Keterangan,
	}

	Satuan_, err := s.repoSatuan.Create(Satuan)
	responseTipe := response.SatuanResponse(&Satuan_)

	return responseTipe, err
}

func (s *serviceSatuan) Update(request request.Satuan, id string) (response.Satuan, error) {
	Satuan, err := s.repoSatuan.GetData(id)

	//Hapus file yang lama

	Satuan.Name = request.Name
	Satuan.Keterangan=request.Keterangan

	Satuan_, err := s.repoSatuan.Update(Satuan)
	responseTipe := response.SatuanResponse(&Satuan_)

	return responseTipe, err
}

func (s *serviceSatuan) Delete(id string) error {
	Satuan, err := s.repoSatuan.GetData(id)

	err = s.repoSatuan.Delete(Satuan)

	return err
}
