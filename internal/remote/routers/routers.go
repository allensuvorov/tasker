package routers

import (
	"github.com/go-chi/chi/v5"

	"github.com/allensuvorov/tasker/internal/remote/handlers"
)

func NewRouter(task handlers.TaskHandler) chi.Router {

	r := chi.NewRouter()

	r.Post("/task", task.CreateTask)
	r.Get("/task/{taskID}", task.GetStatus)
	return r
}
