package custom_logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var custom_logger *zap.Logger

func init() {
	GenerateLogger()
}

func GenerateLogger() *zap.Logger {
	// zap setting.
	config := zap.NewDevelopmentConfig()
	// output destination
	config.OutputPaths = []string{"logging.txt", "stderr"}
	// input setting
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig = encoderConfig
	custom_logger, _ = config.Build()
	return custom_logger
}

func Info(message string, fields ...zap.Field) {
	custom_logger.Info(message, fields...)
}
