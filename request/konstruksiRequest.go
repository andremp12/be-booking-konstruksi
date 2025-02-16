package request

import (
	"encoding/json"
)

type Konstruksi struct {
	ID               int          `json:"id"`
	TipeKonstruksiID int          `json:"tipe_konstruksi_id"`
	ClientID         *int         `json:"client_id"`
	MandorID         *int         `json:"mandor_id"`
	Tipe             string       `json:"tipe" binding:"required"`
	Address          string       `json:"address" binding:"required"`
	Keterangan       string       `json:"keterangan"`
	Status           string       `json:"status"`
	Total            *json.Number `json:"total" binding:"required,number"`
	StartDate        *string      `json:"start_date"`
	EndDate          *string      `json:"end_date"`
}

type Confirmation struct {
	MandorID  *int    `json:"mandor_id" binding:"required"`
	StartDate *string `json:"start_date" binding:"required"`
	EndDate   *string `json:"end_date" binding:"required"`
	Status    string  `json:"status"`
}
