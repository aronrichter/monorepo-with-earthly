package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Log struct {
	localLog *zap.Logger
}

func New() *zap.Logger {
	config := zap.NewProductionConfig()

	// set log time info as we currently use
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.TimeKey = "time"

	// disable stacktrace and caller
	config.DisableStacktrace = true
	config.DisableCaller = true

	logger, err := config.Build()
	if err != nil {
		panic(err)
	}

	return logger
}

func (l *Log) Info(msg string, fields ...zap.Field) {
	l.localLog.Info(msg, fields...)
}

func (l *Log) Error(msg string, fields ...zap.Field) {
	l.localLog.Error(msg, fields...)
}
