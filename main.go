package main

import (
	"booking-konstruksi/controller"
	"booking-konstruksi/database"
	"booking-konstruksi/initializers"
	"booking-konstruksi/middleware"
	"booking-konstruksi/repository"
	"booking-konstruksi/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // Allow all domain
		//c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		//c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		//c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		//c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		//c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		//
		//// Handle OPTIONS method
		//if c.Request.Method == "OPTIONS" {
		//	c.AbortWithStatus(http.StatusOK)
		//	return
		//}

		c.Next()
	}
}

func init() {
	initializers.LoadEnv()
}

func main() {
	dsn := "root:password@tcp(127.0.0.1:3306)/db_booking_konstruksi?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	err = database.AutoMigrate(db)
	if err != nil {
		panic("Failed to migrate database")
	}

	// Add role data
	initializers.AddRole(db)

	// Authentication
	repoAuthentication := repository.NewAuthenticationRepository(db)
	serviceAuthentication := service.NewAuthenticationService(repoAuthentication)
	authenticationController := controller.NewAuthenticationController(serviceAuthentication)

	// Tipe Konstruksi
	repoTipeKonstruksi := repository.NewRepoTipeKonstruksi(db)
	serviceTipeKonstruksi := service.NewServiceTipeKonstruksi(repoTipeKonstruksi)
	tipeKonstruksiController := controller.NewTipeKonstruksiController(serviceTipeKonstruksi)

	// Satuan
	repoSatuan := repository.NewRepoSatuan(db)
	serviceSatuan := service.NewServiceSatuan(repoSatuan)
	satuanController := controller.NewSatuanController(serviceSatuan)

	// Konstruksi
	repoKonstruksi := repository.NewRepoKonstruksi(db)
	serviceKonstruksi := service.NewServiceKonstruksi(repoKonstruksi)
	konstruksiController := controller.NewKonstruksiController(serviceKonstruksi)

	// Time Line Konstruksi
	repoTimeLine := repository.NewRepositoryTimelineKonstruksi(db)
	serviceTimeLine := service.NewServiceTimelineKonstruksi(repoTimeLine)
	timeLineController := controller.NewTimeLineController(serviceTimeLine)

	// Time Line Konstruksi
	repoLaporan := repository.NewRepositoryLaporanKonstruksi(db)
	serviceLaporan := service.NewServiceLaporanKonstruksi(repoLaporan)
	laporanController := controller.NewLaporanKonstruksiController(serviceLaporan)

	// Client
	repoClient := repository.NewRepoClient(db)
	serviceClient := service.NewServiceClient(repoClient)
	clientController := controller.NewMandorController(serviceClient)

	// Mandor
	repoMandor := repository.NewRepoMandor(db)
	serviceMandor := service.NewServiceMandor(repoMandor)
	mandorController := controller.NewMandorController(serviceMandor)

	// Pembayaran
	repoPembayaran := repository.NewRepoPembayaran(db)
	servicePembayaran := service.NewServicePembayaran(repoPembayaran)
	pembayaranController := controller.NewPembayaranController(servicePembayaran)

	// Midtrans
	serviceMidtrans := service.NewMidtransService(repoPembayaran)
	midtransController := controller.NewMidtransController(serviceMidtrans)

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	//router.Use(CORSMiddleware()) // allow cors on API Gateway

	v1 := router.Group("/api/v1")
	v1.POST("/register", authenticationController.Register)
	v1.POST("/login", authenticationController.Login)
	v1.GET("/tipe-konstruksi/landing", tipeKonstruksiController.GetTipeLanding)
	v1.POST("/tipe-konstruksi", tipeKonstruksiController.CreateData)

	authorized := v1.Use(middleware.Auth(db))

	authorized.POST("/validate-token", authenticationController.ValidateToken)
	authorized.POST("/logout", authenticationController.Logout)

	// API tipe konstruksi

	authorized.GET("/tipe-konstruksi", tipeKonstruksiController.GetAllData)
	authorized.GET("/tipe-konstruksi/:id", tipeKonstruksiController.GetData)
	authorized.PUT("/tipe-konstruksi/:id", tipeKonstruksiController.UpdateData)
	authorized.DELETE("/tipe-konstruksi/:id", tipeKonstruksiController.Delete)

	// API satuan
	authorized.GET("/satuan", satuanController.GetAllData)
	authorized.GET("/satuan/:id", satuanController.GetData)
	authorized.POST("/satuan", satuanController.CreateData)
	authorized.PUT("/satuan/:id", satuanController.UpdateData)
	authorized.DELETE("/satuan/:id", satuanController.Delete)
	router.Static("/images", "./images")

	// API konstruksi
	authorized.GET("/riwayat-konstruksi", konstruksiController.GetRiwayatKonstruksi)
	authorized.GET("/mandor/riwayat-konstruksi", konstruksiController.GetRiwayatKonstruksiMandor)
	authorized.GET("/client/riwayat-konstruksi", konstruksiController.GetRiwayatKonstruksiClient)
	authorized.GET("/konstruksi/get-count-status", konstruksiController.GetCountStatus)
	authorized.GET("/konstruksi", konstruksiController.GetAllData)
	authorized.GET("/konstruksi/:id", konstruksiController.GetData)
	authorized.GET("/client/konstruksi", konstruksiController.GetKonstruksiUser)
	authorized.GET("/mandor/konstruksi", konstruksiController.GetKonstruksiMandor)
	authorized.POST("/konstruksi/:tipe_id", konstruksiController.Booking)
	authorized.PUT("/konstruksi/update-status/:id", konstruksiController.UpdateStatus)
	authorized.PUT("/konstruksi/confirmation/:id", konstruksiController.KonfirmasiBooking)

	//API Timeline Konstruksi
	authorized.GET("/timeline-konstruksi/:konstruksi_id", timeLineController.GetAllData)
	authorized.POST("/timeline", timeLineController.Create)
	authorized.PUT("/timeline/:id", timeLineController.Update)
	authorized.DELETE("/timeline/:id", timeLineController.Delete)

	//API Laporan Konstruksi
	authorized.GET("/laporan-konstruksi/:konstruksi_id", laporanController.GetLaporanKonstruksi)
	authorized.POST("/laporan-konstruksi", laporanController.Create)

	// API Mandor
	authorized.GET("/mandor", mandorController.GetAllData)

	// API Client
	authorized.GET("/client", clientController.GetAllData)

	//API Pembayaran
	authorized.GET("/riwayat-pembayaran", pembayaranController.GetRiwayatPembayaran)
	authorized.GET("/client/riwayat-pembayaran", pembayaranController.GetRiwayatPembayaranClient)
	authorized.GET("/pembayaran/:konstruksi_id", pembayaranController.GetPembayaranClient)
	authorized.GET("/total-paid/konstruksi", pembayaranController.GetTotalPaid)
	authorized.GET("/total-paid/konstruksi/:konstruksi_id", pembayaranController.GetTotalPaidKonstruksi)
	authorized.POST("/pembayaran/:konstruksi_id", pembayaranController.Create)
	authorized.POST("/midtrans/:id", midtransController.Create)
	authorized.POST("/success-payment/:id", pembayaranController.SuccessPayment)

	router.Run()
}
