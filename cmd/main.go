package main

import (
	"log"
	"net/http"

	"github.com/allensuvorov/tasker/internal/remote/handlers"
	"github.com/allensuvorov/tasker/internal/remote/routers"
	"github.com/allensuvorov/tasker/internal/service"
	"github.com/allensuvorov/tasker/internal/storage"
)

func main() {
	TaskStorage := storage.NewTaskStorage()
	TaskService := service.NewTaskService(TaskStorage)
	TaskHandler := handlers.NewTaskHandler(TaskService)
	r := routers.NewRouter(TaskHandler)
	serverAddress := ":8080"
	log.Println("Serving on port", serverAddress)
	log.Fatal(http.ListenAndServe(serverAddress, r))
}
