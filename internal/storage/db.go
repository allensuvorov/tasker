package storage

import (
	"github.com/allensuvorov/tasker/internal/domain"
	"log"
)

type TaskStorage struct {
}

func NewTaskStorage() *TaskStorage {
	return &TaskStorage{}
}

func (taskStorage TaskStorage) CreateTask(entity domain.TaskEntity) error {
	log.Println("Storage CreateTask - hello")

	log.Println("Storage CreateTask - bye")

	return nil
}
