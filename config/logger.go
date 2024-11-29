package config

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var logger *zap.Logger

type ECSSyncWriteSyncer struct {
	esClient    *elasticsearch.Client
	appSettings *AppSettings
	index       string
}

type ECSLogEntry struct {
	Timestamp   string                 `json:"@timestamp"`
	LogLevel    string                 `json:"log.level"`
	Message     string                 `json:"message"`
	ServiceName string                 `json:"service.name"`
	Fields      map[string]interface{} `json:"fields"`
}

var notCheckedList []string = []string{"level", "msg"}

// Write converts zap logs to ECS format and sends them to Elasticsearch
func (syncer *ECSSyncWriteSyncer) Write(p []byte) (n int, err error) {
	// Unmarshal the zap log entry
	var logEntry map[string]interface{}
	if err := json.Unmarshal(p, &logEntry); err != nil {
		return 0, err
	}

	customFields := make(map[string]interface{})
	for key, value := range logEntry {
		if !contains(notCheckedList, key) {
			customFields[key] = value
		}
	}

	// Convert to ECS format
	ecsLog := ECSLogEntry{
		Timestamp:   time.Now().Format(time.RFC3339),
		LogLevel:    logEntry["level"].(string),
		Message:     logEntry["msg"].(string),
		ServiceName: syncer.appSettings.Project,
		Fields:      customFields,
	}

	// Marshal ECS log to JSON
	ecsLogBytes, err := json.Marshal(ecsLog)
	if err != nil {
		return 0, err
	}

	// Send to Elasticsearch
	res, err := syncer.esClient.Index(
		syncer.index,
		bytes.NewReader(ecsLogBytes),
		syncer.esClient.Index.WithContext(context.Background()),
	)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	return len(p), nil
}

func contains(list []string, key string) bool {

	for _, value := range list {
		if value == key {
			return true
		}
	}
	return false
}

// Sync is a no-op for this example
func (syncer *ECSSyncWriteSyncer) Sync() error {
	return nil
}

func UseEsLogger(appSettings *AppSettings, esClient *elasticsearch.Client) *zap.Logger {
	ecsSyncer := ECSSyncWriteSyncer{
		esClient:    esClient,
		index:       getLogIndex(appSettings),
		appSettings: appSettings,
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(&ecsSyncer),
		zap.InfoLevel,
	)

	zapLogger := zap.New(core)
	logger = zapLogger.With(
		zap.String("environment", appSettings.Environment),
		zap.String("project", appSettings.Project),
	)
	return logger
}

func GetLogger() *zap.Logger {
	return logger
}

func getLogIndex(appSettings *AppSettings) string {
	return fmt.Sprintf("golang-logger-%s-logs-%s", appSettings.Environment, time.Now().Format("02.01.2006"))
}
