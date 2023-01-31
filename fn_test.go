package fn

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

const (
	MAX = 10
)

var (
	buf *bufio.Writer
)

func panic1[A any](a A, err error) A {
	if err != nil {
		panic(err)
	}
	return a
}
func generate(path string, act func()) {
	file := panic1(os.OpenFile(path, os.O_TRUNC|os.O_CREATE|os.O_RDONLY, os.ModePerm))
	defer PanicFn(file.Close)
	buf = bufio.NewWriter(file)
	_ = panic1(buf.WriteString("package fn \n//generated file should not edit\n\n"))
	act()
	Panic(buf.Flush())
}
func genPanics() {
	for i := 1; i < MAX; i++ {
		pntf("// Panic%d error or return other %d value\n", i, i)
		pntf(`func Panic%d[`, i)
		n(i, A)
		pnt(` any](`)
		n(i, aA)
		pnt(", err error)(")
		n(i, A)
		pnt("){\n\tif err!=nil{\n\t\tpanic(err)\n\t}\n\treturn ")
		n(i, a)
		pnt("\n}\n\n")
	}
}
func genRecovers() {
	for i := 0; i < MAX; i++ {
	no:
		for o := 0; o < MAX; o++ {
			if i == o && o == 0 {
				continue no
			}
			pntf("// Recover%d%d panic with func %d in %d out\n", i, o, i, o)
			pntf(`func Recover%d%d[`, i, o)
			n(i+o, A)
			pnt(` any](fn func(`)
			n(i, A)
			pnt(`)(`)
			r(i, i+o, A)
			pnt(`) )func(`)
			n(i, aA)
			pnt(")(")
			r(i, i+o, A)
			if o != 0 {
				pnt(",")
			}
			pnt(`error){
	return func(`)
			n(i, aA)
			pnt(")(")
			r(i, i+o, aA)
			if o != 0 {
				pnt(",")
			}
			pnt("err error){\n")
			pnt(`		defer func() {
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
		`)
			r(i, i+o, a)
			if o != 0 {
				pnt("= ")
			}
			pnt("fn(")
			n(i, a)
			pnt(")\n\t\treturn\n\t}\n}\n\n")
		}

	}
}
func genDrops() {
	for i := 0; i < MAX; i++ {
	no:
		for o := 0; o < MAX; o++ {
			for d := 1; d < o; d++ {
				if o == 0 {
					continue no
				}
				pntf("// Drop%d%d%d with func( %d in)(%d out) drop first %d result \n", i, o, d, i, o, d)
				pntf(`func Drop%d%d%d[`, i, o, d)
				n(i+o, A)
				pnt(` any](fn func(`)
				//in param
				n(i, A)
				pnt(`)(`)
				//out param
				r(i, i+o, A)
				pnt(`) )func(`)
				//in param
				n(i, aA)
				pnt(")(")
				//out after drop
				r(i+d, i+o, A)
				pnt(`){
	return func(`)
				n(i, aA)
				pnt(")(")
				r(i+d, i+o, aA)
				pnt("){\n\t\t")
				f(d, "_")
				pnt(",")
				r(i+d, i+o, a)
				pnt(" = ")
				pnt("fn(")
				n(i, a)
				pnt(")\n\t\treturn\n\t}\n}\n\n")

				pntf("// DropLast%d%d%d with func( %d in)(%d out) drop last %d result \n", i, o, d, i, o, d)
				pntf(`func DropLast%d%d%d[`, i, o, d)
				n(i+o, A)
				pnt(` any](fn func(`)
				//in param
				n(i, A)
				pnt(`)(`)
				//out param
				r(i, i+o, A)
				pnt(`) )func(`)
				//in param
				n(i, aA)
				pnt(")(")
				//out after drop
				r(i, i+o-d, A)
				pnt(`){
	return func(`)
				n(i, aA)
				pnt(")(")
				r(i, i+o-d, aA)
				pnt("){\n\t\t")
				r(i, i+o-d, a)
				pnt(",")
				f(d, "_")
				pnt(" = ")
				pnt("fn(")
				n(i, a)
				pnt(")\n\t\treturn\n\t}\n}\n\n")
			}
		}

	}
}
func genClosures() {
ni:
	for i := 0; i < MAX; i++ {
		for o := 0; o < MAX; o++ {
			for d := 1; d < i; d++ {
				if i == 0 {
					continue ni
				}
				pntf("// Closure%d%d%d with func( %d in)(%d out) fix first %d argument \n", i, o, d, i, o, d)
				pntf(`func Closure%d%d%d[`, i, o, d)
				n(i+o, A)
				pnt(` any](`)
				r(0, i-d, aA)
				pnt(`,fn func(`)
				//in param
				n(i, A)
				pnt(`)(`)
				//out param
				r(i, i+o, A)
				pnt(`)`)
				pnt(` )func(`)
				//in param
				r(i-d, i, A)
				pnt(")(")
				//outs
				r(i, i+o, A)
				pnt(`){
	return func(`)
				r(i-d, i, aA)
				pnt(")(")
				r(i, i+o, aA)
				pnt("){\n\t\t")
				r(i, i+o, a)
				if o != 0 {
					pnt(" = ")
				}
				pnt("fn(")
				n(i, a)
				pnt(")\n\t\treturn\n\t}\n}\n\n")

				pntf("// ClosureLast%d%d%d with func( %d in)(%d out) fix last %d argument \n", i, o, d, i, o, d)
				pntf(`func ClosureLast%d%d%d[`, i, o, d)
				n(i+o, A)
				pnt(` any](`)
				r(i-d, i, aA)
				pnt(`,fn func(`)
				//in param
				n(i, A)
				pnt(`)(`)
				//out param
				r(i, i+o, A)
				pnt(`)`)
				pnt(` )func(`)
				//in param
				r(0, i-d, A)
				pnt(")(")
				//outs
				r(i, i+o, A)
				pnt(`){
	return func(`)
				r(0, i-d, aA)
				pnt(")(")
				r(i, i+o, aA)
				pnt("){\n\t\t")
				r(i, i+o, a)
				if o != 0 {
					pnt(" = ")
				}
				pnt("fn(")
				n(i, a)
				pnt(")\n\t\treturn\n\t}\n}\n\n")
			}
		}

	}
}
func genPanicsFn() {
	for i := 0; i < MAX; i++ {
	no:
		for o := 0; o < MAX; o++ {
			if i == o && o == 0 {
				continue no
			}
			pntf("// PanicsFn%d%d panic with func %d in %d out (last out is an error), returns a Runnable.\n", i, o, i, o+1)
			pntf(`func PanicsFn%d%d[`, i, o)
			n(i+o, A)
			pnt(` any](`)
			n(i, aA)
			if i != 0 {
				pnt(",")
			}
			pnt(` fn func(`)
			n(i, A)
			pnt(`)(`)
			r(i, i+o, A)
			if o != 0 {
				pnt(",")
			}
			pnt(`error) )func()`)
			pnt(`{
	return func(){
		var err error
		`)
			if o != 0 {
				fr(i, i+o, "_")
				pnt(",")
			}
			pnt("err= ")
			pnt("fn(")
			n(i, a)
			pnt(")\n\t\tif err!=nil{\n\t\t\tpanic(err)\n\t\t}\n\t\treturn\n\t}\n}\n\n")
		}

	}
}
func TestGenerate(t *testing.T) {
	//generate("panics.go", genPanics)
	//generate("recovers.go", genRecovers)
	//generate("drops.go", genDrops)
	//generate("closures.go", genClosures)
	//generate("panics_fn.go", genPanicsFn)
}

func pntf(p string, a ...any) {
	_, _ = buf.WriteString(fmt.Sprintf(p, a...))
}
func pnt(v ...string) {
	for _, x := range v {
		_, _ = buf.WriteString(x)
	}
}
func id(i int, v string) {
	pnt(strings.Repeat("\t", i), v)
}
func a(i int) {
	pnt(string(rune(i + 'a')))
}
func aA(i int) {
	pnt(string(rune(i + 'a')))
	pnt(" ")
	pnt(string(rune(i + 'A')))
}
func A(i int) {
	pnt(string(rune(i + 'A')))
}
func n(i int, v func(j int)) {
	for j := 0; j < i; j++ {
		if j != 0 {
			pnt(",")
		}
		v(j)
	}
}
func f(i int, v string) {
	for j := 0; j < i; j++ {
		if j != 0 {
			pnt(",")
		}
		pnt(v)
	}
}
func fr(i int, u int, v string) {
	for j := i; j < u; j++ {
		if j != i {
			pnt(",")
		}
		pnt(v)
	}
}
func r(i int, u int, v func(j int)) {
	for j := i; j < u; j++ {
		if j != i {
			pnt(",")
		}
		v(j)
	}
}
