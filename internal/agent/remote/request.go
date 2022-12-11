package remote

import "github.com/allensuvorov/tasker/internal/server/domain"

type Request struct{}

func NewRequest() Request {
	return Request{}
}

func (r Request) Request(domain.TaskEntity) domain.ResultEntity {
	return domain.ResultEntity{}
}
