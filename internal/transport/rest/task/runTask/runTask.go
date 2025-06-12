package runTask

import (
	"encoding/json"
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
	"net/http"
	task_repo "tasker/internal/service/task-repo"
)

type RunTaskRequest struct {
	TaskId int `json:"task_id"`
}

type ErrorResponse struct {
	Error  string `json:"error"`
	TaskId uint64 `json:"task_id,omitempty"`
}

func NewRunTaskRouter(log *slog.Logger, tasker *task_repo.TaskFarm) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error

		log.Info("New request for run task", "request_id", middleware.GetReqID(r.Context()))

		w.Header().Set("Content-Type", "application/json")

		var data RunTaskRequest
		err = json.NewDecoder(r.Body).Decode(&data)

		if err != nil {
			log.Error("Failed to decode request body", "error", err, "request_id", r.Header.Get("x-request-id"))
			errorResponse := &ErrorResponse{Error: "Invalid data"}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(errorResponse)
			return
		}

		err = tasker.RunTask(uint64(data.TaskId))

		if err != nil {
			log.Error("Task not exists", "error", err, "request_id", r.Header.Get("x-request-id"))
			errorResponse := &ErrorResponse{Error: "Task not exists", TaskId: uint64(data.TaskId)}
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(errorResponse)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
