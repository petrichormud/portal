package handlers

import (
	"net/http"
	"slices"

	"github.com/a-h/templ"

	"github.com/petrichormud/portal/components"
	"github.com/petrichormud/portal/session"
)

// TODO: Inject the functionality that matches the error from above
func ErrorMiddleware(f Func) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch err := f(w, r); err {
		case nil:
		case ErrNotFound:
			w.WriteHeader(http.StatusNotFound)
			templ.Handler(components.ErrNotFound()).ServeHTTP(w, r)
		case ErrInternal:
			w.WriteHeader(http.StatusInternalServerError)
			templ.Handler(components.ErrInternal()).ServeHTTP(w, r)
		case ErrUnauthorized:
			w.WriteHeader(http.StatusUnauthorized)
			templ.Handler(components.ErrInternal()).ServeHTTP(w, r)
		default:
			w.WriteHeader(http.StatusInternalServerError)
			templ.Handler(components.ErrInternal()).ServeHTTP(w, r)
		}
	})
}

func ContextMiddleware(f FuncWithContext) Func {
	return func(w http.ResponseWriter, r *http.Request) error {
		c := Context{
			Writer:  w,
			Request: r,
		}
		return f(c)
	}
}

func IdentifyPlayerMiddleware(next FuncWithContext) FuncWithContext {
	return func(c Context) error {
		sess, _ := session.Store.Get(c.Request, "session")
		pid, ok := sess.Values["PID"].(int64)
		if ok {
			c.Player.ID = pid
		}
		return next(c)
	}
}

func AuthenticatePlayerMiddleware(next FuncWithContext) FuncWithContext {
	return func(c Context) error {
		if c.Player.ID == 0 {
			return ErrUnauthorized
		}
		return next(c)
	}
}

type WrapInput struct {
	Handler FuncWithContext
	Inner   []MiddlewareFuncWithContext
	Outer   []MiddlewareFunc
}

func Wrap(in WrapInput) http.Handler {
	slices.Reverse(in.Inner)
	slices.Reverse(in.Outer)

	i := in.Handler
	for _, f := range in.Inner {
		i = f(i)
	}

	o := ErrorMiddleware(ContextMiddleware(i))
	for _, f := range in.Outer {
		o = f(o)
	}

	return o
}
