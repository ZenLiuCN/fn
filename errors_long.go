//go:build FullErr

package fn

import (
	"errors"
	"fmt"
)

//Error create new error with line info
func Error(arg ...any) error {
	return errors.New(fmt.Sprint(append([]any{CallerFileN(1)}, arg...)))
}

//Errorf create new error with line info
func Errorf(pattern string, arg ...any) error {
	return errors.New(fmt.Sprintf("%s "+pattern, append([]any{CallerFileN(1)}, arg...)))
}

var (
	Packer ErrorPacker = NewCallFileError
)
