package agent

import (
	"github.com/allensuvorov/tasker/internal/server/domain"
)

func schedule(taskID string) {
	// TODO: task agent with goroutines
	// TODO: go func(){}() that assigns task to agent - taskAgent(taskID)
}

func taskAgent(taskID string) {
	// TODO:
	//receive taskID from <-ch
	//do the task - send request to 3rd party
	//update status of task in DB
}

type Storage interface {
	GetNewTasks() []domain.TaskEntity
	BulkUpdateTaskStatuses(map[string]int)
	BulkCreateTaskResults(results []domain.ResultEntity)
}

type Request interface {
	Request(domain.TaskEntity) domain.ResultEntity
}

func main() {
	// TODO: check DB for new tasks
	// TODO: send tasks to 3rdparty
	// TODO: update DB
	// TODO: add layers for - storage, service, remote

	//Storage := NewDoStorage()
	//DoRemote := NewDoRemote()
	//Service := NewDoService(DoRemote, Storage)
	//reCheckTime := 500*time.Millisecond
	//Service.doStorage.GetNewTasks(reCheckTime)
}
