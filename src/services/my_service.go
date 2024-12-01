package services

import (
	"github.com/Filiphasan/GolangLogging/config"
	"go.uber.org/zap"
)

func CreateSomeLog(logCount int) {
	logger := config.GetLogger()

	for i := 0; i < logCount; i++ {
		logger.Info("Hello, World!", zap.Int32("LogCount", int32(i)))
	}
}

func CreateSomeErrorLog(logCount int) {
	logger := config.GetLogger()

	for i := 0; i < logCount; i++ {
		logger.Error("Hello, World!", zap.Int32("LogCount", int32(i)))
	}
}
