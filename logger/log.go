package logger

import (
	"sync"

	log "github.com/sirupsen/logrus"
)

var (
	logger *log.Logger
	once   sync.Once
)

func InitLogger() {
	once.Do(func() {
		logger = log.New()
	})
}

func Log(level log.Level, message string, fields ...log.Fields) {
	logger.Log(level, message, fields)
}

func Debug(message string, fields ...log.Fields) {
	logger.Debug(message, fields)
}

func Info(message string, fields ...log.Fields) {
	logger.Info(message, fields)
}

func Warn(message string, fields ...log.Fields) {
	logger.Warn(message, fields)
}

func Error(message string, fields ...log.Fields) {
	logger.Error(message, fields)
}

func Fatal(message string, fields ...log.Fields) {
	logger.Fatal(message, fields)
}
