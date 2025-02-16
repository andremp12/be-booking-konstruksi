package response

import (
	database "booking-konstruksi/database/migration"
	"time"
)

type Laporan struct {
	ID              int                    `json:"id"`
	KonstruksiID    int                    `json:"konstruksi_id"`
	Konstruksi      *Konstruksi            `json:"konstruksi"`
	Title           string                 `json:"title"`
	Tipe            string                 `json:"tipe"`
	Client          string                 `json:"client"`
	MandorName      string                 `json:"mandor_name"`
	BahanKonstruksi map[string]interface{} `json:"bahan_konstruksi"`
	Deskripsi       string                 `json:"deskripsi"`
	Image           string                 `json:"image"`
	CreatedAt       time.Time              `json:"created_at"`
}

func LaporanResponse(laporan *database.Laporan) Laporan {
	return Laporan{
		ID:           laporan.ID,
		KonstruksiID: laporan.KonstruksiID,
		Konstruksi:   ResponseKonstruksi(&laporan.Konstruksi),
		Title:        laporan.Title,
		Tipe:         laporan.Tipe,
		Client:       laporan.Client,
		MandorName:   laporan.MandorName,
		Deskripsi:    laporan.Deskripsi,
		Image:        laporan.Image,
		CreatedAt:    laporan.CreatedAt,
	}
}
