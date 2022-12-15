package remote

import (
	"github.com/allensuvorov/tasker/internal/server/domain/entity"
)

type Request struct{}

func NewRequest() Request {
	return Request{}
}

func (r Request) Request(entity.TaskEntity) entity.ResultEntity {
	return entity.ResultEntity{}
}
