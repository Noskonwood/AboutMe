package config

import (
	"github.com/caarlos0/env/v7"
	"github.com/joho/godotenv"
	"log"
	"os"
	_ "os"
)

type BotInfastructureConfig struct {
	TelegramBotToken string `env:"API_KEY,required"`
	LogLevel         string `env:"LOG_LEVEL"`
	LogServer        string `env:"LOG_SERVER"`
	ServiceName      string `env:"SERVICE_NAME"`
	LogOutput        string `env:"LOG_OUTPUT_FILE"`
	//Host_port   string  `env:"HOST_PORT"`
}

func NewBotInfastructureConfig() BotInfastructureConfig {
	var botInfastructureConfig BotInfastructureConfig

	err := env.Parse(&botInfastructureConfig)
	if err != nil {
		return botInfastructureConfig
	}

	loadFile := godotenv.Load()
	if loadFile != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	// Parse environment variables into config struct
	botInfastructureConfig.TelegramBotToken = os.Getenv("API_KEY")
	botInfastructureConfig.LogLevel = os.Getenv("LOG_LEVEL")
	botInfastructureConfig.LogServer = os.Getenv("LOG_SERVER")
	botInfastructureConfig.ServiceName = os.Getenv("SERVICE_NAME")
	botInfastructureConfig.LogOutput = os.Getenv("LOG_OUTPUT_FILE")

	// Validate required fields
	if botInfastructureConfig.TelegramBotToken == "" {
		log.Fatalf("missing required field: API_KEY - %v", err)

	}

	return botInfastructureConfig
}
