package main

import (
	"fmt"
	"github.com/Filiphasan/GolangLogging/config"
	"github.com/Filiphasan/GolangLogging/src/services"
	"go.uber.org/zap"
	"time"
)

func main() {
	config.LoadAppSettings()
	appSettings := config.GetAppSettings()

	fmt.Println(appSettings.Environment)

	esClient := config.UseEsClient(appSettings)

	logger := config.UseEsLogger(appSettings, esClient)
	defer logger.Sync()

	services.CreateSomeLog(12)

	logger.Info("Program finished", zap.Time("FinishTime", time.Now()))
}
