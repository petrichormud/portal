package cmd

import (
	"github.com/spf13/cobra"

	"github.com/petrichormud/portal/internal/app"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the application",
	Long:  `Run the application`,
	Run: func(_ *cobra.Command, _ []string) {
		app.Listen()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
