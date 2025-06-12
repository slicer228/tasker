package app

import (
	"github.com/go-chi/chi/v5"
	"log"
	"log/slog"
	"net/http"
	"tasker/internal/config"
	task_repo "tasker/internal/service/task-repo"
	"tasker/internal/transport/rest"
)

type App struct {
	AbstractApp
	log        *slog.Logger
	restServer *chi.Mux
	tasker     *task_repo.TaskFarm
	cfg        *config.Config
}

func (app *App) MustRun() {
	app.log.Info("starting app...", "address", app.cfg.Address)
	err := http.ListenAndServe(app.cfg.Address, app.restServer)
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}

func (app *App) Stop() {

}

func NewApp(log *slog.Logger, cfg *config.Config) *App {
	tasker := task_repo.NewTaskFarm(log, cfg.MaxTasks)
	return &App{
		log:        log,
		restServer: rest.NewHTTPServer(log, cfg.Timeout, tasker),
		cfg:        cfg,
		tasker:     tasker,
	}
}
