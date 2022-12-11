package service

import (
	"github.com/allensuvorov/tasker/cmd/agent"
	"github.com/allensuvorov/tasker/internal/server/domain"
)

type Service struct {
	as agent.Storage
	ar agent.Request
}

func NewService(as agent.Storage, ar agent.Request) Service {
	return Service{
		as: as,
		ar: ar,
	}
}

func GetNewTasks() []domain.TaskEntity {
	return []domain.TaskEntity{}
}
func DoTasks([]domain.TaskEntity) []domain.ResultEntity {
	return []domain.ResultEntity{}
}
func BulkCreateTaskResults(results []domain.ResultEntity) {}
