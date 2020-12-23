package logger

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/batazor/shortlink/internal/pkg/logger/field"
	"github.com/batazor/shortlink/internal/pkg/logger/tracer"
)

type zapLogger struct { // nolint unused
	logger *zap.Logger
}

func (log *zapLogger) init(config Configuration) error {
	logLevel := log.setLogLevel(config.Level)

	// To keep the example deterministic, disable timestamps in the output.
	encoderCfg := zapcore.EncoderConfig{
		TimeKey:        "@timestamp",
		LevelKey:       "@level",
		NameKey:        "logger",
		CallerKey:      "@caller",
		MessageKey:     "@msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.TimeEncoder(log.timeEncoder(config.TimeFormat)),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	log.logger = zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.Lock(zapcore.AddSync(config.Writer)),
		logLevel,
	), zap.AddCaller(), zap.AddCallerSkip(1))

	return nil
}

func (log *zapLogger) Close() error {
	err := log.logger.Sync()
	return err
}

func (log *zapLogger) SetConfig(config Configuration) error {
	var err error
	logLevel := log.setLogLevel(config.Level)

	cfg := zap.Config{
		Level: logLevel,
	}
	if log.logger, err = cfg.Build(); err != nil {
		return err
	}

	return nil
}

func (log *zapLogger) converter(fields ...field.Fields) []zap.Field {
	var zapFields []zap.Field

	for _, field := range fields {
		for k, v := range field {
			zapFields = append(zapFields, zap.Any(k, v))
		}
	}

	return zapFields
}

func (log *zapLogger) setLogLevel(logLevel int) zap.AtomicLevel {
	atom := zap.NewAtomicLevel()

	switch logLevel {
	case FATAL_LEVEL:
		atom.SetLevel(zap.FatalLevel)
	case ERROR_LEVEL:
		atom.SetLevel(zap.ErrorLevel)
	case WARN_LEVEL:
		atom.SetLevel(zap.WarnLevel)
	case INFO_LEVEL:
		atom.SetLevel(zap.InfoLevel)
	case DEBUG_LEVEL:
		atom.SetLevel(zap.DebugLevel)
	default:
		atom.SetLevel(zap.InfoLevel)
	}

	return atom
}

func (log *zapLogger) timeEncoder(format string) func(time.Time, zapcore.PrimitiveArrayEncoder) {
	return func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format(format))
	}
}

// Fatal ===============================================================================================================

func (log *zapLogger) Fatal(msg string, fields ...field.Fields) {
	zapFields := log.converter(fields...)
	log.logger.Fatal(msg, zapFields...)
}

func (log *zapLogger) FatalWithContext(ctx context.Context, msg string, fields ...field.Fields) {
	fields, err := tracer.NewTraceFromContext(ctx, msg, fields...)
	if err != nil {
		log.logger.Error(fmt.Sprintf("Error send span to opentracing: %s", err.Error()))
	}

	zapFields := log.converter(fields...)
	log.logger.Fatal(msg, zapFields...)
}

// Warn ================================================================================================================

func (log *zapLogger) Warn(msg string, fields ...field.Fields) {
	zapFields := log.converter(fields...)
	log.logger.Warn(msg, zapFields...)
}

func (log *zapLogger) WarnWithContext(ctx context.Context, msg string, fields ...field.Fields) {
	fields, err := tracer.NewTraceFromContext(ctx, msg, fields...)
	if err != nil {
		log.logger.Error(fmt.Sprintf("Error send span to opentracing: %s", err.Error()))
	}

	zapFields := log.converter(fields...)
	log.logger.Warn(msg, zapFields...)
}

// Error ===============================================================================================================

func (log *zapLogger) Error(msg string, fields ...field.Fields) {
	zapFields := log.converter(fields...)
	log.logger.Error(msg, zapFields...)
}

func (log *zapLogger) ErrorWithContext(ctx context.Context, msg string, fields ...field.Fields) {
	fields, err := tracer.NewTraceFromContext(ctx, msg, fields...)
	if err != nil {
		log.logger.Error(fmt.Sprintf("Error send span to opentracing: %s", err.Error()))
	}

	zapFields := log.converter(fields...)
	log.logger.Error(msg, zapFields...)
}

// Info ================================================================================================================

func (log *zapLogger) Info(msg string, fields ...field.Fields) {
	zapFields := log.converter(fields...)
	log.logger.Info(msg, zapFields...)
}

func (log *zapLogger) InfoWithContext(ctx context.Context, msg string, fields ...field.Fields) {
	fields, err := tracer.NewTraceFromContext(ctx, msg, fields...)
	if err != nil {
		log.logger.Error(fmt.Sprintf("Error send span to opentracing: %s", err.Error()))
	}

	zapFields := log.converter(fields...)
	log.logger.Info(msg, zapFields...)
}

// Debug ===============================================================================================================

func (log *zapLogger) Debug(msg string, fields ...field.Fields) {
	zapFields := log.converter(fields...)
	log.logger.Debug(msg, zapFields...)
}

func (log *zapLogger) DebugWithContext(ctx context.Context, msg string, fields ...field.Fields) {
	fields, err := tracer.NewTraceFromContext(ctx, msg, fields...)
	if err != nil {
		log.logger.Error(fmt.Sprintf("Error send span to opentracing: %s", err.Error()))
	}

	zapFields := log.converter(fields...)
	log.logger.Debug(msg, zapFields...)
}
