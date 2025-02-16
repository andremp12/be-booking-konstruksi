package response

import (
	database "booking-konstruksi/database/migration"
	"time"
)

type Konstruksi struct {
	ID               int            `json:"id"`
	NoBooking        string         `json:"no_booking"`
	TipeKonstruksiID int            `json:"tipe_konstruksi_id"`
	TipeKonstruksi   TipeKonstruksi `json:"tipe_konstruksi"`
	ClientID         *int           `json:"client_id"`
	Client           *User          `json:"client"`
	MandorID         *int           `json:"mandor_id"`
	Mandor           *User          `json:"mandor"`
	Tipe             string         `json:"tipe"`
	Address          string         `json:"address"`
	Keterangan       string         `json:"keterangan"`
	Status           string         `json:"status"`
	Total            int            `json:"total"`
	Paid             int            `json:"paid"`
	RemainingPayment *int           `json:"remaining_payment"`
	StartDate        *time.Time     `json:"start_date"`
	EndDate          *time.Time     `json:"end_date"`
	CreatedAt        time.Time      `json:"created_at"`
}

type CountStatusKonstruksi struct {
	Status string `json:"status"`
	Count  int    `json:"count"`
}

func ResponseKonstruksi(konstruksi *database.Konstruksi) *Konstruksi {

	return &Konstruksi{
		ID:               konstruksi.ID,
		NoBooking:        konstruksi.NoBooking,
		TipeKonstruksiID: konstruksi.TipeKonstruksiID,
		TipeKonstruksi:   TipeKonstruksiResponse(&konstruksi.TipeKonstruksi),
		ClientID:         konstruksi.ClientID,
		Client:           UserResponse(&konstruksi.Client),
		MandorID:         konstruksi.MandorID,
		Mandor:           UserResponse(&konstruksi.Mandor),
		Tipe:             string(konstruksi.Tipe),
		Address:          konstruksi.Address,
		Keterangan:       konstruksi.Keterangan,
		Status:           string(konstruksi.Status),
		Total:            konstruksi.Total,
		Paid:             konstruksi.Paid,
		RemainingPayment: konstruksi.RemainingPayment,
		StartDate:        konstruksi.StartDate,
		EndDate:          konstruksi.EndDate,
		CreatedAt:        konstruksi.CreatedAt,
	}
}
