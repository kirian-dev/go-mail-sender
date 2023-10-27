package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	log.Formatter = &logrus.TextFormatter{
		FullTimestamp: true,
	}

	log.Level = logrus.InfoLevel

	log.Out = os.Stdout
}

func GetLogger() *logrus.Logger {
	return log
}
