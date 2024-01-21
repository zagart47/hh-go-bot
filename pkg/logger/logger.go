package logger

import (
	"hh-go-bot/internal/config"
	"log/slog"
	"os"
)

const (
	envDebug   = "debug"
	envRelease = "release"
)

func NewLogger() *slog.Logger {
	env := config.All.LoggerMode
	var log *slog.Logger

	switch env {
	case envDebug:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelDebug,
			}),
		)
	case envRelease:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelWarn,
			}),
		)
	}
	return log
}

var MyLogger = NewLogger()

type Logger interface {
	Debug(string, ...any)
	Info(string, ...any)
	Error(string, ...any)
	Warn(string, ...any)
}

type Log struct {
	logger *slog.Logger
}

func (l Log) Debug(msg string, args ...any) {
	l.logger.Debug(msg, args)
}
func (l Log) Info(msg string, args ...any) {
	l.logger.Info(msg, args)
}
func (l Log) Error(msg string, args ...any) {
	l.logger.Error(msg, args)
}
func (l Log) Warn(msg string, args ...any) {
	l.logger.Warn(msg, args)
}
