package storage

import (
	"database/sql"
	"log"

	"github.com/allensuvorov/tasker/internal/domain"
)

type TaskStorage struct {
	DB *sql.DB
}

func NewTaskStorage() *TaskStorage {
	//db, err := sql.Open("pgx",
	//	"postgres://postgres:sql@localhost:5432/postgres")
	//if err != nil {
	//	panic(err)
	//}

	//
	//_, err = db.Exec(`CREATE TABLE IF NOT EXISTS tasks(
	//ID SERIAL PRIMARY KEY,
	//method TEXT,
	//hash TEXT,
	//client TEXT,
	//deleted BOOL DEFAULT 'false'
	//                          );`)
	//if err != nil {
	//	log.Fatal(err)
	//}

	return &TaskStorage{
		//DB: db,
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
