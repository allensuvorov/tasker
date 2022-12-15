package entity

import "net/http"

type TaskEntity struct {
	ID      string
	Method  string
	URL     string
	Headers http.Header
}

type ResultEntity struct {
	TaskID                 string
	TaskStatus             string
	ResponseHttpStatusCode int
	ResponseHeaders        http.Header
	ResponseBodyLength     int
}
