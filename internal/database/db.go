package database

import (
	"fmt"
	"log"
	"os"

	"employeejwt/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	cfg := config.Load()

	host := config.Pick(cfg.Host, "DB_HOST")
	port := config.Pick(cfg.Port, "DB_PORT")
	user := config.Pick(cfg.User, "DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := config.Pick(cfg.DBName, "DB_NAME")
	sslmode := config.Pick(cfg.SSLMode, "DB_SSLMODE")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("database open failed:", err)
	}
	log.Println("database opened successfully")

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("database handle failed:", err)
	}
	if err := sqlDB.Ping(); err != nil {
		log.Fatal("database connection failed:", err)
	}

	DB = db
	log.Println("database connected successfully")
}
