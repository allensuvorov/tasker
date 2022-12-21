package service

import (
	"fmt"
	"github.com/allensuvorov/tasker/internal/server/domain/entity"
	"log"
	"runtime"
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

var taskCh = make(chan entity.TaskEntity)

var resultCh = make(chan entity.ResultEntity, 5)

func (s Service) createWorkers(n int) {
	for i := 0; i < n; i++ {
		go func() {
			for task := range taskCh {
				s.makeRequest(task)
			}
		}()
	}
}
func (s Service) makeRequest(task entity.TaskEntity) {
	log.Println("Service.makeRequest - Hello")

	result, err := s.request.Request(task)
	if err != nil {
		log.Println(err)
	}
	select {
	case resultCh <- result:
		log.Println("taskCh len is ", len(taskCh))
		log.Println("NumGoroutine len is ", runtime.NumGoroutine())
	default:
		s.flushToDB()
		resultCh <- result
	}

	log.Println("Service.makeRequest - Bye")
}

func (s Service) StartGettingNewTasks(timeInterval time.Duration) error {
	log.Println("Service.StartGettingNewTasks - hello")
	startTimer := time.Now()

	s.createWorkers(1000)

	for i := 0; i < 10; i++ {

		err := s.getNewTasks()
		if err != nil {
			return err
		}

		time.Sleep(1 * timeInterval)
		log.Println("resultCh - len is ", len(resultCh))

		if len(resultCh) > 0 {
			s.flushToDB()
		}

		duration := time.Since(startTimer)
		fmt.Printf("Servis.getNewTasks - Execution Time ms %d\n", duration.Milliseconds())
	}
	log.Println("Service.StartGettingNewTasks - bye")
	return nil
}

func (s Service) flushToDB() {
	log.Println("service.flushToDB - flushing resultCh - len is ", len(resultCh))
	err := s.storage.BulkAddTaskResultsViaCh(resultCh)
	if err != nil {
		log.Println(err)
	}
}

func (s Service) getNewTasks() error {
	log.Println("Service.getNewTasks - hello")
	newTasks, err := s.storage.GetNewTasks()

	//maxLen := 1000
	//if len(newTasks) < maxLen {
	//	resultCh = make(chan entity.ResultEntity, len(newTasks))
	//}

	if err != nil {
		return err
	}

	if len(newTasks) != 0 {
		for _, task := range newTasks {
			taskCh <- task
		}
	}

	log.Println("Service.getNewTasks - bye")

	return nil
}

// TODO need channel here
//var results []entity.ResultEntity

//func (s Service) bulkAddTaskResults(results []entity.ResultEntity) {
//	err := s.storage.BulkAddTaskResults(results)
//	if err != nil {
//		log.Println(err)
//	}
//}

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
//	s.bulkAddTaskResults(results)
//}
