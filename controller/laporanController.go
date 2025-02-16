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
	"os"
	"path/filepath"
	"time"
)

type laporanKonstruksiController struct {
	serviceLaporanKonstruksi service.LaporanKonstruksi
}

func NewLaporanKonstruksiController(serviceLaporanKonstruksi service.LaporanKonstruksi) *laporanKonstruksiController {
	return &laporanKonstruksiController{serviceLaporanKonstruksi: serviceLaporanKonstruksi}
}

func (laporanController *laporanKonstruksiController) GetLaporanKonstruksi(c *gin.Context) {
	konstruksiId := c.Param("konstruksi_id")

	responseLaporan, err := laporanController.serviceLaporanKonstruksi.GetLaporanKonstruksi(konstruksiId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.APIResponse{
			Status:  "error",
			Message: "Failed to get data",
			Data:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.APIResponse{
		Status:  "success",
		Message: "Success get data",
		Data:    responseLaporan,
	})
}

func (laporanController *laporanKonstruksiController) Create(c *gin.Context) {
	var request request.Laporan

	fmt.Println("Laporan : ", c.PostForm("bahan_konstruksi"))

	err := c.ShouldBind(&request)
	request.BahanKonstruksi = c.Request.FormValue("bahan_konstruksi")
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

	//Mendapatkan mengubah nama file dan mendefinisikan path file untuk disimpan ke database
	//fileExt := filepath.Ext(file.Filename)
	if request.FileImage != nil {
		fileName := fmt.Sprintf("%d-%s", time.Now().UnixNano(), request.FileImage.Filename)
		filepath := filepath.Join("images", fileName)
		request.Image = fmt.Sprintf("%s/%s", os.Getenv("BASE_URL"), filepath)

		//Simpan file ke dalam folder images dengan nama yang telah sesuai format
		err = c.SaveUploadedFile(request.FileImage, filepath)
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
	err = laporanController.serviceLaporanKonstruksi.Create(request)
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
	})
}
