//go:build !gen

package fn

import (
	"errors"
	"strings"
	"testing"
)

func TestSimpleFileLine(t *testing.T) {
	e := Recover(func() {
		panic(1)
	}).Error()
	v := strings.Split(e, "\t")
	e = v[len(v)-2]
	if e[len(e)-2:] != "13" {
		t.Log(e)
		t.FailNow()
	}
	e = Recover(func() {
		Panic(errors.New("just fail"))
	}).Error()
	v = strings.Split(e, "\t")
	e = v[len(v)-2]
	if e[len(e)-2:] != "22" {
		t.Log(e)
		t.FailNow()
	}
	e = Recover(Panics(func() error {
		return errors.New("just")
	})).Error()
	v = strings.Split(e, "\t")
	e = v[len(v)-2]
	if e[len(e)-2:] != "30" {
		t.Log(e)
		t.FailNow()
	}
}
