package rest

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
	"tasker/internal/transport/rest/task"
	"time"
)

func NewHTTPServer(log *slog.Logger, timeout time.Duration) *chi.Mux {
	router := chi.NewRouter()
	router.With(middleware.Timeout(timeout), middleware.RequestID)
	task.MustLoadTaskRoutes(log, router)
	return router
}
