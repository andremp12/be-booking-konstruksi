package controller

import (
	"booking-konstruksi/request"
	"booking-konstruksi/response"
	"booking-konstruksi/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type timeLineController struct {
	serviceTimeLine service.TimeLineKonstruksi
}

func NewTimeLineController(serviceTimeLine service.TimeLineKonstruksi) *timeLineController {
	return &timeLineController{serviceTimeLine: serviceTimeLine}
}

func (timeLineController *timeLineController) GetAllData(c *gin.Context) {
	konstruksi_id := c.Param("konstruksi_id")

	timeLines, err := timeLineController.serviceTimeLine.GetTimeLineKonstruksi(konstruksi_id)

	if err != nil {
		c.JSON(400, response.APIResponse{
			Status:  "error",
			Message: "Failed Get All Data",
			Data:    err.Error(),
		})

		return
	}

	// response := gin.H{
	// 	"records":timeLines,
	// }

	c.JSON(200, response.APIResponse{
		Status:  "success",
		Message: "Success Get All Data",
		Data:    timeLines,
	})
}

func (timeLineController *timeLineController) Create(c *gin.Context) {
	var request request.TimeLine

	err := c.ShouldBind(&request)
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

	err = timeLineController.serviceTimeLine.Create(request)
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
	})
}

func (timeLineController *timeLineController) Update(c *gin.Context) {
	var request request.TimeLine

	//id := c.Param("id")

	err := c.ShouldBind(&request)
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

	err = timeLineController.serviceTimeLine.Update(request)
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
	})
}

func (timeLineController *timeLineController) Delete(c *gin.Context) {
	id := c.Param("id")

	err := timeLineController.serviceTimeLine.Delete(id)

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
