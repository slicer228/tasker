package createTask

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
	"net/http"
	"tasker/internal/service/lib/clock"
	task_repo "tasker/internal/service/task-repo"
	"time"
)

type CreateTaskResponse struct {
	TaskId uint64 `json:"task_id"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func NewCreateTaskRouter(log *slog.Logger, tasker *task_repo.TaskFarm) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error

		log.Info("New request for create task", "request_id", middleware.GetReqID(r.Context()))

		w.Header().Set("Content-Type", "application/json")

		//example of func to launch
		//you can use context to provide termination of this func(method StopTask in task-repo)
		taskId, err := tasker.CreateTask(clock.NewTimer(time.UTC), func(ctx *context.Context, args ...any) (any, error) {
			time.Sleep(time.Second * 5)
			return 1, nil
		})

		if err != nil {
			log.Error("Failed to create task", "error", err, "request_id", r.Header.Get("x-request-id"))
			errResponse := ErrorResponse{"Max tasks exceeded"}
			w.WriteHeader(http.StatusNotAcceptable)
			json.NewEncoder(w).Encode(errResponse)
			return
		}

		resp := CreateTaskResponse{TaskId: taskId}
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(resp)

		if err != nil {
			log.Error("Failed to respond to create task request", "error", err, "request_id", r.Header.Get("x-request-id"))
			errResponse := ErrorResponse{"Failed to respond to create task request"}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(errResponse)
			return
		}
	}
}
