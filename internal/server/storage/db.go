package storage

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/allensuvorov/tasker/internal/server/domain"
)

type TaskStorage struct {
	DB *sql.DB
}

func NewTaskStorage() *TaskStorage {
	db, err := sql.Open("pgx",
		"postgres://postgres:sql@localhost:5432/postgres")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS tasks(
	ID SERIAL PRIMARY KEY,
	task_id TEXT,
	request_method TEXT,
	url TEXT,
	headers JSON,
	status TEXT
	                         );`)
	if err != nil {
		log.Fatal(err)
	}

	return &TaskStorage{
		DB: db,
	}
}

func (taskStorage TaskStorage) CreateTask(entity domain.TaskEntity) error {
	log.Println("Storage CreateTask - hello")

	log.Println("Storage CreateTask - bye")

	return nil
}

func (taskStorage TaskStorage) GetTaskStatus(taskID string) domain.ResultEntity {
	log.Println("Storage GetTaskStatus - hello")

	log.Println("Storage GetTaskStatus - bye")

	return domain.ResultEntity{}
}
