package hocon

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type Some struct {
	Name  string `json:"name,omitempty"`
	Param bool   `json:"param,omitempty"`
	Num   int64  `json:"num"`
	Obj   *OBJ   `json:"obj,omitempty"`
	Array []Some `json:"array,omitempty"`
}
type OBJ struct {
	Name  string    `json:"name,omitempty"`
	Param bool      `json:"param,omitempty"`
	Time  time.Time `json:"time"`
	Obj   *Some     `json:"obj,omitempty"`
	Array []Some    `json:"array,omitempty"`
}

var sample = Some{
	Name:  "123",
	Param: false,
	Obj: &OBJ{
		Name:  "11",
		Param: false,
		Obj:   nil,
		Array: nil,
	},
	Array: []Some{{
		Name:  "1a1",
		Param: false,
		Obj:   nil,
		Array: nil,
	}, {
		Name:  "1a2",
		Param: false,
		Obj:   nil,
		Array: nil,
	}},
}

func TestMarshal(t *testing.T) {
	c, err := Marshal(sample)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf(c.String())
	var s Some
	err = Unmarshal(c, &s)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, s, sample)
}
func BenchmarkMarshal(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := Marshal(sample)
		if err != nil {
			b.Fatal(err)
		}
	}
}
func BenchmarkUnmarshal(b *testing.B) {
	c, err := Marshal(sample)
	if err != nil {
		b.Fatal(err)
	}
	var s Some
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		err = Unmarshal(c, &s)
		if err != nil {
			b.Fatal(err)
		}
	}
}
