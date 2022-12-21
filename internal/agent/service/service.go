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
	BulkAddTaskResultsViaCh(resultCh chan entity.ResultEntity) error
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

	s.createWorkers(1000)

	for i := 0; i < 1; i++ {

		err := s.getNewTasks()
		if err != nil {
			return err
		}

		time.Sleep(1 * timeInterval)
	}
	log.Println("Service.StartGettingNewTasks - bye")
	return nil
}

var taskCh = make(chan entity.TaskEntity)
var resultCh = make(chan entity.ResultEntity, 1000)

// TODO need channel here
var results []entity.ResultEntity

func (s Service) createWorkers(n uint) {
	for i := 0; i < 1000; i++ {
		go func() {
			for task := range taskCh {
				s.makeRequest(task)
			}
		}()
	}
}

func (s Service) makeRequest(task entity.TaskEntity) {
	result, err := s.request.Request(task)
	if err != nil {
		log.Println(err)
	}
	select {
	case resultCh <- result:
	default:

	}
	resultCh <- result
	//select {
	//case db.BufferCh <- v:
	//default:
	//	db.flushBufferToDB()
	//}

	//results = append(results, result)
	//if len(results) == 1000 {
	//	s.bulkCreateTaskResults(results)
	//	results = nil
	//}
}

func (s Service) getNewTasks() error {
	startTimer := time.Now()
	log.Println("Service.getNewTasks - hello")

	newTasks, err := s.storage.GetNewTasks()
	if err != nil {
		return err
	}

	if len(newTasks) != 0 {
		for _, task := range newTasks {
			taskCh <- task
		}
	}

	log.Println("Service.getNewTasks - bye")
	duration := time.Since(startTimer)
	fmt.Printf("Servis.getNewTasks - Execution Time ms %d\n", duration.Milliseconds())
	return nil
}

func (s Service) bulkCreateTaskResults(results []entity.ResultEntity) {
	err := s.storage.BulkAddTaskResults(results)
	if err != nil {
		log.Println(err)
	}
}

//func (s Service) schedule(newTasks []entity.TaskEntity) {
//	var results []entity.ResultEntity
//	// TODO go func add - fan-out
//	// TODO for very long lists such as 10^10, let's process lists concurrently
//
//	for _, task := range newTasks {
//
//		// go func
//		result, err := s.request.Request(task)
//		if err != nil {
//			log.Println(err)
//		}
//		results = append(results, result)
//		// results to ch buffer - fan-in
//	}
//
//	s.bulkCreateTaskResults(results)
//}
