package request

type UserRequest struct {
	RoleID   int    `form:"role_id" json:"role_id" binding:"required"`
	Name     string `form:"name" binding:"required"`
	Email    string `form:"email" binding:"required"`
	NoWA     string `form:"no_wa" json:"no_wa" binding:"required"`
	Password string `form:"password" binding:"required"`
}
