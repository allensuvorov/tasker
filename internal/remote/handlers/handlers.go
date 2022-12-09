package handlers

import (
	"log"
	"net/http"
)

type TaskService interface {
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
	log.Println("Handler CreateTask - bye")
}

func (th TaskHandler) GetStatus(w http.ResponseWriter, r *http.Request) {

}
