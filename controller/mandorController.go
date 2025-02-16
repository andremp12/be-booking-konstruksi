package controller

import (
	"booking-konstruksi/response"
	"booking-konstruksi/service"
	"github.com/gin-gonic/gin"
)

type mandorController struct {
	serviceMandor service.Mandor
}

func NewMandorController(serviceMandor service.Mandor) *mandorController {
	return &mandorController{serviceMandor: serviceMandor}
}

func (mandorController *mandorController) GetAllData(c *gin.Context) {
	mandors, err := mandorController.serviceMandor.GetAllData()

	if err != nil {
		c.JSON(400, response.APIResponse{
			Status:  "error",
			Message: "Failed Get All Data",
			Data:    err.Error(),
		})

		return
	}

	// response := gin.H{
	// 	"records":Satuans,
	// }

	c.JSON(200, response.APIResponse{
		Status:  "success",
		Message: "Success Get All Data",
		Data:    mandors,
	})
}
