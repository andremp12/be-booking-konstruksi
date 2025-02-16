package controller

import (
	"booking-konstruksi/request"
	"booking-konstruksi/response"
	"booking-konstruksi/service"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type tipeKonstruksiController struct {
	serviceTipeKonstruksi service.TipeKonstruksi
}

func NewTipeKonstruksiController(serviceTipeKonstruksi service.TipeKonstruksi) *tipeKonstruksiController {
	return &tipeKonstruksiController{serviceTipeKonstruksi: serviceTipeKonstruksi}
}

func (tipeController *tipeKonstruksiController) GetTipeLanding(c *gin.Context) {
	tipeKonstruksis, err := tipeController.serviceTipeKonstruksi.GetTipeLanding()

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
		Message: "Success Get All Tipe Konstruksi",
		Data:    tipeKonstruksis,
	})
}

func (tipeController *tipeKonstruksiController) GetAllData(c *gin.Context) {
	tipeKonstruksis, err := tipeController.serviceTipeKonstruksi.GetAllData()

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
		Message: "Success Get All Tipe Konstruksi",
		Data:    tipeKonstruksis,
	})
}

func (tipeController *tipeKonstruksiController) GetData(c *gin.Context) {
	id := c.Param("id")

	tipeKonstruksi, err := tipeController.serviceTipeKonstruksi.GetData(id)
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
		Message: "Success Get Tipe Konstruksi",
		Data:    tipeKonstruksi,
	})
}

func (tipeController *tipeKonstruksiController) CreateData(c *gin.Context) {
	var requestTipe request.TipeKonstruksi

	err := c.ShouldBind(&requestTipe)
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

	//file, err := c.FormFile("file_image")
	//if err != nil {
	//	c.JSON(400, response.APIResponse{
	//		Status:  "error",
	//		Message: "File required",
	//		Data:    err.Error(),
	//	})
	//
	//	return
	//}

	//Mendapatkan mengubah nama file dan mendefinisikan path file untuk disimpan ke database
	//fileExt := filepath.Ext(file.Filename)
	if requestTipe.FileImage != nil {
		fileName := fmt.Sprintf("%d-%s", time.Now().UnixNano(), requestTipe.FileImage.Filename)
		filepath := filepath.Join("images", fileName)
		requestTipe.Image = fmt.Sprintf("%s/%s", os.Getenv("BASE_URL"), filepath)

		//Simpan file ke dalam folder images dengan nama yang telah sesuai format
		err = c.SaveUploadedFile(requestTipe.FileImage, filepath)
		if err != nil {
			c.JSON(500, response.APIResponse{
				Status:  "error",
				Message: "Failed Save File",
				Data:    err.Error(),
			})

			return
		}
	}

	//Simpan path file ke database
	responseTipekonstruksi, err := tipeController.serviceTipeKonstruksi.Create(requestTipe)
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
		Message: "Success Create Tipe Konstruksi",
		Data:    responseTipekonstruksi,
	})
}

func (tipeController *tipeKonstruksiController) UpdateData(c *gin.Context) {
	var requestTipe request.TipeKonstruksi

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

	//Mendapatkan mengubah nama file dan mendefinisikan path file untuk disimpan ke database
	//fileExt := filepath.Ext(file.Filename)
	if requestTipe.FileImage != nil {
		fileName := fmt.Sprintf("%d-%s", time.Now().UnixNano(), requestTipe.FileImage.Filename)
		filepath := filepath.Join("images", fileName)
		requestTipe.Image = fmt.Sprintf("%s/%s", os.Getenv("BASE_URL"), filepath)

		//Simpan file ke dalam folder images dengan nama yang telah sesuai format
		err = c.SaveUploadedFile(requestTipe.FileImage, filepath)
		if err != nil {
			c.JSON(500, response.APIResponse{
				Status:  "error",
				Message: "Failed Save File",
				Data:    err.Error(),
			})

			return
		}

		//Simpan file ke dalam folder images dengan nama yang telah sesuai format
		// _, err = os.Stat(filepath)
		// if os.IsNotExist(err) {
		// 	err = c.SaveUploadedFile(requestTipe.FileImage, filepath)

		// 	if err != nil {
		// 		c.JSON(500, response.APIResponse{
		// 			Status:  "error",
		// 			Message: "Failed Save File",
		// 			Data:    err.Error(),
		// 		})

		// 		return
		// 	}
		// }
	}

	//Simpan path file ke database
	responseTipekonstruksi, err := tipeController.serviceTipeKonstruksi.Update(requestTipe, id)
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
		Data:    responseTipekonstruksi,
	})
}

func (tipeController *tipeKonstruksiController) Delete(c *gin.Context) {
	id := c.Param("id")

	err := tipeController.serviceTipeKonstruksi.Delete(id)

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
