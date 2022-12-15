package handlers

import (
	"github.com/allensuvorov/tasker/internal/server/domain"
	"log"
	"net/http"
)

type TaskService interface {
	CreateTask(te domain.TaskEntity)
	GetTaskStatus(id string) domain.ResultEntity
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

	// TODO: gen TaskID
	// TODO: decode JSON body to struct
	// TODO: return JSON with TaskID and http status 200

	th.taskService.CreateTask(domain.TaskEntity{})

	log.Println("Handler CreateTask - bye")
}

func (th TaskHandler) GetTaskStatus(w http.ResponseWriter, r *http.Request) {
	//th.taskService.GetTaskStatus(taskID)
}
