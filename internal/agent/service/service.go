package service

import (
	"github.com/allensuvorov/tasker/internal/server/domain/entity"
	"log"
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

func (s Service) StartGettingNewTasks(timeInterval time.Duration) error {
	log.Println("Service.StartGettingNewTasks - hello")

	for i := 0; i < 2; i++ {
		// TODO infinite loop calls this function, gets data from DB and updates statuses
		s.GetNewTasks()
		time.Sleep(1 * timeInterval)
	}
	log.Println("Service.StartGettingNewTasks - bye")
	return nil
}

func (s Service) GetNewTasks() ([]entity.TaskEntity, error) {
	log.Println("Service.GetNewTasks - hello")

	newTasks := s.storage.GetNewTasks()
	log.Println("Service.GetNewTasks - bye")

	return newTasks, nil
}

func (s Service) DoTasks([]entity.TaskEntity) []entity.ResultEntity {
	return []entity.ResultEntity{}
}
func (s Service) BulkCreateTaskResults(results []entity.ResultEntity) {}
