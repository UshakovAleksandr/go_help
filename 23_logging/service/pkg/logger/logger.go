package logger

import (
	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type AllLogger struct {
	Log   *zap.Logger
	LogMV func(next httprouter.Handle) httprouter.Handle
}

func NewAllLogger() AllLogger {
	log := AllLogger{
		Log:   NewLogger(),
		LogMV: NewLoggerMiddleware(NewLogger()),
	}

	return log
}

func NewLogger() *zap.Logger {
	var level zapcore.Level
	level = zapcore.InfoLevel

	cfg := zap.NewProductionConfig()
	cfg.Sampling = nil
	cfg.DisableStacktrace = true
	cfg.Level = zap.NewAtomicLevelAt(level)
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	cfg.EncoderConfig.CallerKey = "logLine"

	logger, err := cfg.Build()
	if err != nil {
		logger.Error("logger init error")
	}

	return logger
}
