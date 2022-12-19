package storage

import (
	"database/sql"
	"encoding/json"
	"errors"
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

func (s *Storage) GetNewTasks() ([]entity.TaskEntity, error) {
	log.Println("Storage.getNewTasks - hello")

	if s.DB == nil {
		return nil, errors.New("you haven`t opened the database connection")
	}
	// TODO: do this in transaction, so that if there is an error amidst execution rollback saves the day
	rows, err := s.DB.Query(`UPDATE tasks 
		SET task_status = $1 
		WHERE task_status = $2 RETURNING task_id, task_request_method, task_headers, task_url;`,
		"in_process", "new")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if rows.Err() != nil {
		return nil, err
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
	log.Println("Storage.getNewTasks got new tasks:", newTasks)
	// and updates statuses
	log.Println("Storage.getNewTasks - bye")
	return newTasks, nil
}
func (s *Storage) BulkAddTaskResults(results []entity.ResultEntity) error {
	log.Println("Storage.BulkAddTaskResults - hello")

	if s.DB == nil {
		return errors.New("you haven`t opened the database connection")
	}

	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(
		`UPDATE tasks 
		SET task_status = $2, result_http_status_code = $3, result_headers = $4, result_body_length = $5  
    	WHERE task_id = $1;`,
	)
	if err != nil {
		return err
	}

	defer stmt.Close()

	for _, r := range results {
		if _, err = stmt.Exec(r.TaskID, r.TaskStatus, r.ResponseHttpStatusCode, r.ResponseHeaders, r.ResponseBodyLength); err != nil {
			return err
		}
	}

	log.Println("Storage.BulkAddTaskResults - bye")

	return tx.Commit()
}
