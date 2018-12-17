package logger

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(config zap.Config) *zap.Logger {
	config.EncoderConfig.LineEnding     = zapcore.DefaultLineEnding
	config.EncoderConfig.EncodeLevel    = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.EncodeTime     = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.EncodeDuration = zapcore.StringDurationEncoder
	config.EncoderConfig.EncodeCaller   = zapcore.ShortCallerEncoder

	logger,err := config.Build()
	if err != nil {
		panic(err);
	}
	return logger
}

func GetLogger(ctx context.Context) *zap.Logger  {
	return ctx.Value("logger").(*zap.Logger)
}

func SetLogger(ctx context.Context,logger *zap.Logger) context.Context {
	return context.WithValue(ctx,"logger",logger)
}