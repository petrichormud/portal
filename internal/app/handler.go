/*
Copyright © 2023 Alec DuBois <alec@petrichormud.com>
*/
package app

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"

	"github.com/petrichormud/portal/internal/component"
)

func Handler(h *http.ServeMux) {
	h.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		c := component.Home()
		handler := templ.Handler(c)
		handler.ServeHTTP(w, r)
	})

	h.HandleFunc("GET /about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "About")
	})
}
