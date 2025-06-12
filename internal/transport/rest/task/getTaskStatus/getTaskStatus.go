package getTaskStatus

import (
	"encoding/json"
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
	"net/http"
	"strconv"
	task_repo "tasker/internal/service/task-repo"
)

type ErrorResponse struct {
	Error  string `json:"error"`
	TaskId uint64 `json:"task_id,omitempty"`
}

func NewGetTaskInfoRouter(log *slog.Logger, tasker *task_repo.TaskFarm) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error

		log.Info("New request for task_info", "request_id", middleware.GetReqID(r.Context()))

		w.Header().Set("Content-Type", "application/json")

		taskId, err := strconv.ParseUint(r.URL.Query().Get("task_id"), 10, 64)

		if err != nil {
			log.Error("Invalid parameters", "error", err, "request_id", r.Header.Get("x-request-id"))
			errorResponse := &ErrorResponse{Error: "Invalid parameters"}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(errorResponse)
			return
		}

		res, err := tasker.GetTaskInfo(taskId)

		if err != nil {
			log.Error("Task not found", "error", err, "request_id", r.Header.Get("x-request-id"))
			errorResponse := &ErrorResponse{Error: "Task not found"}
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(errorResponse)
			return
		}

		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(res)

		if err != nil {
			log.Error("Failed to respond to send task info", "error", err, "request_id", r.Header.Get("x-request-id"))
			errorResponse := &ErrorResponse{Error: "Failed to respond to send task info", TaskId: taskId}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(errorResponse)
			return
		}

	}
}
