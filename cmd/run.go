package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/spf13/cobra"

	"github.com/petrichormud/portal/internal/component"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the application",
	Long:  `Run the application`,
	Run: func(_ *cobra.Command, _ []string) {
		mux := http.NewServeMux()

		static := http.FileServer(http.Dir("./static"))
		mux.Handle("GET /static/", http.StripPrefix("/static/", static))

		mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/" {
				handler := templ.Handler(component.Home())
				handler.ServeHTTP(w, r)
				return
			}

			handler := templ.Handler(component.NotFound())
			w.WriteHeader(http.StatusNotFound)
			handler.ServeHTTP(w, r)
		})

		mux.HandleFunc("GET /about", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "About")
		})

		s := &http.Server{
			Addr:    ":8008",
			Handler: mux,
		}

		log.Fatal(s.ListenAndServe())
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
