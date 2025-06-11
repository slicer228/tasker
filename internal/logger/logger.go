package logger

import (
	"log/slog"
	"os"
)

const (
	Prod  = "prod"
	Local = "local"
)

func NewLogger(env string) *slog.Logger {
	switch env {
	case Prod:
		return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	case Local:
		return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	default:
		panic("Invalid env for logger")
	}
}
