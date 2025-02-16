package request

import "mime/multipart"

type Laporan struct {
	KonstruksiID    int    `form:"konstruksi_id" binding:"required"`
	MandrorID       string `form:"mandor_id" binding:"required"`
	Title           string `form:"title" binding:"required"`
	Status          string `form:"status" binding:"required"`
	Tipe            string `form:"tipe" binding:"required"`
	Client          string `form:"client"   binding:"required"`
	MandorName      string `form:"mandor_name" binding:"required"`
	BahanKonstruksi string
	Deskripsi       string                `form:"deskripsi"`
	FileImage       *multipart.FileHeader `form:"file_image" binding:"required"`
	Image           string                `form:"image"`
}
