package logging

import (
	"github.com/joho/godotenv"
	"os"

	"github.com/sirupsen/logrus"
)

func Init() {
	// Load the environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		logrus.Fatalf("Error to read .env file %v", err)
	}

	// Set the logging level based on the value in the .env file
	logLevel, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		logLevel = logrus.InfoLevel
	}
	logrus.SetLevel(logLevel)

	// Set the logging output to a file based on the value in the .env file
	logFile := os.Getenv("LOG_OUTPUT_FILE")
	if logFile != "" {
		file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err == nil {
			logrus.SetOutput(file)
		} else {
			logrus.Warnf("Failed to logging to file, using default stderr: %v", err)
		}
	}

	// Configure the formatter to output JSON
	logrus.SetFormatter(&logrus.JSONFormatter{})
}
