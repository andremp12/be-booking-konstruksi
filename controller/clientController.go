package controller

import (
	"booking-konstruksi/response"
	"booking-konstruksi/service"
	"github.com/gin-gonic/gin"
)

type ClientController struct {
	serviceClient service.Client
}

func NewClientController(serviceClient service.Client) *ClientController {
	return &ClientController{serviceClient: serviceClient}
}

func (ClientController *ClientController) GetAllData(c *gin.Context) {
	clients, err := ClientController.serviceClient.GetAllData()

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
		Data:    clients,
	})
}
