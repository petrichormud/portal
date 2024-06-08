package handlers

import (
	"net/http"

	"github.com/a-h/templ"

	"github.com/petrichormud/portal/components"
)

var Home http.Handler = Wrap(WrapInput{
	Handler: home,
	Inner:   []MiddlewareFuncWithContext{IdentifyPlayerMiddleware},
	Outer:   []MiddlewareFunc{},
})

func home(c Context) error {
	if c.Request.URL.Path == "/" {
		// TODO: Build a handler that renders a template given the Context
		cmp := components.Home()
		templ.Handler(cmp).ServeHTTP(c.Writer, c.Request)
		return nil
	}

	return ErrNotFound
}
