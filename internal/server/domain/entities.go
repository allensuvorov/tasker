package domain

type TaskEntity struct {
	ID           string
	Method       string
	URL          string
	Headers      map[string]string
	ResultEntity int
}

type ResultEntity struct {
	ID                     string
	TaskStatus             string
	ResponseHttpStatusCode string
	ResponseHeaders        map[string]string
	ResponseBodyLength     int
}
