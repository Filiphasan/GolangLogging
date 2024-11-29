package config

import (
	"os"

	"github.com/joho/godotenv"
)

type AppSettings struct {
	Environment string
	Project     string
	Elastic     ElasticSettings
}

type ElasticSettings struct {
	User     string
	Password string
	Url      string
}

var appSettings AppSettings

func LoadAppSettings() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	appSettings = AppSettings{
		Environment: os.Getenv("ENVIRONMENT"),
		Project:     os.Getenv("PROJECT"),
		Elastic: ElasticSettings{
			User:     os.Getenv("ELASTIC_USER"),
			Password: os.Getenv("ELASTIC_PASSWORD"),
			Url:      os.Getenv("ELASTIC_HOST"),
		},
	}
}

func GetAppSettings() *AppSettings {
	return &appSettings
}
