package main

import (
	"github.com/allensuvorov/tasker/internal/server/remote/handlers"
	"github.com/allensuvorov/tasker/internal/server/remote/routers"
	"github.com/allensuvorov/tasker/internal/server/service"
	"github.com/allensuvorov/tasker/internal/server/storage"
	"log"
	"net/http"
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
