package fn

import (
	"fmt"
	"gotest.tools/v3/assert"
	"testing"
)

func TestHttpErrorFormat(t *testing.T) {
	e := NewHttpError(200, fmt.Errorf("some err"))
	t.Log(string(defaultFormatter(e)))
}
func TestStringData(t *testing.T) {
	v := "123456787"
	b := stringData(v)
	assert.DeepEqual(t, b, []byte(v))
}
