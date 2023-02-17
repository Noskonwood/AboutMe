package main

import (
	"github.com/caarlos0/env/v7"
	"github.com/joho/godotenv"
	"log"
)

type Config struct {
	TelegramBotToken string `env:"API_KEY,required"`
}

func LoadConfig() Config {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	return cfg
}
