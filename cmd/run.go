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

		mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
			c := component.Home()
			handler := templ.Handler(c)
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
