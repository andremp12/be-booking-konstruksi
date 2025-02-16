package request

type Satuan struct {
	Name string `json:"name" form:"name" binding:"required"`
	Keterangan string `json:"keterangan" form:"keterangan"`
}
