// Provides a custom logger over the slog package.
package logger

import (
	"cdvet/app/config"
	"fmt"
	"log/slog"
)

type Logger interface {
	Debug(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}

type Log struct {
	*slog.Logger
}

func (l *Log) Debug(format string, args ...interface{}) {
	l.Logger.Debug(format, args...)
}

func (l *Log) Info(format string, args ...interface{}) {
	l.Logger.Info(format, args...)
}

func (l *Log) Warn(format string, args ...interface{}) {
	l.Logger.Warn(format, args...)
}

func (l *Log) Error(format string, args ...interface{}) {
	l.Logger.Error(format, args...)
}

func (l *Log) Debugf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.Logger.Debug(msg)
}

func (l *Log) Infof(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.Logger.Info(msg)
}

func (l *Log) Warnf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.Logger.Warn(msg)
}

func (l *Log) Errorf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.Logger.Error(msg)
}

func New(cfg *config.LoggingConfig) Logger {
	opts := &slog.HandlerOptions{
		Level: getLogLevelFromString(cfg.Level),
	}

	slog.SetLogLoggerLevel(opts.Level.Level())

	logger := &Log{
		Logger: slog.Default(),
	}

	return logger
}

func getLogLevelFromString(level string) slog.Level {
	switch level {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	}
	slog.Warn("Invalid log level. Defaulting to info", "level", level)
	return slog.LevelInfo
}
