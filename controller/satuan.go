package controller

import (
	"booking-konstruksi/request"
	"booking-konstruksi/response"
	"booking-konstruksi/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type satuanController struct {
	serviceSatuan service.Satuan
}

func NewSatuanController(serviceSatuan service.Satuan) *satuanController {
	return &satuanController{serviceSatuan: serviceSatuan}
}

func (satuanController *satuanController) GetAllData(c *gin.Context) {
	Satuans, err := satuanController.serviceSatuan.GetAllData()

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
		Data:  Satuans,
	})
}

func (satuanController *satuanController) GetData(c *gin.Context) {
	id := c.Param("id")

	Satuan, err := satuanController.serviceSatuan.GetData(id)
	if err != nil {
		c.JSON(400, response.APIResponse{
			Status:  "error",
			Message: "Failed Get All Tipe Konstruksi",
			Data:    err.Error(),
		})

		return
	}

	c.JSON(200, response.APIResponse{
		Status:  "success",
		Message: "Success Get Data",
		Data:    Satuan,
	})
}

func (satuanController *satuanController) CreateData(c *gin.Context) {
	var requestTipe request.Satuan

	err := c.ShouldBind(&requestTipe)
	if err != nil {
		errMessages := []string{}

		for _, e := range err.(validator.ValidationErrors) {
			errMessage := fmt.Sprintf("Error Field %s, is %s", e.Field(), e.Tag())
			errMessages = append(errMessages, errMessage)
		}

		c.JSON(400, response.APIResponse{
			Status:  "error",
			Message: "Failed Saved data",
			Data:    errMessages,
		})

		return
	}

	responseSatuan, err := satuanController.serviceSatuan.Create(requestTipe)
	if err != nil {
		c.JSON(500, response.APIResponse{
			Status:  "error",
			Message: "Failed Save File",
			Data:    err.Error(),
		})

		return
	}

	c.JSON(200, response.APIResponse{
		Status:  "success",
		Message: "Success Create Data",
		Data:    responseSatuan,
	})
}

func (satuanController *satuanController) UpdateData(c *gin.Context) {
	var requestTipe request.Satuan

	id := c.Param("id")

	err := c.ShouldBind(&requestTipe)
	if err != nil {
		errMessages := []string{}

		for _, e := range err.(validator.ValidationErrors) {
			errMessage := fmt.Sprintf("Error Field %s, is %s", e.Field(), e.Tag())
			errMessages = append(errMessages, errMessage)
		}

		c.JSON(400, response.APIResponse{
			Status:  "error",
			Message: "Failed Saved data",
			Data:    errMessages,
		})

		return
	}

	responseSatuan, err := satuanController.serviceSatuan.Update(requestTipe, id)
	if err != nil {
		c.JSON(500, response.APIResponse{
			Status:  "error",
			Message: "Failed Save File",
			Data:    err.Error(),
		})

		return
	}

	c.JSON(200, response.APIResponse{
		Status:  "success",
		Message: "Success Update Data",
		Data:    responseSatuan,
	})
}

func (satuanController *satuanController) Delete(c *gin.Context) {
	id := c.Param("id")

	err := satuanController.serviceSatuan.Delete(id)

	if err != nil {
		c.JSON(500, response.APIResponse{
			Status:  "error",
			Message: "Failed Delete Data",
			Data:    err.Error(),
		})

		return
	}

	c.JSON(200, response.APIResponse{
		Status:  "success",
		Message: "Success Delete Data",
	})
}
