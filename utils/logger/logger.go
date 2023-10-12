// logger.go
package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

// InitializeLogger initializes the logger with the given log level.
func InitializeLogger(logLevel zapcore.Level) error {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	cfg.Level = zap.NewAtomicLevelAt(logLevel)

	// Customize the encoder to remove the JSON structure for error messages.
	cfg.EncoderConfig.EncodeName = zapcore.FullNameEncoder
	cfg.EncoderConfig.TimeKey = ""
	cfg.EncoderConfig.MessageKey = "message"

	var err error
	logger, err = cfg.Build()
	if err != nil {
		return err
	}

	return nil
}

// GetLogger returns the initialized logger.
func GetLogger() *zap.Logger {
	if logger == nil {
		// Initialize with a default log level of INFO if not initialized previously.
		_ = InitializeLogger(zapcore.InfoLevel)
	}
	return logger
}
