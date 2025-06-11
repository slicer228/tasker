package app

import (
	"log/slog"
	"tasker/internal/config"
)

type App struct {
	AbstractApp
	logger *slog.Logger
	RestServer
}

func (app *App) MustRun() {

}

func (app *App) Stop() {

}

func NewApp(logger *slog.Logger, cfg *config.Config) *App {

}
