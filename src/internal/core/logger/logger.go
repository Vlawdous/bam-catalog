package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct {
	*zap.Logger
}

func NewLogger(logFilePath string, env string) *Logger {
	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	writer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logFilePath + "/book.log",
		MaxSize:    100,
		MaxBackups: 3,
		MaxAge:     28,
	})

	logLevel := zapcore.DebugLevel

	if env == "production" {
		logLevel = zapcore.ErrorLevel
	}

	zapCore := zapcore.NewCore(encoder, writer, logLevel)
	logger := zap.New(zapCore)

	return &Logger{logger: logger}
}
