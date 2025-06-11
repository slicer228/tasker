package task

import (
	"github.com/go-chi/chi/v5"
	"log/slog"
	"tasker/internal/transport/rest/task/createTask"
)

func MustLoadTaskRoutes(log *slog.Logger, r *chi.Mux) {
	r.Route("/task", func(r chi.Router) {
		r.Post("/", createTask.NewCreateTaskRouter(log))
	})
}
