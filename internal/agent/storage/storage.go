package storage

import (
	"database/sql"
	"github.com/allensuvorov/tasker/internal/server/domain"
)

type Storage struct {
	DB *sql.DB
}

func NewStorage() *Storage {

	return &Storage{
		//DB: db,
	}
}

func (s Storage) GetNewTasks() []domain.TaskEntity {
	return []domain.TaskEntity{}
}
func (s Storage) BulkUpdateTaskStatuses(map[string]int)               {}
func (s Storage) BulkCreateTaskResults(results []domain.ResultEntity) {}
