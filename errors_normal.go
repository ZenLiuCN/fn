//go:build NormalErr

package fn

import (
	"errors"
	"fmt"
)

//Error create new error with line info
func Error(arg ...any) error {
	return errors.New(fmt.Sprint(arg...))
}

//Errorf create new error with line info
func Errorf(pattern string, arg ...any) error {
	return errors.New(fmt.Sprintf(pattern, arg...))
}

var (
	Packer ErrorPacker = func(a any, _ uint) error {
		switch e := a.(type) {
		case error:
			return e
		default:
			return Error(e)
		}
	}
)
