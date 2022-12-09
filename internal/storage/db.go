package storage

import "github.com/allensuvorov/tasker/internal/domain"

type TaskStorage struct {
}

func NewTaskStorage() *TaskStorage {
	return &TaskStorage{}
}

func (taskStorage TaskStorage) CreateTask(entity domain.TaskEntity) error {
	return nil
}
