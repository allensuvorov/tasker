package domain

type TaskEntity struct {
	ID      string
	Method  string
	URL     string
	Headers map[string]string
}
