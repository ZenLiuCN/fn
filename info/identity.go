package info

import (
	"bytes"
	"fmt"
	"github.com/ZenLiuCN/fn"
)

const (
	KindUnknown Kind = iota
	KindType
	KindSlice
	KindMap
	KindPointer
	KindFunc
	KindArray
	KindBuiltIn
	KindChan
)

type (
	Kind     int
	Identity interface {
		Kind() Kind
		fmt.Stringer
	}
)
type IdBuiltIn string

func (i IdBuiltIn) Kind() Kind {
	return KindBuiltIn
}
func (i IdBuiltIn) String() string {
	return string(i)
}

type IdType [2]string

func (i IdType) Kind() Kind {
	return KindType
}
func (i IdType) String() string {
	return fmt.Sprintf("%s.%s",
		i[0],
		i[1],
	)
}

type IdSlice [1]Identity

func (i IdSlice) Kind() Kind {
	return KindSlice
}
func (i IdSlice) String() string {
	return fmt.Sprintf("[]%s",
		i[0].String(),
	)
}

type IdPointer [1]Identity

func (i IdPointer) Kind() Kind {
	return KindPointer
}
func (i IdPointer) String() string {
	return fmt.Sprintf("*%s",
		i[0].String(),
	)
}

type IdMap [2]Identity

func (i IdMap) Kind() Kind {
	return KindMap
}
func (i IdMap) String() string {
	return fmt.Sprintf("map[%s]%s",
		i[0].String(),
		i[1].String(),
	)
}

type IdFunc [2][]Identity

func (i IdFunc) Kind() Kind {
	return KindFunc
}
func (i IdFunc) String() string {
	if len(i[1]) == 0 {
		return fmt.Sprintf("func(%s)",
			fn.SliceJoinString(i[0], ",", Identity.String))
	}
	return fmt.Sprintf("func(%s)(%s)",
		fn.SliceJoinString(i[0], ",", Identity.String),
		fn.SliceJoinString(i[1], ",", Identity.String))
}

type IdArray struct {
	Length  uint
	Element Identity
}

func (i IdArray) Kind() Kind {
	return KindArray
}
func (i IdArray) String() string {
	return fmt.Sprintf("[%d]%s", i.Length, i.Element)
}

type IdChan struct {
	Dir     int8 // 0: send and receive, positive: send, negative: receive
	Element Identity
}

func (i IdChan) Kind() Kind {
	return KindChan
}
func (i IdChan) String() string {
	if i.Dir > 0 {
		return "chan<- " + i.Element.String()
	} else if i.Dir < 0 {
		return "<-chan " + i.Element.String()
	}
	return "chan " + i.Element.String()
}

var (
	stringReader = Reader[string, []byte](0)
)

func ParseId(i Id) (r Identity) {
	b := []byte(i)
	r, _ = parseType(b)
	return r
}

func parseType(i []byte) (r Identity, t int) {
	if bytes.HasPrefix(i, []byte("chan")) || bytes.HasPrefix(i, []byte("<-chan")) {
		return parseChan(i)
	}
	if bytes.HasPrefix(i, []byte("map")) {
		return parseMap(i)
	}
	if bytes.HasPrefix(i, []byte("func")) {
		return parseFunc(i)
	}
	if bytes.HasPrefix(i, []byte("[]")) {
		return parseSlice(i)
	}
	if bytes.HasPrefix(i, []byte("[")) {
		return parseArray(i)
	}
	if bytes.HasPrefix(i, []byte("*")) {
		return parsePointer(i)
	}
	var m IdType
	n := len(i)
	b := new(bytes.Buffer)
loop:
	for t = 0; t < n; t++ {
		switch c := i[t]; c {
		case '.':
			m[0] = b.String()
			b.Reset()
		case ' ', ',', '<', '*', '(', ')', '[', ']':
			m[1] = b.String()
			b.Reset()
			t -= 1
			break loop
		default:
			b.WriteByte(c)
		}
	}
	if b.Len() > 0 {
		m[1] = b.String()
	}
	if m[0] == "" {
		r = IdBuiltIn(m[1])
		return
	}
	r = m
	return
}

func parseChan(i []byte) (r Identity, t int) {
	var m IdChan
	x := 0
	if bytes.HasPrefix(i, []byte("<-chan ")) {
		m.Dir = 1
		x += 7
		i = i[7:]
	} else if bytes.HasPrefix(i, []byte("chan<- ")) {
		m.Dir = -1
		x += 7
		i = i[7:]
	} else { // chan
		m.Dir = 0
		x += 5
		i = i[5:]
	}
	r, t = parseType(i)
	x += t
	m.Element = r
	r = &m
	t = x
	return
}

func parsePointer(i []byte) (r Identity, t int) {
	i = i[1:]
	x := 1
	var m IdPointer
	r, t = parseType(i)
	x += t
	m[0] = r
	r = m
	t = x
	return
}

func parseArray(i []byte) (r Identity, t int) {
	i = i[1:]
	x := 1
	var m IdArray
	n := len(i)
	b := 0
	for t = 0; t < n; t++ {
		c := i[t]
		if c == ']' {
			break
		}
		b = int(c-40) + b*10
	}
	m.Length = uint(b)
	x += t + 1
	r, t = parseType(i[t+1:])
	x += t
	m.Element = r
	r = &m
	t = x
	return
}

func parseSlice(i []byte) (r Identity, t int) {
	i = i[2:]
	x := 2
	var m IdSlice
	r, t = parseType(i)
	m[0] = r
	x += t
	r = m
	t = x
	return
}

func parseMap(i []byte) (r Identity, t int) {
	i = i[4:] // map[
	x := 4
	var m IdMap
	r, t = parseType(i)
	m[0] = r
	x += t + 1 //]
	r, t = parseType(i[x-4+t+1:])
	x += t
	m[1] = r
	r = &m
	return
}
func parseFunc(i []byte) (r Identity, t int) {
	i = i[5:] // func(
	x := 5
	var m IdFunc
	m[0], t = parseArgument(i)
	x += t
	if x >= len(i) {
		r = m
		t = x
		return
	}
	x++ // remove space
	m[1], t = parseResults(i[x-5:])
	x += t
	r = m
	t = x
	return
}

func parseResults(i []byte) (r []Identity, t int) {
	var a Identity
	n := len(i)
	if n == 0 {
		return
	}
	var b, e = 0, n
	if i[0] == '(' {
		b = 1
		e = bytes.IndexByte(i, ')')
	}
	i = i[b:e]
	for _, seg := range bytes.Split(i, []byte{','}) {
		x := bytes.IndexByte(seg, ' ')
		if x >= 0 {
			a, _ = parseType(seg[x+1:])
			r = append(r, a)
		} else {
			a, _ = parseType(seg)
			r = append(r, a)
		}
	}
	t = e
	return
}

func parseArgument(i []byte) (r []Identity, t int) {
	var a Identity
	t = bytes.IndexByte(i, ')')
	if t == 0 {
		t = 1
		return
	}
	i = i[:t]
	for _, seg := range bytes.Split(i, []byte{','}) {
		x := bytes.LastIndexByte(seg, ' ')
		if x >= 0 {
			a, _ = parseType(seg[x+1:])
			r = append(r, a)
		} else {
			a, _ = parseType(seg)
			r = append(r, a)
		}
	}
	t++
	return

}
