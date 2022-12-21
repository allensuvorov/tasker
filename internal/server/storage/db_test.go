package storage

import (
	"fmt"
	"github.com/allensuvorov/tasker/internal/server/domain/entity"
	"log"
	"net/http"
	"reflect"
	"testing"
)

//func TestNewTaskStorage(t *testing.T) {
//	tests := []struct {
//		name string
//		want *TaskStorage
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := NewTaskStorage(); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("NewTaskStorage() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

// TODO: create Test with mock database
// TODO add cleanup - garbage in, garbage out

func TestTaskStorage_CreateTask_RealDB(t *testing.T) {

	ts := NewTaskStorage()

	type args struct {
		te entity.TaskEntity
	}
	tests := []struct {
		name    string
		fields  *TaskStorage
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "created",
			fields: ts,
			args: args{
				te: entity.TaskEntity{
					ID:      "abcde",
					Method:  http.MethodGet,
					URL:     "http://google.com",
					Headers: entity.Headers{"Authentication": "Basic bG9naW46cGFzc3dvcmQ=", "Content-type": "JSON"},
				},
			},
			wantErr: false,
		},
		{
			name:   "created, this task is to check its own status",
			fields: ts,
			args: args{
				te: entity.TaskEntity{
					ID:      "checkstatus",
					Method:  http.MethodGet,
					URL:     "http://localhost:8080/task/checkstatus",
					Headers: entity.Headers{"Content-type": "JSON"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := TaskStorage{
				DB: tt.fields.DB,
			}
			if err := ts.CreateTask(tt.args.te); (err != nil) != tt.wantErr {
				t.Errorf("CreateTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTaskStorage_Create30Tasks_RealDB(t *testing.T) {

	ts := NewTaskStorage()

	type args struct {
		te entity.TaskEntity
	}
	tests := []struct {
		name    string
		fields  *TaskStorage
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "created",
			fields: ts,
			args: args{
				te: entity.TaskEntity{
					ID:      "boredapi",
					Method:  http.MethodGet,
					URL:     "https://www.boredapi.com/api/activity",
					Headers: entity.Headers{"Content-type": "JSON"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := TaskStorage{
				DB: tt.fields.DB,
			}

			for i := 0; i < 30; i++ {
				tt.args.te.ID = "testtask" + fmt.Sprint(i)
				if err := ts.CreateTask(tt.args.te); (err != nil) != tt.wantErr {
					t.Errorf("CreateTask() error = %v, wantErr %v", err, tt.wantErr)
				}
			}

			// run agent

			// clean Data

		})
	}
}

//func TestTaskStorage_CreateTask(t *testing.T) {
//	type fields struct {
//		DB *sql.DB
//	}
//	type args struct {
//		te domain.TaskEntity
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			ts := TaskStorage{
//				DB: tt.fields.DB,
//			}
//			if err := ts.CreateTask(tt.args.te); (err != nil) != tt.wantErr {
//				t.Errorf("CreateTask() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}

// TODO add an clean up - garbage in, garbage out
func TestTaskStorage_GetTaskStatus_RealDB(t *testing.T) {

	ts := NewTaskStorage()

	type args struct {
		taskID string
		result entity.ResultEntity
	}
	tests := []struct {
		name    string
		fields  *TaskStorage
		args    args
		want    entity.ResultEntity
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "found",
			args: args{
				taskID: "abc",
				result: entity.ResultEntity{
					TaskID:                 "abc",
					TaskStatus:             "done",
					ResponseHttpStatusCode: 200,
					//ResponseHeaders:        http.Header{"Result-Header1": {"sample-data"}, "Content-type": {"JSON"}},
					ResponseHeaders:    entity.Headers{"Result-Header1": "sample-data", "Content-type": "JSON"},
					ResponseBodyLength: "50",
				},
			},
			fields: ts,
			want: entity.ResultEntity{
				TaskID:                 "abc",
				TaskStatus:             "done",
				ResponseHttpStatusCode: 200,
				//ResponseHeaders:        http.Header{"Result-Header1": {"sample-data"}, "Content-type": {"JSON"}},
				ResponseHeaders:    entity.Headers{"Result-Header1": "sample-data", "Content-type": "JSON"},
				ResponseBodyLength: "50",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := TaskStorage{
				DB: tt.fields.DB,
			}

			_, err := ts.DB.Exec(
				`INSERT INTO tasks (task_id, task_status, result_http_status_code, result_headers, result_body_length)
				VALUES ($1, $2, $3, $4, $5);`,
				tt.args.result.TaskID, tt.args.result.TaskStatus, tt.args.result.ResponseHttpStatusCode, tt.args.result.ResponseHeaders, tt.args.result.ResponseBodyLength,
			)
			if err != nil {
				log.Fatal(err)
			}

			got, err := ts.GetTaskStatus(tt.args.taskID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTaskStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTaskStatus() got = %v, want %v", got, tt.want)
			}
		})
	}
}

//func TestTaskStorage_GetTaskStatus(t *testing.T) {
//	type fields struct {
//		DB *sql.DB
//	}
//	type args struct {
//		taskID string
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    domain.ResultEntity
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			ts := TaskStorage{
//				DB: tt.fields.DB,
//			}
//			got, err := ts.GetTaskStatus(tt.args.taskID)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("GetTaskStatus() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("GetTaskStatus() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
