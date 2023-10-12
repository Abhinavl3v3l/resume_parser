// log.go
package logger

import (
	"go.uber.org/zap"
)

// logWithStruct logs a message with optional key-value pairs and a struct.
func logWithStruct(logFn func(string, ...zap.Field), message string, err error, args ...interface{}) {
	fields := make([]zap.Field, 0, len(args)/2)

	for i := 0; i < len(args)-1; i += 2 {
		key, ok := args[i].(string)
		if !ok {
			key = "unknown"
		}

		fields = append(fields, zap.Any(key, args[i+1]))
	}

	if err != nil {
		fields = append(fields, zap.Error(err))
	}

	logFn(message, fields...)
}

// Info logs an informational message with optional key-value pairs and a struct.
func Info(message string, args ...interface{}) {
	logWithStruct(GetLogger().Info, message, nil, args...)
}

// Debug logs a debug message with optional key-value pairs and a struct.
func Debug(message string, args ...interface{}) {
	logWithStruct(GetLogger().Debug, message, nil, args...)
}

// Warn logs a warning message with optional key-value pairs and a struct.
func Warn(message string, args ...interface{}) {
	logWithStruct(GetLogger().Warn, message, nil, args...)
}

// Error logs an error message without the JSON structure with optional key-value pairs and a struct.
func Error(message string, args ...interface{}) {
	var err error
	var fields []zap.Field

	for i := 0; i < len(args); i++ {
		arg := args[i]
		switch arg := arg.(type) {
		case error:
			err = arg
		default:
			if i+1 < len(args) {
				key, ok := args[i].(string)
				if !ok {
					key = "unknown"
				}
				value := args[i+1]
				fields = append(fields, zap.Any(key, value))
				i++
			}
		}
	}

	if err != nil {
		fields = append(fields, zap.Error(err))
	}

	GetLogger().Error(message, fields...)
}
