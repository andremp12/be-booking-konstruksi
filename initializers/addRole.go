package initializers

import (
	database "booking-konstruksi/database/migration"
	"gorm.io/gorm"
)

func AddRole(db *gorm.DB) {
	roles := []database.Role{
		{Name: "admin"},
		{Name: "mandor"},
		{Name: "client"},
	}

	for _, role := range roles {
		db.FirstOrCreate(&role, database.Role{Name: role.Name})
	}
}
