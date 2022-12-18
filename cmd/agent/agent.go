package main

import (
	"log"
	"time"

	"github.com/allensuvorov/tasker/internal/agent/remote"
	"github.com/allensuvorov/tasker/internal/agent/service"
	"github.com/allensuvorov/tasker/internal/agent/storage"
)

func main() {
	Storage := storage.NewStorage()
	Request := remote.NewRequest()
	Service := service.NewService(Storage, Request)
	newTaskCheckInterval := 500 * time.Millisecond

	log.Println("Getting new tasks from storage...")
	Service.StartGettingNewTasks(newTaskCheckInterval)
}
