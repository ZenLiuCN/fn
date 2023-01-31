package fn

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

const (
	MAX = 10 //never bigger than 10
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
		P(pf("// Panic%d error or return other %d value\n", i, i))
		P(pf(`func Panic%d[`, i), n(i, A), p(` any]`), p(` (`), n(i, aA), p(", err error)"), q(i > 1, n(i, A)), p(`{
	if err!=nil{
		panic(err)
	}
	return `), n(i, a), p(`
}

`))
	}
}
func genRecovers() {
	P(pf("import \"fmt\" \n"))
	for i := 0; i < MAX; i++ {
	no:
		for o := 0; o < MAX; o++ {
			if i == o && o == 0 {
				continue no
			}
			P(pf("// Recover%d%d panic with func %d in %d out\n", i, o, i, o))
			P(pf(`func Recover%d%d[`, i, o), n(i+o, A), p(` any](fn func(`), n(i, A), p(")"), q(o > 1, r(i, i+o, A)),
				p(` )func(`), n(i, aA), p(")"), w(o > 0, "("), r(i, i+o, A), w(o > 0, ","), p("error"), w(o > 0, ")"), p(`{
	return func(`), n(i, aA), p(")("), r(i, i+o, aA), w(o > 0, ","), p(`err error){
`), p(`		defer func() {
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
		`), r(i, i+o, a), w(o > 0, " = "), p("fn("), n(i, a), p(`)
		return
	}
}

`))
		}

	}
}
func genDrops() {
	for i := 0; i < MAX; i++ {
	no:
		for o := 0; o < MAX; o++ {
			for d := 1; d <= o; d++ {
				if o == 0 {
					continue no
				}

				do := o - d
				ts := i + o

				P(pf("// Drop%d%d%d with func( %d in)(%d out) drop first %d result \n", i, o, d, i, o, d))
				P(pf(`func Drop%d%d%d[`, i, o, d), n(ts, A), p(` any] (fn func(`), n(i, A), p(`)`), q(o > 1, r(i, i+o, A)), p(`) func(`), n(i, A), p(")"), q(do > 1, r(i+d, i+o, A)), p(`{
	return func(`), n(i, aA), p(")"), q(do > 0, r(i+d, i+o, aA)), p(`{
		`), ww(d > 0, f(d, "_")), ww(do > 0, w(d > 0, ","), r(i+d, i+o, a)), w(o > 0, " = "), p("fn("), n(i, a), p(`)
		return
	}
}

`))

				if o == d {
					continue
				}
				P(pf("// DropLast%d%d%d with func( %d in)(%d out) drop last %d result \n", i, o, d, i, o, d))
				P(pf(`func DropLast%d%d%d[`, i, o, d), n(ts, A), p(` any](fn func(`), n(i, A), p(`)`), q(o > 1, r(i, i+o, A)), p(`) func(`), n(i, A), p(")"), q(do > 1, r(i, i+do, A)), p(` {
	return func(`), n(i, aA), p(")"), q(do > 0, r(i, i+do, aA)), p(`{
		`), ww(do != 0, r(i, i+do, a)), ww(d != 0, w(do != 0, ","), f(d, "_")), w(do != 0, " = "), p("fn("), n(i, a), p(`)
		return
	}
}

`))
			}
		}

	}
}
func genClosures() {
ni:
	for i := 0; i < MAX; i++ {
		for o := 0; o < MAX; o++ {
			for d := 1; d <= i; d++ {
				if i == 0 {
					continue ni
				}
				P(pf("// Closure%d%d%d with func( %d in)(%d out) closure first %d argument \n", i, o, d, i, o, d))
				ts := i + o
				P(pf(`func Closure%d%d%d[`, i, o, d), n(ts, A), p(` any](`), r(0, d, aA), w(d > 0, ","), p(`fn func(`), n(i, A), p(`)`), q(o > 1, r(i, i+o, A)), p(` )func(`), r(d, i, A), p(")"), q(o > 1, r(i, ts, A)), p(`{
	return func(`), r(d, i, aA), p(")"), q(o > 0, r(i, ts, aA)), p(`{
		`), r(i, ts, a), w(o > 0, " = "), p("fn("), n(i, a), p(`)
		return
	}
}

`))
				if i == d {
					continue
				}
				P(pf("// ClosureLast%d%d%d with func( %d in)(%d out) fix last %d argument \n", i, o, d, i, o, d))
				P(pf(`func ClosureLast%d%d%d[`, i, o, d), n(ts, A), p(` any](`), r(i-d, i, aA), w(d > 0, ","), p(`fn func(`), n(i, A), p(`)`), q(o > 1, r(i, i+o, A)), p(` )func(`), r(0, i-d, A), p(")"), q(o > 1, r(i, i+o, A)), p(`{
	return func(`), r(0, i-d, aA), p(")"), q(o > 0, r(i, i+o, aA)), p(`{
		`), r(i, i+o, a), w(o != 0, " = "), p("fn("), n(i, a), p(`)
		return
	}
}

`))
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
			P(pf("// PanicsFn%d%d panic with func %d in %d out (last out is an error), returns a Runnable.\n", i, o, i, o+1))
			P(pf(`func PanicsFn%d%d[`, i, o), n(i+o, A), p(` any](`), n(i, aA), w(i > 0, ","), p(` fn func(`), n(i, A), p(`)`), w(o > 0, "("), r(i, i+o, A), w(o > 0, ","), p(`error`), w(o > 0, ")"), p(` )func()`), p(`{
	return func(){
		var err error
		`), ww(o > 0, fr(i, i+o, "_"), p(",")), p("err= "), p("fn("), n(i, a), p(`)
		if err!=nil{
			panic(err)
		}
		return
	}
}

`))
		}

	}
}
func TestGenerate(t *testing.T) {
	generate("panics.go", genPanics)
	//generate("recovers.go", genRecovers)
	//generate("drops.go", genDrops)
	//generate("closures.go", genClosures)
	//generate("panics_fn.go", genPanicsFn)
	/*	buf = bufio.NewWriter(os.Stdout)
		P(n(10, A))
		Panic(buf.Flush())*/
}

func pf(p string, a ...any) func() {
	return func() {
		_, _ = buf.WriteString(fmt.Sprintf(p, a...))
	}
}
func p(v ...string) func() {
	return func() {
		for _, u := range v {
			_, _ = buf.WriteString(u)
		}
	}
}

func a(i int) {
	p(string(rune(i + 'a')))()
}
func aA(i int) {
	p(string(rune(i + 'a')))()
	p(" ")()
	p(string(rune(i + 'A')))()
}
func A(i int) {
	p(string(rune(i + 'A')))()
}

func n(i int, v func(j int)) func() {
	return func() {
		for j := 0; j < i; j++ {
			if j != 0 {
				p(",")()
			}
			v(j)
		}
	}
}
func f(i int, v string) func() {
	return func() {
		for j := 0; j < i; j++ {
			if j != 0 {
				p(",")()
			}
			p(v)()
		}
	}
}
func fr(i int, u int, v string) func() {
	return func() {
		for j := i; j < u; j++ {
			if j != i {
				p(",")()
			}
			p(v)()
		}
	}
}
func r(i int, u int, v func(j int)) func() {
	if i < 0 || i >= u {
		return func() {

		}
	}
	return func() {
		for j := i; j < u; j++ {
			if j != i {
				p(",")()
			}
			v(j)
		}
	}
}
func q(b bool, act func()) func() {
	return func() {
		if b {
			p("(")()
		}
		act()
		if b {
			p(")")()
		}
	}
}
func w(b bool, v string) func() {
	return func() {
		if b {
			p(v)()
		}
	}
}
func ww(b bool, act ...func()) func() {
	return func() {
		if b {
			for _, f2 := range act {
				f2()
			}
		}
	}
}
func P(v ...func()) {
	for _, f2 := range v {
		f2()
	}
}
