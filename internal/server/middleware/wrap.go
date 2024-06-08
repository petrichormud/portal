package middleware

import (
	"net/http"
	"slices"

	"github.com/a-h/templ"

	"github.com/petrichormud/portal/internal/component"
	"github.com/petrichormud/portal/internal/server/handler"
	"github.com/petrichormud/portal/internal/server/session"
)

// TODO: Inject the functionality that matches the error from above
func Error(f handler.Func) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch err := f(w, r); err {
		case nil:
		case handler.ErrNotFound:
			w.WriteHeader(http.StatusNotFound)
			templ.Handler(component.ErrNotFound()).ServeHTTP(w, r)
		case handler.ErrInternal:
			w.WriteHeader(http.StatusInternalServerError)
			templ.Handler(component.ErrInternal()).ServeHTTP(w, r)
		case handler.ErrUnauthorized:
			w.WriteHeader(http.StatusUnauthorized)
			templ.Handler(component.ErrInternal()).ServeHTTP(w, r)
		default:
			w.WriteHeader(http.StatusInternalServerError)
			templ.Handler(component.ErrInternal()).ServeHTTP(w, r)
		}
	})
}

func Context(f handler.FuncWithContext) handler.Func {
	return func(w http.ResponseWriter, r *http.Request) error {
		c := handler.Context{
			Writer:  w,
			Request: r,
		}
		return f(c)
	}
}

func IdentifyPlayer(next handler.FuncWithContext) handler.FuncWithContext {
	return func(c handler.Context) error {
		sess, _ := session.Store.Get(c.Request, "session")
		pid, ok := sess.Values["PID"].(int64)
		if ok {
			c.Player.ID = pid
		}
		return next(c)
	}
}

func AuthenticatePlayer(next handler.FuncWithContext) handler.FuncWithContext {
	return func(c handler.Context) error {
		if c.Player.ID == 0 {
			return handler.ErrUnauthorized
		}
		return next(c)
	}
}

type WrapInput struct {
	Handler handler.FuncWithContext
	Inner   []FuncWithContext
	Outer   []Func
}

func Wrap(in WrapInput) http.Handler {
	slices.Reverse(in.Inner)
	slices.Reverse(in.Outer)

	i := in.Handler
	for _, f := range in.Inner {
		i = f(i)
	}

	o := Error(Context(i))
	for _, f := range in.Outer {
		o = f(o)
	}

	return o
}
