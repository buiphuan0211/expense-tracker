package plogger

import (
	"encoding/json"
	"go.uber.org/zap"
)

// Debug ...
func Debug(content string, data LogData) {
	jsonData, _ := json.Marshal(data)
	zapLogger.Debug(content, zap.String("data", string(jsonData)))
}
