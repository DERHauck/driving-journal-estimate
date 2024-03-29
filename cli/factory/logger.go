package factory

import (
	"derhauck/driving-journal-estimate/public/logger"
)

var staticLogger logger.Inf

func GetLogger() logger.Inf {
	if staticLogger == nil {
		staticLogger = logger.New(logger.DEFAULT)
	}
	return staticLogger
}

func SetLogLevel(value string) {
	level, err := logger.ParseLevel(value)
	if err != nil {
		GetLogger().Error(err)
		return
	}
	GetLogger().SetLevel(level)
}
