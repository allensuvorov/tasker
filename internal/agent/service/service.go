package service

import (
	"fmt"
	"github.com/allensuvorov/tasker/internal/server/domain/entity"
	"log"
	"time"
)

type Storage interface {
	GetNewTasks() ([]entity.TaskEntity, error)
	BulkAddTaskResults(results []entity.ResultEntity) error
}

type Request interface {
	Request(taskEntity entity.TaskEntity) (entity.ResultEntity, error)
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
		startTimer := time.Now()

		err := s.getNewTasks()
		if err != nil {
			return err
		}
		duration := time.Since(startTimer)
		fmt.Printf("Servis.StartGettingNewTasks - Execution Time ms %d\n", duration.Milliseconds())

		time.Sleep(1 * timeInterval)
	}
	log.Println("Service.StartGettingNewTasks - bye")
	return nil
}

func (s Service) getNewTasks() error {
	log.Println("Service.getNewTasks - hello")

	newTasks, err := s.storage.GetNewTasks()
	if err != nil {
		return err
	}

	if len(newTasks) != 0 {
		s.schedule(newTasks)
	}

	log.Println("Service.getNewTasks - bye")

	return nil
}

func (s Service) schedule(newTasks []entity.TaskEntity) {
	var results []entity.ResultEntity
	// TODO go func add - fan-out

	for _, task := range newTasks {

		// go func
		result, err := s.request.Request(task)
		if err != nil {
			log.Println(err)
		}
		results = append(results, result)
		// results to ch buffer - fan-in
	}

	s.bulkCreateTaskResults(results)
}

func (s Service) bulkCreateTaskResults(results []entity.ResultEntity) {
	s.storage.BulkAddTaskResults(results)
}
