package domain

type TaskEntity struct {
	ID           string
	Method       string
	URL          string
	Headers      map[string]string
	ResultEntity int
}

type ResultEntity struct {
	ID             string
	Status         string
	HttpStatusCode string
	Headers        map[string]string
	Length         int
}
