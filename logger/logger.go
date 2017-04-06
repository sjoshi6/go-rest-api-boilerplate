package logger

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/markbates/going/defaults"
)

var logger *logrus.Logger

func GetLogger() *logrus.Logger {

	if logger == nil {
		logger = logrus.New()
		var ENV = defaults.String(os.Getenv("GO_EUN_ENV"), "development")

		logger.Out = os.Stdout
		logger.Formatter = &logrus.TextFormatter{}

		if ENV == "development" {
			logger.Level = logrus.DebugLevel
		} else {
			logger.Level = logrus.InfoLevel
		}
	}

	return logger
}
