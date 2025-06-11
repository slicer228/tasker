package app

import (
	"github.com/go-chi/chi/v5"
	"log"
	"log/slog"
	"net/http"
	"tasker/internal/config"
	"tasker/internal/transport/rest"
)

type App struct {
	AbstractApp
	log        *slog.Logger
	restServer *chi.Mux
	cfg        *config.Config
}

func (app *App) MustRun() {
	err := http.ListenAndServe(app.cfg.Address, app.restServer)
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}

func (app *App) Stop() {

}

func NewApp(log *slog.Logger, cfg *config.Config) *App {
	return &App{
		log:        log,
		restServer: rest.NewHTTPServer(log, cfg.Timeout),
		cfg:        cfg,
	}
}
