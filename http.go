package fn

import (
	"fmt"
	"net/http"
)

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
				_, _ = w.Write([]byte(x.Error()))
			default:
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = fmt.Fprintf(w, `{"code":500,"error":"%s"}`, x)
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
