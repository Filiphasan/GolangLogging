package main

import (
	"fmt"
	"github.com/Filiphasan/GolangLogging/config"
	"go.uber.org/zap"
)

func main() {
	config.LoadAppSettings()
	appSettings := config.GetAppSettings()

	fmt.Println(appSettings.Environment)

	esClient := config.UseEsClient(appSettings)

	logger := config.UseEsLogger(appSettings, esClient)
	defer logger.Sync()

	logger.Info("Hello, World!", zap.String("name", "Hasan"))

}

// Minifying Url
