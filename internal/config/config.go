package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUrl     string
	JwtSecret string
	Port      string
}

var AppConfig Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("NO  .env file found")
	}
	AppConfig = Config{
		DBUrl:     os.Getenv("DATABASE_URL"),
		JwtSecret: os.Getenv("JwtSecret"),
		Port:      os.Getenv("PORT"),
	}
	if AppConfig.DBUrl == "" || AppConfig.JwtSecret == "" {
		log.Fatal("missing the environment varialbels")
	}

}
