package controller

import (
	"booking-konstruksi/request"
	"booking-konstruksi/response"
	"booking-konstruksi/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"strconv"
)

type statusRequest struct {
	Status string `json:"status" binding:"required"`
}

type konstruksiController struct {
	serviceKonstruksi service.Konstruksi
}

func NewKonstruksiController(serviceKonstruksi service.Konstruksi) *konstruksiController {
	return &konstruksiController{serviceKonstruksi: serviceKonstruksi}
}

func (konstruksiController *konstruksiController) GetRiwayatKonstruksi(c *gin.Context) {
	konstruksis, err := konstruksiController.serviceKonstruksi.GetRiwayatKonstruksi()

	if err != nil {
		c.JSON(400, response.APIResponse{
			Status:  "error",
			Message: "Failed Get All Data",
			Data:    err.Error(),
		})

		return
	}

	c.JSON(200, response.APIResponse{
		Status:  "success",
		Message: "Success Get All Data",
		Data:    konstruksis,
	})
}

func (konstruksiController *konstruksiController) GetRiwayatKonstruksiMandor(c *gin.Context) {
	mandorId := c.Query("mandor_id")
	konstruksis, err := konstruksiController.serviceKonstruksi.GetRiwayatKonstruksiMandor(mandorId)

	if err != nil {
		c.JSON(400, response.APIResponse{
			Status:  "error",
			Message: "Failed Get All Data",
			Data:    err.Error(),
		})

		return
	}

	c.JSON(200, response.APIResponse{
		Status:  "success",
		Message: "Success Get All Data",
		Data:    konstruksis,
	})
}

func (konstruksiController *konstruksiController) GetRiwayatKonstruksiClient(c *gin.Context) {
	clientId := c.Query("client_id")
	konstruksis, err := konstruksiController.serviceKonstruksi.GetRiwayatKonstruksiClient(clientId)

	if err != nil {
		c.JSON(400, response.APIResponse{
			Status:  "error",
			Message: "Failed Get All Data",
			Data:    err.Error(),
		})

		return
	}

	c.JSON(200, response.APIResponse{
		Status:  "success",
		Message: "Success Get All Data",
		Data:    konstruksis,
	})
}

func (konstruksiController *konstruksiController) GetCountStatus(c *gin.Context) {
	konstruksi, err := konstruksiController.serviceKonstruksi.GetCountStatus()

	if err != nil {
		c.JSON(400, response.APIResponse{
			Status:  "error",
			Message: "Failed Get All Data",
			Data:    err.Error(),
		})

		return
	}

	c.JSON(200, response.APIResponse{
		Status:  "success",
		Message: "Success Get All Data",
		Data:    konstruksi,
	})
}

func (konstruksiController *konstruksiController) GetAllData(c *gin.Context) {
	konstruksis, err := konstruksiController.serviceKonstruksi.GetAllData()

	if err != nil {
		c.JSON(400, response.APIResponse{
			Status:  "error",
			Message: "Failed Get All Data",
			Data:    err.Error(),
		})

		return
	}

	c.JSON(200, response.APIResponse{
		Status:  "success",
		Message: "Success Get All Data",
		Data:    konstruksis,
	})
}

func (konstruksiController *konstruksiController) GetData(c *gin.Context) {
	id := c.Param("id")

	konstruksis, err := konstruksiController.serviceKonstruksi.GetData(id)
	if err != nil {
		c.JSON(400, response.APIResponse{
			Status:  "error",
			Message: "Failed Get Data",
			Data:    err.Error(),
		})

		return
	}

	c.JSON(200, response.APIResponse{
		Status:  "success",
		Message: "Success Get Tipe Konstruksi",
		Data:    konstruksis,
	})
}

func (konstruksiController *konstruksiController) GetKonstruksiUser(c *gin.Context) {
	clientId := c.Query("client_id")

	konstruksis, err := konstruksiController.serviceKonstruksi.GetKonstruksiUser(clientId)
	if err != nil {
		c.JSON(400, response.APIResponse{
			Status:  "error",
			Message: "Failed Get Data",
			Data:    err.Error(),
		})

		return
	}

	c.JSON(200, response.APIResponse{
		Status:  "success",
		Message: "Success Get Tipe Konstruksi",
		Data:    konstruksis,
	})
}

