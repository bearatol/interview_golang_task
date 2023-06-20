package logger

import (
	"fmt"
	"time"

	"github.com/getsentry/sentry-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Option func(*zap.SugaredLogger)

func NewLogger() (*zap.Logger, error) {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, err := config.Build()
	if err != nil {
		return nil, err
	}

	return logger, nil
}

func NewSugarLogger(logger *zap.Logger, opts ...Option) *zap.SugaredLogger {
	log := logger.Sugar()

	for _, opt := range opts {
		opt(log)
	}

	return log
}

func WithSentry(sentryHub *sentry.Hub) Option {
	return func(zapLogger *zap.SugaredLogger) {
		opts := zap.WrapCore(func(zapCore zapcore.Core) zapcore.Core {
			return zapcore.RegisterHooks(zapCore, func(zapEntry zapcore.Entry) error {
				defer sentryHub.Flush(2 * time.Second)

				switch zapEntry.Level {
				case zap.ErrorLevel:
					sentryHub.CaptureException(fmt.Errorf("%s, Line No: %d :: %s", zapEntry.Caller.File, zapEntry.Caller.Line, zapEntry.Message))
				}

				return nil
			})
		})

		*zapLogger = *(zapLogger.WithOptions(opts))
	}
}
