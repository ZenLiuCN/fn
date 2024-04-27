package fn

import (
	"net/http"
)

type (
	ErrorResponseFn  func(http.ResponseWriter, *http.Request, any)
	ErrorTranslateFn func(any) (code int, message []byte)
)

func ErrorResponder(fn ErrorTranslateFn) ErrorResponseFn {
	return func(w http.ResponseWriter, r *http.Request, a any) {
		code, buf := fn(a)
		if code > 0 {
			w.WriteHeader(code)
		}
		if len(buf) > 0 {
			_, _ = w.Write(buf)
		}
	}
}
func ErrorResponseHandler(fn ErrorResponseFn, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if re := recover(); re != nil {
				fn(w, r, re)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
func ErrorResponseFunc(fn ErrorResponseFn, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if re := recover(); re != nil {
				fn(w, r, re)
			}
		}()
		next.ServeHTTP(w, r)
	}
}
