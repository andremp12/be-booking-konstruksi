package service

import (
	database "booking-konstruksi/database/migration"
	"booking-konstruksi/repository"
	"booking-konstruksi/request"
	"booking-konstruksi/response"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type AuthService interface {
	Login(User database.User) (response.AuthUserResponse, error)
	Logout(token string) error
	ValidateToken(token string) (response.AuthUserResponse, error)
	Register(Request request.UserRequest) (response.User, error)
	CheckUser(email string) (database.User, error)
}

type authenticationService struct {
	authRepo repository.AuthenticationRepository
}

func NewAuthenticationService(authRepo repository.AuthenticationRepository) *authenticationService {
	return &authenticationService{authRepo: authRepo}
}

func (s *authenticationService) Login(User database.User) (response.AuthUserResponse, error) {

	//Generate JWT token
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Issuer:    "booking-konstruksi",
		ID:        string(User.ID),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	authUser := database.AuthUser{
		UserID:    User.ID,
		Role:      database.TypeRole(User.Role.Name), // Convert type string to TypeRole
		Token:     tokenStr,
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}

	err = s.authRepo.Login(authUser)
	fmt.Println("Role :", database.TypeRole(User.Role.Name))

	userResponse := response.UserResponse(&User)

	authResponse := response.AuthUserResponse{
		User:      userResponse,
		Role:      string(User.Role.Name), // Convert TypeRole to string
		Token:     tokenStr,
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}

	return authResponse, err
}

func (s *authenticationService) Logout(token string) error {
	err := s.authRepo.Logout(token)

	return err
}

func (s *authenticationService) ValidateToken(Token string) (response.AuthUserResponse, error) {
	var authUser database.AuthUser

	authUser, err := s.authRepo.ValidateToken(Token)

	// Create userResponse from auth.User
	userResponse := response.User{
		ID:     authUser.User.ID,
		Name:   authUser.User.Name,
		Email:  authUser.User.Email,
		NoWA:   authUser.User.NoWA,
		RoleID: authUser.User.RoleID,
	}

	// Create authResponse
	authResponse := response.AuthUserResponse{
		User:      &userResponse,
		UserId:    authUser.User.ID,
		Role:      string(authUser.Role), //Convert Type Role to String
		Token:     authUser.Token,
		ExpiresAt: authUser.ExpiresAt,
	}

	return authResponse, err
}

func (s *authenticationService) Register(request request.UserRequest) (response.User, error) {
	// Hashing for request password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)

	User := database.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: string(hashPassword),
		RoleID:   request.RoleID,
		NoWA:     request.NoWA,
	}

	User, err = s.authRepo.Register(User)

	UserResponse := response.UserResponse(&User)

	return *UserResponse, err
}

func (s *authenticationService) CheckUser(email string) (database.User, error) {
	var user database.User

	user, err := s.authRepo.CheckUser(email)

	return user, err
}
