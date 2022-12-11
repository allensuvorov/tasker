package scheduler

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

func main() {
	schedule("taskID")
}
