package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func ex1() {
	//cfg := zap.NewProductionConfig()
	//cfg.Sampling = nil
	//cfg.DisableStacktrace = true
	//cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	//cfg.EncoderConfig.CallerKey = "logLine"
	////cfg.EncoderConfig.EncodeTime = syslogTimeEncoder
	////cfg.EncoderConfig.EncodeLevel = customLevelEncoder
	////cfg.Level = zap.NewAtomicLevelAt(config.LogLevel)
	//
	//logger, _ := cfg.Build()
	//
	//logger.Info("This should have a syslog style timestamp")
	//logger.Error("aaa")
	//logger.Warn("qqq")
	//logger.Debug("2222")
	////logger.Panic("333")
}

//func syslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
//	enc.AppendString(t.Format("2006-01-02 15:04:05"))
//}
//
//func customLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
//	enc.AppendString("[" + level.String() + "]")
//}

// Global system logger.
var sysLog *zap.Logger

// Init system logger.
func init() {
	sysLog = NewLogger()
}

func NewLogger() *zap.Logger {
	cfg := zap.NewProductionConfig()
	cfg.Sampling = nil
	cfg.DisableStacktrace = true
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	cfg.EncoderConfig.CallerKey = "logLine"

	logger, _ := cfg.Build()
	return logger
}

func main() {
	sysLog.Info("Info log")
	sysLog.Error("Error log")
	sysLog.Warn("Warn log")
}
