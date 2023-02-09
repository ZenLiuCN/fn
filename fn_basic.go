package fn

// manual write functions
import "fmt"

type PackedError struct {
	Origin any
	Caller string
}

func (p PackedError) Error() string {
	switch err := p.Origin.(type) {
	case error:
		return fmt.Sprintf("%s\t%s", p.Caller, err)
	default:
		return fmt.Sprintf("%s\t%#v", p.Caller, err)
	}
}

//NewCallFileError impl ErrorPacker with caller file line
func NewCallFileError(origin any, caller int) error {
	return &PackedError{Origin: origin, Caller: CallerFileN(caller)}
}

//NewCallFuncError impl ErrorPacker with caller func and file line
func NewCallFuncError(origin any, caller int) error {
	return &PackedError{Origin: origin, Caller: CallerN(caller)}
}

//ErrorPacker func to pack error with origin caller location
type ErrorPacker func(any, int) error

var (
	Packer ErrorPacker = NewCallFileError
)

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

//Panic error
func Panic(err error) {
	if err != nil {
		panic(Packer(err, 2))
	}
}

//Panics return a Runnable
func Panics(fn func() error) func() {
	return func() {
		if err := fn(); err != nil {
			panic(Packer(err, 3))
		}
	}
}
