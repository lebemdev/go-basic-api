package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// defining envireoment variables Structure
type Config struct {
	DB_URI     string
	PORT       string
	JWT_SECRET string
	//UPSTASH_URI string
}

func SetConfig() (*Config, error) {
	//if os.Getenv("APP_ENV") == "dev" {
	godotenv.Load()
	//}

	port := os.Getenv("PORT")
	fmt.Printf("PORT from the .env : %s", port)

	db_uri := os.Getenv("DB_URI")
	fmt.Printf("Database URI : %s", db_uri)

	return &Config{
		DB_URI:     os.Getenv("DB_URI"),
		PORT:       os.Getenv("PORT"),
		JWT_SECRET: os.Getenv("JWT_SECRET"),
		//UPSTASH_URI: viper.GetString("UPSTASH_URI"),
	}, nil
}
