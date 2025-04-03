package utils

import (
	"os"

	"github.com/charmbracelet/log"
)

func GetDefaultLogger() *log.Logger {

	var level log.Level

	if os.Getenv("DEBUG") != "" {
		level = log.DebugLevel
	} else {
		level = log.ErrorLevel
	}

	return log.NewWithOptions(os.Stdout, log.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
		Level:           level,
	})
}
