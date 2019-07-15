package user

import (
	"context"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// LogError logs a error to stderr including some internal information like
// error code and message.
func LogError(ctx context.Context, logger *zap.Logger, route string, msg string, err error, extraFields ...zap.Field) {
	var errCode string
	var errMsg string

	switch terr := err.(type) {
	case *Error:
		errCode = terr.ErrCode
		errMsg = terr.Message

	default:
		errMsg = terr.Error()
	}

	fields := []zap.Field{
		zap.String(string(ContextKeyTransactionID), GetTransactionID(ctx)),
		zap.String(string(ContextKeyInternalID), GetInternalID(ctx)),
		zap.String("route", route),
		zap.String("error-code", errCode),
		zap.String("error-message", errMsg),
	}
	if len(extraFields) > 0 {
		fields = append(fields, extraFields...)
	}

	logger.Error(msg, fields...)
}

// ConfigLog returns a zap configuration struct.
func ConfigLog(level zap.AtomicLevel) zap.Config {
	return zap.Config{
		Level:         level,
		Development:   false,
		DisableCaller: true,
		Sampling:      nil,
		Encoding:      "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: millisDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

func millisDurationEncoder(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendInt(int(float64(d) / float64(time.Millisecond)))
}
