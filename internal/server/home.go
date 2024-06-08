package server

import (
	"net/http"

	"github.com/a-h/templ"

	"github.com/petrichormud/portal/internal/component"
	"github.com/petrichormud/portal/internal/server/handler"
	"github.com/petrichormud/portal/internal/server/middleware"
)

var Home http.Handler = middleware.Wrap(middleware.WrapInput{
	Handler: home,
	Inner:   []middleware.FuncWithContext{middleware.IdentifyPlayer},
	Outer:   []middleware.Func{},
})

func home(c handler.Context) error {
	if c.Request.URL.Path == "/" {
		// TODO: Build a handler that renders a template given the Context
		templ.Handler(component.Home()).ServeHTTP(c.Writer, c.Request)
		return nil
	}

	return handler.ErrNotFound
}
