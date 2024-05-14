package fn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatch(t *testing.T) {
	corpus := []byte("ABC ABCDAB ABCDABCDABDE")
	find := []byte("ABCDABD")
	assert.Equal(t, 15, Match(corpus, find...))
}
func TestAnyMatch(t *testing.T) {
	corpus := []byte("ABC ABCDAB ABCDABCDABDE")
	find := []byte("ABCDABD")
	assert.Equal(t, 15, AnyMatch(corpus, '?', '*', find...))
}
