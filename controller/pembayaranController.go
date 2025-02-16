package controller

import (
	"booking-konstruksi/request"
	"booking-konstruksi/response"
	"booking-konstruksi/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
)

type PembayaranController struct {
	servicePembayaran service.Pembayaran
}

func NewPembayaranController(servicePembayaran service.Pembayaran) *PembayaranController {
	return &PembayaranController{servicePembayaran: servicePembayaran}
}

func (pembayaranController *PembayaranController) GetRiwayatPembayaran(c *gin.Context) {
	pembayarans, err := pembayaranController.servicePembayaran.GetRiwayatPembayaran()
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
		Data: gin.H{
			"records": pembayarans,
		},
	})
}

func (pembayaranController *PembayaranController) GetRiwayatPembayaranClient(c *gin.Context) {
	clientId := c.Param("client_id")
	pembayarans, err := pembayaranController.servicePembayaran.GetRiwayatPembayaranClient(clientId)

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
		Data: gin.H{
			"records": pembayarans,
		},
	})
}

func (pembayaranController *PembayaranController) GetPembayaranClient(c *gin.Context) {
	konstruksiId := c.Param("konstruksi_id")

	pembayarans, err := pembayaranController.servicePembayaran.GetPembayaranClient(konstruksiId)
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
		Data: gin.H{
			"records": pembayarans,
		},
	})
}

func (pembayaranController *PembayaranController) Create(c *gin.Context) {
	var request request.Pembayaran

	konstruksiID := c.Param("konstruksi_id")
	clientId := c.Query("client_id")

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

	request.KonstruksiID = konstruksiID
	request.ClientID = clientId

	pembayaranResponse, err := pembayaranController.servicePembayaran.Create(request)
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
		Message: "Success Save Data",
		Data: gin.H{
			"record": pembayaranResponse,
		},
	})
}

func (pembayaranController *PembayaranController) SuccessPayment(c *gin.Context) {
	id := c.Param("id")

	err := pembayaranController.servicePembayaran.SuccessPayment("Paid", id)
	if err != nil {
		c.JSON(500, response.APIResponse{
			Status:  "error",
			Message: "Failed Update Payment Status",
		})

		return
	}

	c.JSON(200, response.APIResponse{
		Status:  "success",
		Message: "Success Payment",
	})
}

func (pembayaranController *PembayaranController) GetTotalPaid(c *gin.Context) {

	totalPaid, err := pembayaranController.servicePembayaran.GetTotalPaid()
	if err != nil {
		c.JSON(500, response.APIResponse{
			Status:  "error",
			Message: "Failed Get Total Paid",
		})

		return
	}

	c.JSON(200, response.APIResponse{
		Status:  "success",
		Message: "Success Get Total Paid",
		Data:    totalPaid,
	})
}

func (pembayaranController *PembayaranController) GetTotalPaidKonstruksi(c *gin.Context) {

	kosntruksi_id := c.Param("konstruksi_id")

	totalPaid, err := pembayaranController.servicePembayaran.GetTotalPaidKonstruksi(kosntruksi_id)
	if err != nil {
		c.JSON(500, response.APIResponse{
			Status:  "error",
			Message: "Failed Get Total Paid",
		})

		return
	}

	c.JSON(200, response.APIResponse{
		Status:  "success",
		Message: "Success Get Total Paid",
		Data:    totalPaid,
	})
}

func (pembayaranController *PembayaranController) Update(c *gin.Context) {
	var request request.Pembayaran

	id := c.Param("id")
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

	pembayaranResponse, err := pembayaranController.servicePembayaran.Update(id, request)
	if err != nil {
		c.JSON(500, response.APIResponse{
			Status:  "error",
			Message: "Failed Save Data",
		})
		return
	}

	c.JSON(200, response.APIResponse{
		Status:  "success",
		Message: "Success Update Data",
		Data:    pembayaranResponse,
	})
}

func (pembayaranController *PembayaranController) Delete(c *gin.Context) {
	id := c.Param("id")

	err := pembayaranController.servicePembayaran.Delete(id)
	if err != nil {
		c.JSON(500, response.APIResponse{
			Status:  "error",
			Message: "Failed Delete Data",
		})
		return
	}

	c.JSON(200, response.APIResponse{
		Status:  "success",
		Message: "Success Delete Data",
	})
}
