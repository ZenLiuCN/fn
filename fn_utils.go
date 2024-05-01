package fn

import (
	"bytes"
	"errors"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
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
func CallerN(n uint) string {
	pc, file, lineno, ok := runtime.Caller(int(n) + 1)
	if ok {
		return fmt.Sprintf("%s\t%s:%d", runtime.FuncForPC(pc).Name(), file, lineno)
	}
	return "unknown:unknown"
}

// CallerFileN return call stack of file:line, skip n+1
func CallerFileN(n uint) string {
	_, file, lineno, ok := runtime.Caller(int(n) + 1)
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

// CallerShortFileN return call stack of file:line, skip n+1
func CallerShortFileN(n uint) string {
	_, file, lineno, ok := runtime.Caller(int(n) + 1)
	if ok {
		return fmt.Sprintf("%s:%d", filepath.Base(file), lineno)
	}
	return "unknown:unknown"
}

// CallerShortFile return call stack of file:line  ,skip 1
func CallerShortFile() string {
	_, file, lineno, ok := runtime.Caller(1)
	if ok {
		return fmt.Sprintf("%s:%d", filepath.Base(file), lineno)
	}
	return "unknown:unknown"
}

// CallerStackN dump  stack info of "function  file:line"
func CallerStackN(n, s uint) string {
	pc := make([]uintptr, s)
	ok := runtime.Callers(1, pc)
	if ok > 0 {
		frames := runtime.CallersFrames(pc)
		var frame runtime.Frame
		b := new(strings.Builder)
		more := n > 0
		for more {
			frame, more = frames.Next()
			b.WriteString(fmt.Sprintf("%s\t%s:%d\n", frame.Function, frame.File, frame.Line))
		}
		return b.String()
	}
	return "unknown:unknown"
}

// CallerShortStackN dump short stack "info of file:line"
func CallerShortStackN(n, s uint) string {
	pc := make([]uintptr, s)
	ok := runtime.Callers(1, pc)
	if ok > 0 {
		frames := runtime.CallersFrames(pc)
		var frame runtime.Frame
		b := new(strings.Builder)
		more := n > 0
		for more {
			frame, more = frames.Next()
			b.WriteString(fmt.Sprintf("%s:%d\n", frame.File, frame.Line))
		}
		return b.String()
	}
	return "unknown:unknown"
}

type ErrLine struct {
	Err  error
	Line string
}

func (e *ErrLine) Error() string {
	return fmt.Sprintf("\n%s\t%s", e.Line, e.Err)
}

// NewErrLine create an Error with original caller line
func NewErrLine(err error, skip uint) error {
	return &ErrLine{Err: err, Line: CallerFileN(skip)}
}

type ErrCombine struct {
	Err []error
}

func (e *ErrCombine) Error() string {
	var buffer bytes.Buffer
	buf := &buffer
	buf.Reset()
	for _, e := range e.Err {
		buf.WriteRune('\n')
		buf.WriteString(e.Error())
	}
	return buf.String()
}

// NewErrCombine combine errors
func NewErrCombine(err ...error) error {
	return &ErrCombine{Err: err}
}
