package entity

type Headers map[string]string

type TaskEntity struct {
	ID      string
	Method  string
	URL     string
	Headers Headers
}

type ResultEntity struct {
	TaskID                 string
	TaskStatus             string
	ResponseHttpStatusCode int
	ResponseHeaders        Headers
	ResponseBodyLength     int
}
