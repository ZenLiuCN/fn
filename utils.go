package fn

import "errors"

var (
	//Nothing placeholder value
	Nothing = struct{}{}
	//ErrDuplicate for found duplicated value
	ErrDuplicate = errors.New("duplicated value")
)
