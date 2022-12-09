package handlers

import (
	"log"
	"net/http"

	"github.com/allensuvorov/tasker/internal/domain"
)

type TaskService interface {
	CreateTask(te domain.TaskEntity)
}

type TaskHandler struct {
	taskService TaskService
}

func NewTaskHandler(ts TaskService) TaskHandler {
	return TaskHandler{
		taskService: ts,
	}
}

func (th TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	log.Println("Handler CreateTask - hello")

	// TODO: gen ID
	// TODO: decode JSON body to struct
	// TODO: return JSON with ID and http status 200

	th.taskService.CreateTask(domain.TaskEntity{})

	log.Println("Handler CreateTask - bye")
}

func (th TaskHandler) GetStatus(w http.ResponseWriter, r *http.Request) {

}
