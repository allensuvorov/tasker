package service

import (
	"github.com/allensuvorov/tasker/internal/domain"
	"log"
)

type TaskStorage interface {
	CreateTask(te domain.TaskEntity) error
	GetTaskStatus(taskID string) domain.ResultEntity
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

	return taskService.taskStorage.GetTaskStatus(taskID)
}
