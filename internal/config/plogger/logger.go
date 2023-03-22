package plogger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type (
	LogData struct {
		Source  string
		Message string
		Data    interface{}
	}

	Map map[string]interface{}
)

var (
	zapLogger *zap.Logger
	err       error
)

// Init ...
func Init(appName, server string) {
	cfg := zap.Config{
		Encoding:      "console",
		Level:         zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths:   []string{"stdout"},
		InitialFields: map[string]interface{}{"server": server, "capture": appName},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:  "message",
			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,
			TimeKey:     "time",
			EncodeTime:  zapcore.ISO8601TimeEncoder,
		},
	}
	zapLogger, err = cfg.Build()
	if err != nil {
		panic(err)
	}
}
