package user

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

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
