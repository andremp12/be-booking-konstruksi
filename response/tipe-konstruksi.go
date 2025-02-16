package response

import database "booking-konstruksi/database/migration"

type TipeKonstruksi struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	HargaFull int    `json:"harga_full"`
	HargaJasa int    `json:"harga_jasa"`
	Image     string `json:"image"`
}

func TipeKonstruksiResponse(data *database.TipeKonstruksi) TipeKonstruksi {
	return TipeKonstruksi{
		ID:        data.ID,
		Name:      data.Name,
		HargaFull: data.HargaFull,
		HargaJasa: data.HargaJasa,
		Image:     data.Image,
	}
}
