package service

import (
	"github.com/allensuvorov/tasker/internal/server/domain"
	"time"
)

type Storage interface {
	GetNewTasks() []domain.TaskEntity
	BulkUpdateTaskStatuses(map[string]int)
	BulkCreateTaskResults(results []domain.ResultEntity)
}

type Request interface {
	Request(domain.TaskEntity) domain.ResultEntity
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

func (s Service) GetNewTasks(timeInterval time.Duration) []domain.TaskEntity {
	return []domain.TaskEntity{}
}
func (s Service) DoTasks([]domain.TaskEntity) []domain.ResultEntity {
	return []domain.ResultEntity{}
}
func (s Service) BulkCreateTaskResults(results []domain.ResultEntity) {}
