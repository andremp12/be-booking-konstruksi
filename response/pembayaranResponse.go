package response

import (
	database "booking-konstruksi/database/migration"
	"time"
)

type Pembayaran struct {
	ID           int         `json:"id"`
	KonstruksiID int         `json:"konstruksi_id"`
	Konstruksi   *Konstruksi `json:"konstruksi"`
	ClientID     int         `json:"client_id"`
	Client       *User       `json:"client"`
	Name         string      `json:"name"`
	Kode         string      `json:"kode"`
	Status       string      `json:"status"`
	Total        int         `json:"total"`
	DueDate      *time.Time  `json:"due_date"`
	PaymentDate  *time.Time  `json:"payment_date"`
	TotalPaid    TotalPaid   `json:"total_paid"`
}

type TotalPaid struct {
	KonstruksiID int   `json:"konstruksi_id"`
	Total        int64 `json:"total"`
}

func ResponsePembayaran(pembayaran *database.Pembayaran) Pembayaran {

	return Pembayaran{
		ID:           pembayaran.ID,
		KonstruksiID: pembayaran.KonstruksiID,
		Konstruksi:   ResponseKonstruksi(&pembayaran.Konstruksi),
		ClientID:     pembayaran.ClientID,
		Client:       UserResponse(&pembayaran.Client),
		Name:         pembayaran.Name,
		Status:       pembayaran.Status,
		Total:        pembayaran.Total,
		Kode:         pembayaran.Kode,
		DueDate:      pembayaran.DueDate,
		PaymentDate:  pembayaran.PaymentDate,
	}
}
