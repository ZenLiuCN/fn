package fn

import "errors"

var (
	//EmptyValue placeholder value
	EmptyValue = struct{}{}
	//ErrDuplicate for found duplicated value
	ErrDuplicate = errors.New("duplicated value")
)