func (konstruksiController *konstruksiController) GetKonstruksiMandor(c *gin.Context) {
	mandorId := c.Query("mandor_id")

	konstruksis, err := konstruksiController.serviceKonstruksi.GetKonstruksiMandor(mandorId)
	if err != nil {
		c.JSON(400, response.APIResponse{
			Status:  "error",
			Message: "Failed Get Data",
			Data:    err.Error(),
		})

		return
	}

	c.JSON(200, response.APIResponse{
		Status:  "success",
		Message: "Success Get Data",
		Data:    konstruksis,
	})
}

func (konstruksiController *konstruksiController) Booking(c *gin.Context) {
	var request request.Konstruksi

	err := c.ShouldBind(&request)
	if err != nil {
		errMessages := []string{}
		log.Printf("Error: %T - %v", err, err)

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

	clientId, _ := strconv.Atoi(c.Query("client_id"))

	request.TipeKonstruksiID, _ = strconv.Atoi(c.Param("tipe_id"))
	request.ClientID = &clientId
	request.Status = "Booking"

	responseKonstruksi, err := konstruksiController.serviceKonstruksi.Booking(request)
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
		Data:    responseKonstruksi,
	})
}

func (konstruksiController *konstruksiController) UpdateStatus(c *gin.Context) {
	var request statusRequest

	id := c.Param("id")

	err := c.ShouldBind(&request)
	if err != nil {
		errMessages := []string{}

		for _, e := range err.(validator.ValidationErrors) {
			errMessage := fmt.Sprintf("Error Field %s, is %s", e.Field(), e.Tag())
			errMessages = append(errMessages, errMessage)
		}

		c.JSON(400, response.APIResponse{
			Status:  "error",
			Message: "Failed Save Data",
			Data:    errMessages,
		})

		return
	}

	err = konstruksiController.serviceKonstruksi.UpdateStatus(id, request.Status)
	if err != nil {
		c.JSON(500, response.APIResponse{
			Status:  "error",
			Message: "Failed Update Status",
			Data:    err.Error(),
		})

		return
	}

	c.JSON(200, response.APIResponse{
		Status:  "success",
		Message: "Success Update Status",
		Data:    nil,
	})
}

func (konstruksiController *konstruksiController) KonfirmasiBooking(c *gin.Context) {
	id := c.Param("id")

	var request request.Confirmation

	err := c.ShouldBind(&request)
	if err != nil {
		errMessages := []string{}

		for _, e := range err.(validator.ValidationErrors) {
			errMessage := fmt.Sprintf("Error Field %s, is %s", e.Field(), e.Tag())
			errMessages = append(errMessages, errMessage)
		}
		c.JSON(400, response.APIResponse{
			Status:  "error",
			Message: "Failed Save Data",
			Data:    errMessages,
		})

		return
	}

	konstruksiResponse, err := konstruksiController.serviceKonstruksi.KonfirmasiBooking(request, id)
	if err != nil {
		c.JSON(500, response.APIResponse{
			Status:  "error",
			Message: "Failed Save Data",
			Data:    err.Error(),
		})

		return
	}

	c.JSON(200, response.APIResponse{
		Status:  "success",
		Message: "Success Konfirmasi Booking",
		Data: gin.H{
			"record": konstruksiResponse,
		},
	})
}

func (konstruksiController *konstruksiController) Delete(c *gin.Context) {
	id := c.Param("id")

	err := konstruksiController.serviceKonstruksi.Delete(id)

	if err != nil {
		c.JSON(500, response.APIResponse{
			Status:  "error",
			Message: "Failed Delete Tipe Konstruksi",
			Data:    err.Error(),
		})

		return
	}

	c.JSON(200, response.APIResponse{
		Status:  "success",
		Message: "Success Delete Tipe Konstruksi",
	})
}
