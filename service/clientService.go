package service

import (
	database "booking-konstruksi/database/migration"
	"booking-konstruksi/repository"
	"booking-konstruksi/response"
)

type Client interface {
	GetAllData() ([]response.User, error)
}

type serviceClient struct {
	repoClient repository.Client
}

func NewServiceClient(repoClient repository.Client) *serviceClient {
	return &serviceClient{repoClient: repoClient}
}

func (s *serviceClient) GetAllData() ([]response.User, error) {
	var client []database.User
	var responseClient []response.User

	client, err := s.repoClient.GetAllData()

	for _, m := range client {
		responseClient_ := response.UserResponse(&m)
		responseClient = append(responseClient, *responseClient_)
	}

	return responseClient, err
}
