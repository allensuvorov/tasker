package storage

import (
	"github.com/allensuvorov/tasker/internal/server/domain"
	"net/http"
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

// TODO: create Test with the actual database
func TestTaskStorage_CreateTask(t *testing.T) {

	ts := NewTaskStorage()

	type args struct {
		te domain.TaskEntity
	}
	tests := []struct {
		name    string
		fields  *TaskStorage
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "create",
			fields: ts,
			args: args{
				te: domain.TaskEntity{
					ID:      "abc",
					Method:  http.MethodGet,
					URL:     "http://google.com",
					Headers: map[string]string{"Authentication": "Basic bG9naW46cGFzc3dvcmQ="},
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
