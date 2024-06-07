/*
Copyright © 2023 Alec DuBois <alec@petrichormud.com>
*/
package app

import (
	"log"
	"net/http"
)

func Listen() {
	mux := http.NewServeMux()

	Handler(mux)

	s := &http.Server{
		Addr:    ":8008",
		Handler: mux,
	}

	log.Fatal(s.ListenAndServe())
}
