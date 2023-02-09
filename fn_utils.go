package fn

import (
	"errors"
	"fmt"
	"runtime"
)

var (
	//Nothing placeholder value
	Nothing = struct{}{}
	//ErrDuplicate for found duplicated value
	ErrDuplicate = errors.New("duplicated value")
)

// Caller return call stack ,skip 1
func Caller() string {
	pc, file, lineno, ok := runtime.Caller(1)
	if ok {
		return fmt.Sprintf("%s\t%s:%d", runtime.FuncForPC(pc).Name(), file, lineno)
	}
	return "unknown:unknown"
}

// CallerN return call stack skip n+1
func CallerN(n int) string {
	pc, file, lineno, ok := runtime.Caller(n + 1)
	if ok {
		return fmt.Sprintf("%s\t%s:%d", runtime.FuncForPC(pc).Name(), file, lineno)
	}
	return "unknown:unknown"
}

// CallerFileN return call stack of file:line, skip n+1
func CallerFileN(n int) string {
	_, file, lineno, ok := runtime.Caller(n + 1)
	if ok {
		return fmt.Sprintf("%s:%d", file, lineno)
	}
	return "unknown:unknown"
}

// CallerFile return call stack of file:line  ,skip 1
func CallerFile() string {
	_, file, lineno, ok := runtime.Caller(1)
	if ok {
		return fmt.Sprintf("%s:%d", file, lineno)
	}
	return "unknown:unknown"
}
