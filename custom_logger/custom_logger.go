package custom_logger

import (
	"fmt"
	"io"
	"log"
	"os"
)

// import (
// 	"go.uber.org/zap"
// 	"go.uber.org/zap/zapcore"
// )

// var custom_logger *zap.Logger

// func init() {
// 	GenerateLogger()
// }

// func GenerateLogger() *zap.Logger {
// 	// zap setting.
// 	config := zap.NewDevelopmentConfig()
// 	// output destination
// 	config.OutputPaths = []string{"logging.txt", "stderr"}
// 	// input setting
// 	encoderConfig := zap.NewDevelopmentEncoderConfig()
// 	encoderConfig.TimeKey = "timestamp"
// 	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
// 	config.EncoderConfig = encoderConfig
// 	custom_logger, _ = config.Build()
// 	return custom_logger
// }

// 	func Info(message any, fields ...zap.Field) {
// 	custom_logger.Info(message, fields...)
// }

func Info(v ...any) {
	f, err := os.OpenFile("logging.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(io.MultiWriter(f, os.Stdin))
	log.Println(v...)
}
