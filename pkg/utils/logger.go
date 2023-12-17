package utils

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func init() {
	Log = logrus.New()

	// Determine log format based on an environment variable
	logFormat := os.Getenv("LOG_FORMAT")
	if strings.ToLower(logFormat) == "json" {
		Log.SetFormatter(&logrus.JSONFormatter{}) // Structured logging
	} else {
		Log.SetFormatter(&logrus.TextFormatter{ // Human-readable logging
			FullTimestamp: true,
		})
	}

	// TODO: Make log level configurable as well
	Log.SetLevel(logrus.WarnLevel) // Default log level
	Log.SetOutput(os.Stdout)
}
