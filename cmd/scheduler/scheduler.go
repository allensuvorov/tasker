package scheduler

import (
	"github.com/allensuvorov/tasker/internal/domain"
)

func schedule(taskID string) {
	// TODO: task scheduler with goroutines
	// TODO: go func(){}() that assigns task to agent - taskAgent(taskID)
}

func taskAgent(taskID string) {
	// TODO:
	//receive taskID from <-ch
	//do the task - send request to 3rd party
	//update status of task in DB
}

type DoService interface {
	GetNewTasks() []domain.TaskEntity
	ScheduleTasks([]domain.TaskEntity) []domain.ResultEntity
	BulkCreateTaskResults(results []domain.ResultEntity)
}

type DoStorage interface {
	GetNewTasks() []domain.TaskEntity
	BulkUpdateTaskStatuses(map[string]int)
	BulkCreateTaskResults(results []domain.ResultEntity)
}

type DoRequest interface {
	Request(domain.TaskEntity) domain.ResultEntity
}

func main() {
	// TODO: check DB for new tasks
	// TODO: send tasks to 3rdparty
	// TODO: update DB
	// TODO: add layers for - storage, service, remote

	//DoStorage := NewDoStorage()
	//DoRemote := NewDoRemote()
	//DoService := NewDoService(DoRemote, DoStorage)
	//reCheckTime := 500*time.Millisecond
	//DoService.doStorage.GetNewTasks(reCheckTime)
}
