package logger

import (
	"git.foxminded.ua/foxstudent104181/telegrambot/config"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"log"
)

const (
	developmentLevel = "DEBUG"
	productionLevel  = "PRODUCTION"
)

func NewBotInfrastructureLogger(string) (*zap.SugaredLogger, error) {
	var l *zap.Logger
	var err error

	cfg := config.NewBotInfastructureConfig()

	// Use the configuration variables to create a logger
	logger, err := NewBotInfrastructureLogger(cfg.LogLevel)
	if err != nil {
		logrus.Fatalf("Failed to create logger: %v", err)
	}

	// Use the logger
	logger.Info("This is an example log message.")

	switch cfg.LogLevel {
	case productionLevel:
		l, err = zap.NewProduction()
		if err != nil {
			return nil, err
		}
	case developmentLevel:
		l, err = zap.NewDevelopment()
		if err != nil {
			return nil, err
		}
	default:
		l, err = zap.NewDevelopment()
		if err != nil {
			return nil, err
		}
	}

	sugar := l.Sugar()

	return sugar, nil
}

func Close(l *zap.SugaredLogger) {
	err := l.Sync()
	if err != nil {
		log.Println("Couldn't flush logging buffer")
	}
}
