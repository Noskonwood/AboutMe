package logger

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type ServiceLogger struct {
	logger *logrus.Logger
}

func (sl *ServiceLogger) log(entry *logrus.Entry) {
	entry = entry.WithFields(logrus.Fields{
		"service": os.Getenv("SERVICE_NAME"),
		"time":    entry.Time.Format("2006-01-02 15:04:05"),
	})
	sl.logger.WithFields(entry.Data).Log(entry.Level, entry.Message)
}

func (sl *ServiceLogger) Debug(args ...interface{}) {
	entry := logrus.NewEntry(sl.logger)
	entry.Message = fmt.Sprint(args...)
	entry.Level = logrus.DebugLevel
	sl.log(entry)
}

func (sl *ServiceLogger) Info(args ...interface{}) {
	entry := logrus.NewEntry(sl.logger)
	entry.Message = fmt.Sprint(args...)
	entry.Level = logrus.InfoLevel
	sl.log(entry)
}

func Init() *ServiceLogger {
	// Load the environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		logrus.Fatalf("Error reading .env file: %v", err)
	}

	// Set the logger level based on the value in the .env file
	logLevel, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		logLevel = logrus.InfoLevel
	}

	// Set the logger output to a file based on the value in the .env file
	logFile := os.Getenv("LOG_OUTPUT_FILE")
	var file *os.File
	if logFile != "" {
		file, err = os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			logrus.Warnf("Failed to logger to file, using default stderr: %v", err)
			file = os.Stderr
		}
	} else {
		file = os.Stderr
	}

	// Configure the formatter to output JSON
	formatter := &logrus.JSONFormatter{}
	logger := logrus.New()
	logger.SetLevel(logLevel)
	logger.SetOutput(file)
	logger.SetFormatter(formatter)

	return &ServiceLogger{logger: logger}
}
