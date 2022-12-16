package handlers

import (
	"net/http"
	"testing"
)

func TestTaskHandler_CreateTask(t *testing.T) {
	type fields struct {
		taskService TaskService
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			th := TaskHandler{
				taskService: tt.fields.taskService,
			}
			th.CreateTask(tt.args.w, tt.args.r)
		})
	}
}
