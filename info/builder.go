package info

import (
	"bytes"
	"fmt"
	"go/types"
	"io"
	"runtime"
)

var (
	Sizes = types.SizesFor("gc", runtime.GOARCH)
)

func BuildType(b io.Writer, pkg, name string, t *types.Struct, exported bool) (r int) {
	w := writer{b}
	n := t.NumFields()
	offset := 0
	size := 0

	fields := make([]*types.Var, t.NumFields())
	for i := 0; i < t.NumFields(); i++ {
		fields[i] = t.Field(i)
	}
	sizes := Sizes
	sizer := sizes.Offsetsof
	offsets := sizer(fields)
	w.f(`	Type{
	Name: "%s",
	Path: "%s",
	Fields:[]Field{
`, name, pkg)
	for i := 0; i < n; i++ {
		f := t.Field(i)
		if exported && !f.Exported() {
			continue
		}
		r++
		size = int(Sizes.Sizeof(f.Type()))
		offset = int(offsets[i])
		w.i(1).f(`	{
		Name:"%s",
		Tag:"%s",
		Id:"%s",
		Embedded:%t,
		Offset:%d,
		Size:%d,
		},
`, f.Name(), t.Tag(i), f.Type().String(), f.Embedded(), offset, size)
	}
	w.i(1).f("},	\n}")
	return
}

type writer [1]io.Writer

func (w writer) f(format string, args ...any) writer {
	_, _ = fmt.Fprintf(w[0], format, args...)
	return w
}
func (w writer) i(n int) writer {
	_, _ = w[0].Write(bytes.Repeat([]byte{'\t'}, n))
	return w
}
func (w writer) n() writer {
	_, _ = w[0].Write([]byte{'\n'})
	return w
}
