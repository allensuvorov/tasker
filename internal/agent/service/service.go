package service

import (
	"github.com/allensuvorov/tasker/internal/server/domain/entity"
	"time"
)

type Storage interface {
	GetNewTasks() []entity.TaskEntity
	BulkUpdateTaskStatuses(map[string]int)
	BulkCreateTaskResults(results []entity.ResultEntity)
}

type Request interface {
	Request(entity.TaskEntity) entity.ResultEntity
}

type Service struct {
	storage Storage
	request Request
}

func NewService(s Storage, r Request) Service {
	return Service{
		storage: s,
		request: r,
	}
}

func (s Service) GetNewTasks(timeInterval time.Duration) []entity.TaskEntity {
	return []entity.TaskEntity{}
}
func (s Service) DoTasks([]entity.TaskEntity) []entity.ResultEntity {
	return []entity.ResultEntity{}
}
func (s Service) BulkCreateTaskResults(results []entity.ResultEntity) {}
