package handlers

import (
	"github.com/allensuvorov/tasker/internal/server/domain/entity"
	localError "github.com/allensuvorov/tasker/internal/server/domain/error"
	"log"
	"net/http"
)

type TaskService interface {
	CreateTask(te entity.TaskEntity)
	GetTaskStatus(id string) entity.ResultEntity
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

	// TODO: decode JSON body to struct
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, localError.ErrWrongContentType.Error(), http.StatusBadRequest)
		return
	}
	
	// TODO: gen TaskID
	// TODO: return JSON with TaskID and http status 200

	th.taskService.CreateTask(entity.TaskEntity{})

	log.Println("Handler CreateTask - bye")
}

func (th TaskHandler) GetTaskStatus(w http.ResponseWriter, r *http.Request) {
	//th.taskService.GetTaskStatus(taskID)
}
