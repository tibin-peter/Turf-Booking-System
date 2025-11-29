package config


import (
	"fmt"
	"log"
	"os"

	"github.com/tibin-peter/Turf-Booking-System/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect with the database", err)
	}

	err = db.AutoMigrate(
		&model.User{},
		&model.Turf{},
		&model.TimeSlot{},
		&model.Payment{},
		&model.Booking{},
		&model.RefreshToken{},
	)
	if err != nil {
		log.Fatal("Migration failed", err)
	}

	fmt.Println("Auto migration completed successfully")

	DB = db
}
