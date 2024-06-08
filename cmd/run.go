package cmd

import (
	"log"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/petrichormud/portal/internal/server"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the application",
	Long:  `Run the application`,
	Run: func(_ *cobra.Command, _ []string) {
		mux := http.NewServeMux()

		static := http.FileServer(http.Dir("./static"))
		mux.Handle("GET /static/", http.StripPrefix("/static/", static))

		mux.Handle("GET /", server.Home)

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
