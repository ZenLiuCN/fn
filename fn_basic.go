package fn

// manual write functions
import "fmt"

// Recover a runnable
func Recover(fn func()) (err error) {
	defer func() {
		z := recover()
		switch z.(type) {
		case error:
			err = z.(error)
			return
		case nil:
			return
		default:
			err = fmt.Errorf("%#v", z)
		}
	}()
	fn()
	return
}

//Panic error
func Panic(err error) {
	if err != nil {
		panic(err)
	}
}

//Panics return a Runnable
func Panics(fn func() error) func() {
	return func() {
		if err := fn(); err != nil {
			panic(err)
		}
	}
}
