package storage

import (
	"database/sql"
	"github.com/allensuvorov/tasker/internal/server/domain/entity"
	"log"
)

type Storage struct {
	DB *sql.DB
}

func NewStorage() *Storage {

	return &Storage{
		//DB: db,
	}
}

func (s Storage) GetNewTasks() []entity.TaskEntity {
	log.Println("Storage.GetNewTasks - hello")

	log.Println("Storage.GetNewTasks - bye")
	return []entity.TaskEntity{}
}
func (s Storage) BulkUpdateTaskStatuses(map[string]int)               {}
func (s Storage) BulkCreateTaskResults(results []entity.ResultEntity) {}
