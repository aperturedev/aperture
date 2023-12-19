package cli

import (
	"fmt"
	"github.com/aperturedev/aperture/internal/app"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init app-id",
	Short: "Initialize new aperture app",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println()

		id := args[0]

		fmt.Println(fmt.Sprintf("> Initializing %s app", id))

		a := app.New(id)

		err := a.Save()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, err.Error())
		}

		fmt.Println("> Writing aperture.yaml")
		fmt.Println("> Done!")
		fmt.Println()
		fmt.Println("> Now try running: aperture run")
	},
}
