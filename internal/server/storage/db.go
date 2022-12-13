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
	ID serial PRIMARY KEY,
	task_id text,
	task_request_method text,
	task_url text,
	task_headers json,
	task_status text DEFAULT 'new',
	result_http_status_code text,
	result_headers json,
	result_body_length INTEGER
	                         );`)
	if err != nil {
		log.Fatal(err)
	}

	return &TaskStorage{
		DB: db,
	}
}

func (ts TaskStorage) CreateTask(te domain.TaskEntity) error {
	log.Println("Storage CreateTask - hello")

	_, err := ts.DB.Exec(
		`INSERT INTO tasks
		(task_id , task_request_method , task_url, task_headers)
		VALUES
		($1, $2, $3, $4);`,
		te.ID, te.Method, te.URL, te.Headers,
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Storage CreateTask - bye")

	return nil
}

func (ts TaskStorage) GetTaskStatus(taskID string) domain.ResultEntity {
	log.Println("Storage GetTaskStatus - hello")

	log.Println("Storage GetTaskStatus - bye")

	return domain.ResultEntity{}
}
