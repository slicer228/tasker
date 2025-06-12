package task

import (
	"github.com/go-chi/chi/v5"
	"log/slog"
	task_repo "tasker/internal/service/task-repo"
	"tasker/internal/transport/rest/task/createTask"
	"tasker/internal/transport/rest/task/deleteTask"
	"tasker/internal/transport/rest/task/getTaskStatus"
	"tasker/internal/transport/rest/task/runTask"
)

func MustLoadTaskRoutes(log *slog.Logger, r *chi.Mux, tasker *task_repo.TaskFarm) {
	r.Route("/task", func(r chi.Router) {
		r.Post("/", createTask.NewCreateTaskRouter(log, tasker))
		r.Post("/run", runTask.NewRunTaskRouter(log, tasker))
		r.Get("/", getTaskStatus.NewGetTaskInfoRouter(log, tasker))
		r.Delete("/", deleteTask.NewDeleteTaskRouter(log, tasker))
	})
}
