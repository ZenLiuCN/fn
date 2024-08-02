package fn

// manual write functions
import (
	"errors"
	"fmt"
)

// NewCallShortFileError use [errors.Join] with caller func and file line (v0.1.34)
func NewCallShortFileError(origin any, caller uint) error {
	switch x := origin.(type) {
	case nil:
		return errors.New(CallerShortFileN(caller))
	case error:
		return errors.Join(x, errors.New(CallerShortFileN(caller)))
	default:
		return errors.Join(fmt.Errorf("%#+v", x), errors.New(CallerShortFileN(caller)))
	}

}

// NewCallFileError use [errors.Join] with caller func and file line (v0.1.34)
func NewCallFileError(origin any, caller uint) error {
	switch x := origin.(type) {
	case nil:
		return errors.New(CallerFileN(caller))
	case error:
		return errors.Join(x, errors.New(CallerFileN(caller)))
	default:
		return errors.Join(fmt.Errorf("%#+v", x), errors.New(CallerFileN(caller)))
	}
}

// NewCallFuncError use [errors.Join] with caller func and file line (v0.1.34)
func NewCallFuncError(origin any, caller uint) error {
	switch x := origin.(type) {
	case nil:
		return errors.New(CallerN(caller))
	case error:
		return errors.Join(x, errors.New(CallerN(caller)))
	default:
		return errors.Join(fmt.Errorf("%#+v", x), errors.New(CallerN(caller)))
	}
}

// ErrorPacker func to pack error with origin caller location
type ErrorPacker func(any, uint) error

// Recover a runnable
func Recover(fn func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = Packer(r, 3)
		}
	}()
	fn()
	return
}

// Panic error
func Panic(err error) {
	if err != nil {
		panic(Packer(err, 2))
	}
}

// Panics return a Runnable
func Panics(fn func() error) func() {
	return func() {
		if err := fn(); err != nil {
			panic(Packer(err, 3))
		}
	}
}
