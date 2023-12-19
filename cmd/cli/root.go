package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "aperture",
	Short: `A Fast and Flexible Event Sourcing Projection Runner`,
	Run: func(cmd *cobra.Command, args []string) {
	},
	CompletionOptions: cobra.CompletionOptions{
		HiddenDefaultCmd: true,
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {

		fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}
}
