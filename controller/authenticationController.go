package controller

import (
	"booking-konstruksi/request"
	"booking-konstruksi/response"
	"booking-konstruksi/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
)

type AuthenticationController struct {
	authService service.AuthService
}

func NewAuthenticationController(authService service.AuthService) *AuthenticationController {
	return &AuthenticationController{authService: authService}
}

func (authController *AuthenticationController) Login(c *gin.Context) {
	var requestAuth request.LoginRequest

	err := c.ShouldBind(&requestAuth)
	fmt.Println(requestAuth)

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

	user, err := authController.authService.CheckUser(requestAuth.Email)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.APIResponse{
			Status:  "error",
			Message: "Bad Credentials",
			Data:    err.Error(),
		})

		return
	}

	if user.ID == 0 {
		c.JSON(404, response.APIResponse{
			Status:  "unathorized",
			Message: "Email or Password is Invalid",
			Data:    err.Error(),
		})

		return
	}

	//Compare user password  with request password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestAuth.Password))

	if err != nil {
		c.JSON(404, response.APIResponse{
			Status:  "unathorized",
			Message: "Email or Password is Invalid",
			Data:    err.Error(),
		})

		return
	}

	// Saved auth User
	authUser, err := authController.authService.Login(user)
	if err != nil {
		c.JSON(500, response.APIResponse{
			Status:  "error",
			Message: "Internal Server Error",
			Data:    err.Error(),
		})
		return
	}

	c.JSON(200, response.APIResponse{
		Status:  "success",
		Message: "Success Login",
		Data:    authUser,
	})
}

func (authController *AuthenticationController) Register(c *gin.Context) {
	var requestUser request.UserRequest

	err := c.ShouldBind(&requestUser)

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

	user, err := authController.authService.Register(requestUser)
	if err != nil {
		c.JSON(500, response.APIResponse{
			Status:  "error",
			Message: "Internal Server Error",
			Data:    err.Error(),
		})

		return
	}

	c.JSON(200, response.APIResponse{
		Status:  "success",
		Message: "Success Register",
		Data:    user,
	})
}

func (authController *AuthenticationController) ValidateToken(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")

	// Check the prefix header as a Bearer Token
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{
			"messages": "Authorization token not provided",
		})

		return
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

	authResponse, err := authController.authService.ValidateToken(tokenStr)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"messages": "Unauthorization",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   "success",
		"messages": "Success Retrive Data",
		"data":     authResponse,
	})

}

func (authController *AuthenticationController) Logout(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{
			"messages": "Authorization token not provided",
		})

		return
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

	err := authController.authService.Logout(tokenStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"messages": "Internal Server Error",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   "success",
		"messages": "Success Logout",
	})
}
