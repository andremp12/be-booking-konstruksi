package request

import "encoding/json"

type MidtransRequest struct {
	UserId   json.Number `json:"user_id"`
	Amount   json.Number `json:"amount"`
	ItemID   string      `json:"item_id"`
	ItemName string      `json:"item_name"`
}
