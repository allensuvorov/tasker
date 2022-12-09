package service

import "github.com/allensuvorov/tasker/internal/domain"

type TaskStorage interface {
}

type TaskService struct {
	taskStorage TaskStorage
}

func NewTaskService(taskStorage TaskStorage) TaskService {
	return TaskService{
		taskStorage: taskStorage,
	}
}

func (taskService TaskService) CreateTask(entity domain.TaskEntity) {

}
