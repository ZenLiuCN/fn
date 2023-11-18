package fn

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	fmt.Println(Recover(func() {
		func() {
			panic(Error(1))
		}()
	}))
}
