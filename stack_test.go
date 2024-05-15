package fn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	corpus = []byte("ABC ABCDAB ABCDABCDABDE")
	find   = []byte("ABCDABD")
)

func TestMatchKMP(t *testing.T) {

	assert.Equal(t, 15, MatchKMP(corpus, find...))
}
func TestMatchZ(t *testing.T) {
	assert.Equal(t, 15, MatchZ(corpus, '?', find...))
}
func BenchmarkMatchZ(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		MatchZ(corpus, '?', find...)
	}
}
func BenchmarkMatchKMP(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		MatchKMP(corpus, find...)
	}
}
