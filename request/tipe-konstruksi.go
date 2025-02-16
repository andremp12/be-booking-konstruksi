package request

import (
	"encoding/json"
	"mime/multipart"
)

type TipeKonstruksi struct {
	Name      string                `json:"name" form:"name" binding:"required"`
	HargaFull json.Number           `json:"harga_full" form:"harga_full" binding:"required,number,gt=0"`
	HargaJasa json.Number           `json:"harga_jasa" form:"harga_jasa" binding:"required,number,gt=0"`
	FileImage *multipart.FileHeader `json:"file_image" form:"file_image""`
	Image     string                `json:"image" form:"image"`
}
