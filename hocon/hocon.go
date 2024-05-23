package hocon

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ZenLiuCN/fn"
	"github.com/go-akka/configuration"
	ho "github.com/go-akka/configuration/hocon"
	"io"
	"reflect"
	"strings"
	"time"
)

const (
	MimeType = "application/hocon"
)

type (
	Encoder struct {
		io.Writer
	}
	Decoder struct {
		io.Reader
	}
)

func Marshal(obj any) (*configuration.Config, error) {
	return fn.Recover11(configuration.FromObject)(obj)
}
func Unmarshal(c *configuration.Config, p any) (err error) {
	if c == nil || c.IsEmpty() {
		return nil
	}
	var buf *bytes.Buffer
	buf, err = convertJson(c, reflect.TypeOf(p))
	if err != nil {
		return err
	}
	defer func() {
		if buf != nil {
			fn.PutBuffer(buf)
		}
	}()
	//println("json", buf.String())
	err = json.NewDecoder(buf).Decode(p)
	return
}

type writer struct {
	*bytes.Buffer
}

func (w writer) Null() {
	fn.Panic1(w.WriteString("null"))
}

func (w writer) Quote(text string) {
	fn.Panic1(w.Write(fn.Panic1(json.Marshal(text))))
}

func (w writer) Bool(boolean bool) {
	if boolean {
		fn.Panic1(w.WriteString("true"))
	} else {
		fn.Panic1(w.WriteString("false"))
	}
}

func (w writer) Int(v int64) {
	fn.Panic1(fmt.Fprintf(w.Buffer, "%d", v))
}

func (w writer) Float(v float64) {
	fn.Panic1(fmt.Fprintf(w.Buffer, "%f", v))
}

func (w writer) ArrayBegin() {
	fn.Panic(w.WriteByte('['))
}
func (w writer) ArrayEnd() {
	fn.Panic(w.WriteByte(']'))
}
func (w writer) ObjectBegin() {
	fn.Panic(w.WriteByte('{'))
}
func (w writer) ObjectEnd() {
	fn.Panic(w.WriteByte('}'))
}
func (w writer) Comma() {
	fn.Panic(w.WriteByte(','))
}

func (w writer) Colon() {
	fn.Panic(w.WriteByte(':'))
}
func convertJson(c *configuration.Config, t reflect.Type) (buf *bytes.Buffer, err error) {
	defer func() {
		switch r := recover().(type) {
		case nil:
		case error:
			if err != nil {
				err = errors.Join(err, fn.Packer(r, 1))
			} else {
				err = fn.Packer(r, 3)
			}
		default:
			if err != nil {
				err = errors.Join(err, fmt.Errorf("%#+v, %s", r, fn.CallerN(2)))
			} else {
				err = fmt.Errorf("%#+v, %s", r, fn.CallerN(2))
			}
		}
	}()
	if t.Kind() != reflect.Pointer {
		return nil, fmt.Errorf("target not a pointer type")
	}
	t = t.Elem()
	buf = fn.GetBuffer()
	w := writer{buf}
	if c.IsEmpty() {
		w.Null()
		return
	}
	err = writeJson(w, c.Root(), t)
	return
}

var (
	cases = map[string]func(v *ho.HoconValue, w writer){
		"time/Time": func(v *ho.HoconValue, w writer) {
			fn.Panic1(w.Write(fn.Panic1(json.Marshal(fn.Panic1(time.Parse(time.RFC3339, v.GetString()))))))
		},
	}
	jsonMarshaler = reflect.TypeOf((*json.Marshaler)(nil)).Elem()
)

func writeJson(w writer, v *ho.HoconValue, t reflect.Type) (err error) {
	switch t.Kind() {
	case reflect.Pointer:
		if v.IsEmpty() {
			w.Null()
		} else {
			return writeJson(w, v, t.Elem())
		}
	case reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64:
		w.Int(v.GetInt64())
	case
		reflect.Float32,
		reflect.Float64:
		w.Float(v.GetFloat64())
	case reflect.Bool:
		w.Bool(v.GetBoolean())
	case reflect.Array, reflect.Slice:
		if v.IsArray() {
			return writeArray(w, v.GetArray(), t.Elem())
		}
		panic(fmt.Errorf("unassignable type: %s ", t.Kind()))
	case reflect.Map, reflect.Struct:
		if v.IsObject() {
			return writeObject(w, v.GetObject(), t)
		}
		if t.Implements(jsonMarshaler) {
			w.Quote(v.String())
			return
		}
		id := t.PkgPath() + "/" + t.Name()
		if fun, ok := cases[id]; ok {
			fun(v, w)
			return
		}
		panic(fmt.Errorf("unassignable type: %s %s ", id, t.Kind()))
	case reflect.Chan, reflect.Func, reflect.UnsafePointer, reflect.Uintptr, reflect.Interface:
		panic(fmt.Errorf("unassignable type: %s ", t.Kind()))
	default:
		w.Quote(v.GetString())
	}
	return nil
}

func writeArray(w writer, v []*ho.HoconValue, t reflect.Type) (err error) {
	w.ArrayBegin()
	for i, value := range v {
		if i > 0 {
			w.Comma()
		}
		err = writeJson(w, value, t)
		if err != nil {
			return
		}
	}
	w.ArrayEnd()
	return
}
func writeObject(w writer, v *ho.HoconObject, t reflect.Type) (err error) {
	w.ObjectBegin()
	n := 0

	if t.Kind() == reflect.Map {
		for key, value := range v.Items() {
			if n > 0 {
				w.Comma()
			}
			w.Quote(key)
			w.Colon()
			err = writeJson(w, value, t.Elem())
			if err != nil {
				return
			}
			n++
		}
	} else if t.Kind() == reflect.Struct {
		info := map[string]reflect.Type{}
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			t := f.Tag.Get("json")
			if x := strings.IndexByte(t, ','); x >= 0 {
				t = t[:x]
			}
			if t != "" {
				info[t] = f.Type
			} else {
				info[f.Name] = f.Type
			}
		}
		for key, value := range v.Items() {
			if n > 0 {
				w.Comma()
			}
			w.Quote(key)
			w.Colon()
			err = writeJson(w, value, info[key])
			if err != nil {
				return
			}
			n++
		}
	} else {
		panic(fmt.Errorf("not map or struct of %v", t))
	}
	w.ObjectEnd()
	return
}
func can[T any](v func() T) (ok bool) {
	defer func() {
		switch recover().(type) {
		case nil:
		default:
			ok = false
		}
	}()
	v()
	return true
}
