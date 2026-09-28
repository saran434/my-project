package setup

import (
	"log"

	"employeejwt/internal/database"
	"employeejwt/internal/models"
)

func SetupDatabase() {
	database.ConnectDB()

	if err := database.DB.AutoMigrate(&models.User{}, &models.Employee{}); err != nil {
		log.Fatalf("migrate database schema: %v", err)
	}
}
