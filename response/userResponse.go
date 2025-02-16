package response

import database "booking-konstruksi/database/migration"

type User struct {
	ID     int    `json:"id"`
	RoleID int    `json:"role_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	NoWA   string `json:"no_wa"`
}

func UserResponse(data *database.User) *User {
	return &User{
		ID:     data.ID,
		Name:   data.Name,
		RoleID: data.RoleID,
		Email:  data.Email,
		NoWA:   data.NoWA,
	}
}
