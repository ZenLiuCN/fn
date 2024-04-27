package fn

import (
	"context"
	"net/http"
)

type (
	AuthExtractFn[A any]     func(*http.Request) (A, error)
	AuthValidateFn[A, U any] func(*http.Request, A) (U, error)
)

const (
	AuthorizationContextKey  = "Authorization"
	AuthenticationContextKey = "Authentication"
)

// AuthorizationRead authorization from context
func AuthorizationRead[A any](c context.Context) (a A, ok bool) {
	if c == nil {
		ok = false
		return
	}
	a, ok = c.Value(AuthorizationContextKey).(A)
	return
}

// AuthenticationRead authentication from context
func AuthenticationRead[A any](c context.Context) (a A, ok bool) {
	if c == nil {
		ok = false
		return
	}
	a, ok = c.Value(AuthenticationContextKey).(A)
	return
}

// AuthHandler warp a handler with authority process
func AuthHandler[A, U any](
	extractor AuthExtractFn[A],
	validator AuthValidateFn[A, U],
	handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var v A
		var e error
		v, e = extractor(r)
		if e != nil {
			panic(e)
		}
		var u U
		u, e = validator(r, v)
		if e != nil {
			panic(e)
		}
		ctx := r.Context()
		if ctx == nil {
			r = r.WithContext(context.WithValue(context.WithValue(context.Background(), AuthorizationContextKey, v), AuthenticationContextKey, u))
		} else {
			r = r.WithContext(context.WithValue(context.WithValue(ctx, AuthorizationContextKey, v), AuthenticationContextKey, u))
		}
		handler.ServeHTTP(w, r)
	})
}

// AuthFunc warp a handler function with authority process
func AuthFunc[A, U any](
	extractor AuthExtractFn[A],
	validator AuthValidateFn[A, U],
	next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var v A
		var e error
		v, e = extractor(r)
		if e != nil {
			panic(e)
		}
		var u U
		u, e = validator(r, v)
		if e != nil {
			panic(e)
		}
		ctx := r.Context()
		if ctx == nil {
			r = r.WithContext(context.WithValue(context.WithValue(context.Background(), AuthorizationContextKey, v), AuthenticationContextKey, u))
		} else {
			r = r.WithContext(context.WithValue(context.WithValue(ctx, AuthorizationContextKey, v), AuthenticationContextKey, u))
		}
		next(w, r)
	}
}
