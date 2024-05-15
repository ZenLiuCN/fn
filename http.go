package fn

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-akka/configuration"
	"io"
	"net/http"
	"time"
	"unsafe"
)

var (
	defaultFormatter          = ConvString2BytesRef[HttpError]((*HttpError).Error)
	HttpErrorFormatterDefault = defaultFormatter
)

type HttpErrorFormatter = func(*HttpError) []byte
type HttpError struct {
	Code int
	Err  error
}

func NewHttpError(code int, err error) *HttpError {
	return &HttpError{Code: code, Err: err}
}

func (h *HttpError) Error() string {
	return fmt.Sprintf(`{"code":%d,"error":"%s"}`, h.Code, h.Err)
}

func HttpSafeFunc(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			switch x := recover().(type) {
			case nil:
			case *HttpError:
				w.WriteHeader(x.Code)
				_, _ = w.Write(HttpErrorFormatterDefault(x))
			default:
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = fmt.Fprintf(w, `{"code":500,"error":"%s"}`, x)
			}
		}()
		h.ServeHTTP(w, r)
	}
}
func HttpSafeResolveFunc(resolver func(any) *HttpError, h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			switch x := recover().(type) {
			case nil:
			case *HttpError:
				w.WriteHeader(x.Code)
				_, _ = w.Write(HttpErrorFormatterDefault(x))
			default:
				if resolver != nil {
					v := resolver(x)
					w.WriteHeader(v.Code)
					_, _ = w.Write(HttpErrorFormatterDefault(v))
				} else {
					w.WriteHeader(http.StatusInternalServerError)
					_, _ = fmt.Fprintf(w, `{"code":500,"error":"%s"}`, x)
				}
			}
		}()
		h.ServeHTTP(w, r)
	}
}
func HttpSafeHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			switch x := recover().(type) {
			case nil:
			case *HttpError:
				w.WriteHeader(x.Code)
				_, _ = w.Write([]byte(x.Error()))
			default:
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = fmt.Fprintf(w, `{"code":500,"error":"%s"}`, x)
			}
		}()
		h.ServeHTTP(w, r)
	})
}
func HttpMiddleware(h http.HandlerFunc) func(http.HandlerFunc) http.HandlerFunc {
	return func(fn http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
			fn.ServeHTTP(w, r)
		}
	}
}
func HttpJsonSender[T any]() func(T, http.ResponseWriter) {
	return HttpJsonSend[T]
}
func HttpJsonSend[T any](t T, w http.ResponseWriter) {
	w.Header().Set("Content-Type", JsonMime)
	buf := GetBuffer()
	defer PutBuffer(buf)
	Panic(json.NewEncoder(buf).Encode(t))
	_ = Panic1(w.Write(buf.Bytes()))
}

const (
	JsonMime  = "application/json"
	HoconMime = "application/hocon"
)

// stringData copy from fn/info
func stringData(v string) []byte {
	n := len(v)
	if n == 0 {
		return nil
	}
	type sh struct {
		Data *[0x7fff0000]byte
		Len  int
	}
	p := (*sh)(unsafe.Pointer(&v))
	return p.Data[:n:n]
}
func HttpHoconSend(t *configuration.Config, w http.ResponseWriter) {
	w.Header().Set("Content-Type", HoconMime)
	Panic1(w.Write(stringData(t.String())))
}
func HttpHoconReceive(r *http.Request, checkMime bool) (c *configuration.Config, read bool) {
	if checkMime && r.Header.Get("Content-Type") != HoconMime {
		return
	}
	IgnoreClose(r.Body)
	read = true
	buf := GetBuffer()
	defer PutBuffer(buf)
	Panic1(io.Copy(buf, r.Body))
	c = configuration.ParseString(buf.String())
	return
}

type SSE interface {
	Request() *http.Request
	io.Closer
	GetWriter() EventWriter
	Send(writer EventWriter) bool //false when connection close
}
type EventWriter interface {
	private()
	WithId(id string) EventWriter            // only write once.
	WithEvent(event string) EventWriter      // only write once.
	WithData(data string) EventWriter        // only write once.
	WithRetry(dur time.Duration) EventWriter // only write once.
}
type eventWriter struct {
	id    bool
	event bool
	data  bool
	retry bool
	*bytes.Buffer
}

func (e *eventWriter) private() {

}

var (
	dataBytes  = []byte("data:")
	eventBytes = []byte("event:")
	idBytes    = []byte("id:")
	retryBytes = []byte("retry:")
	fieldEnd   = []byte{'\n'}
)

func (e *eventWriter) WithId(id string) EventWriter {
	if e.id {
		return e
	}
	e.id = true
	_, _ = e.Write(idBytes)
	_, _ = e.WriteString(id)
	_ = e.WriteByte('\n')
	return e
}

func (e *eventWriter) WithEvent(event string) EventWriter {
	if e.event {
		return e
	}
	e.event = true
	_, _ = e.Write(eventBytes)
	_, _ = e.WriteString(event)
	_ = e.WriteByte('\n')
	return e
}

func (e *eventWriter) WithData(data string) EventWriter {
	if e.data {
		return e
	}
	e.data = true
	_, _ = e.Write(dataBytes)
	_, _ = e.WriteString(data)
	_ = e.WriteByte('\n')
	return e
}

func (e *eventWriter) WithRetry(dur time.Duration) EventWriter {
	if e.retry {
		return e
	}
	e.retry = true
	_, _ = e.Write(retryBytes)
	_, _ = fmt.Fprintf(e.Buffer, "%d", dur.Milliseconds())
	_ = e.WriteByte('\n')
	return e
}

type sse struct {
	req *http.Request
	http.ResponseWriter
	http.Flusher
}

func (s *sse) Close() error {
	if s.req.Close {
		return nil
	}
	s.Send(s.GetWriter().WithEvent("close"))
	return nil
}

func (s *sse) GetWriter() EventWriter {
	return &eventWriter{Buffer: GetBuffer()}
}

func (s *sse) Send(writer EventWriter) bool {
	e := writer.(*eventWriter)
	defer PutBuffer(e.Buffer)
	if s.req.Close {
		return false
	}
	_, _ = e.WriteTo(s.ResponseWriter)
	_, _ = s.ResponseWriter.Write(fieldEnd)
	s.Flush()
	return true
}

func (s *sse) Request() *http.Request {
	return s.req
}

// HttpSSEFunc synchronized SSE function
func HttpSSEFunc(fn func(SSE, context.Context)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		sse := &sse{
			r,
			w,
			w.(http.Flusher),
		}
		fn(sse, r.Context())
	}
}
