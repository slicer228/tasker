package rest

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
	task_repo "tasker/internal/service/task-repo"
	"tasker/internal/transport/rest/task"
	"time"
)

func NewHTTPServer(log *slog.Logger, timeout time.Duration, tasker *task_repo.TaskFarm) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Timeout(timeout), middleware.RequestID)
	task.MustLoadTaskRoutes(log, router, tasker)
	return router
}
