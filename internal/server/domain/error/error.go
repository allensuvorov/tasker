package err

import "errors"

var ErrorNotFound = errors.New("not found")
var ErrorRecordExists = errors.New("record exists")
var ErrorWrongContentType = errors.New("wrong content type in request")
