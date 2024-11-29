package config

import (
	"github.com/elastic/go-elasticsearch/v8"
)

func UseEsClient(appSettings *AppSettings) *elasticsearch.Client {
	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Username: appSettings.Elastic.User,
		Password: appSettings.Elastic.Password,
		Addresses: []string{
			appSettings.Elastic.Url,
		},
	})

	if err != nil {
		panic(err)
	}

	return client
}
