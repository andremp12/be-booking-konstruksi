package controller

import (
	"booking-konstruksi/request"
	"booking-konstruksi/response"
	"booking-konstruksi/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
)

type MidtransController struct {
	serviceMidtrans service.MidtransService
}

func NewMidtransController(serviceMidtrans service.MidtransService) *MidtransController {
	return &MidtransController{
		serviceMidtrans: serviceMidtrans,
	}
}

func (midtransController *MidtransController) Create(c *gin.Context) {
	var request request.MidtransRequest

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

	midtransResponse, err := midtransController.serviceMidtrans.Create(request, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.APIResponse{
		Status:  "success",
		Message: "success",
		Data:    midtransResponse,
	})
}
