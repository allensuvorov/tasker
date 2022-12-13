package domain

type TaskEntity struct {
	ID           string
	Method       string
	URL          string
	Headers      map[string]string
	Status       string
	ResultEntity int
}

type ResultEntity struct {
	ID                     string
	ResponseHttpStatusCode string
	ResponseHeaders        map[string]string
	ResponseBodyLength     int
}
