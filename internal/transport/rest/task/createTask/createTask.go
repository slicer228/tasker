package createTask

import (
	"log/slog"
	"net/http"
)

func NewCreateTaskRouter(log *slog.Logger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello World"))
		if err != nil {
			log.Error("Failed to respond to create task request", "error", err)
		}
	}
}
