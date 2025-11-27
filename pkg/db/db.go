package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/tibin-peter/Turf-Booking-System/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is missing in environment varialbels")
	}

	// Create a custom logger for debugging DB operations
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // Log slow queries
			LogLevel:      logger.Info, // Log DB actions
			Colorful:      true,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		log.Fatal("Failed to connect with the database", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Error getting database instance:", err)
	}

	// Connection pooling config
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Quick test
	if err := sqlDB.Ping(); err != nil {
		log.Fatal("Database ping failed:", err)
	}

	fmt.Println(" Database connection established successfully!")

	err = db.AutoMigrate(
		&model.User{},
		&model.Turf{},
		&model.TimeSlot{},
		&model.Payment{},
		&model.Booking{},
	)
	if err != nil {
		log.Fatal("Migration failed", err)
	}

	fmt.Println("Auto migration completed successfully")

	DB = db
}
