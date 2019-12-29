package utils

import (
	"github.com/sirupsen/logrus"
)

// Logger entry
var Logger = logrus.New()

// InitLogger initiate app logger
func InitLogger() {
	Logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:    true,
		DisableSorting: false,
	})
}

// DebugLogging log in debug mode
func DebugLogging() {
	Logger.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
		DataKey:     "data",
	})
	Logger.SetReportCaller(true)
}
