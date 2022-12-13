package service

import (
	"log"

	"github.com/allensuvorov/tasker/internal/server/domain"
)

type TaskStorage interface {
	CreateTask(te domain.TaskEntity) error
	GetTaskStatus(taskID string) (domain.ResultEntity, error)
}

type TaskService struct {
	taskStorage TaskStorage
}

func NewTaskService(taskStorage TaskStorage) TaskService {
	return TaskService{
		taskStorage: taskStorage,
	}
}

func (taskService TaskService) CreateTask(te domain.TaskEntity) {
	log.Println("Service CreateTask - hello")

	err := taskService.taskStorage.CreateTask(te)
	if err != nil {
		log.Println(err)
	}

	log.Println("Service CreateTask - bye")
}

func (taskService TaskService) GetTaskStatus(taskID string) domain.ResultEntity {
	log.Println("Service GetTaskStatus - hello")

	log.Println("Service GetTaskStatus - bye")

	re, err := taskService.taskStorage.GetTaskStatus(taskID)
	if err != nil {
		log.Println(err)
	}
	return re
}
