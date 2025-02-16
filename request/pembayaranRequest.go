package request

import (
	"encoding/json"
)

type Pembayaran struct {
	ID           int          `json:"id"`
	KonstruksiID string       `json:"konstruksi_id"`
	ClientID     string       `json:"client_id"`
	Name         string       `json:"name" binding:"required"`
	Total        *json.Number `json:"total" binding:"required"`
	DueDate      *string      `json:"due_date" binding:"required"`
	//PaymentDate  *string      `json:"payment_date" binding:"required"`
}
