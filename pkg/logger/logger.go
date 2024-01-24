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

func newLogger() *slog.Logger {
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

var Log = newLogger()
