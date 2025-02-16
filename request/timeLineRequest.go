package request

type TimeLine struct {
	KonstruksiID int     `json:"konstruksi_id"`
	Name         string  `json:"name"`
	Date         string  `json:"date"`
	Description  *string `json:"description"`
}
