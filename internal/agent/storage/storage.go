package storage

import (
	"database/sql"
	"encoding/json"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/allensuvorov/tasker/internal/server/domain/entity"
)

type Storage struct {
	DB *sql.DB
}

func NewStorage() *Storage {
	db, err := sql.Open("pgx",
		"postgres://postgres:sql@localhost:5432/postgres")
	if err != nil {
		panic(err)
	}
	return &Storage{
		DB: db,
	}
}

func (s *Storage) GetNewTasks() []entity.TaskEntity {
	log.Println("Storage.GetNewTasks - hello")

	rows, err := s.DB.Query(`SELECT task_id, task_request_method, task_headers, task_url
		FROM tasks WHERE task_status = $1;`, "new")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	if rows.Err() != nil {
		log.Fatal(err)
	}
	var newTasks []entity.TaskEntity
	for rows.Next() {
		task := entity.TaskEntity{}
		var responseHeadersBuffer []byte
		err := rows.Scan(&task.ID, &task.Method, &responseHeadersBuffer, &task.URL)
		if err != nil {
			log.Println(err)
		}
		err = json.Unmarshal(responseHeadersBuffer, &task.Headers)
		if err != nil {
			log.Println(err)
		}

		newTasks = append(newTasks, task)
	}
	log.Println("Storage.GetNewTasks got new tasks:", newTasks)
	// and updates statuses
	log.Println("Storage.GetNewTasks - bye")
	return newTasks
}
func (s Storage) BulkUpdateTaskStatuses(map[string]int)               {}
func (s Storage) BulkCreateTaskResults(results []entity.ResultEntity) {}
