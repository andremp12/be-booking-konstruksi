package response

import (
	"time"
)

type AuthUserResponse struct {
	User      *User     `json:"user"`
	UserId    int       `json:"user_id"`
	Token     string    `json:"token"`
	Role      string    `json:"role"`
	ExpiresAt time.Time `json:"expires_at"`
}
