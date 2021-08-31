package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-framework/app"
	"os"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show app version",

	Run: func(cmd *cobra.Command, args []string) {
		_, _ = fmt.Fprintf(os.Stdout, "[%s] %s version v1", app.Config.App.Name, app.Config.App.Env)
	},
}

func LoadVersion() {
	Root.AddCommand(versionCmd)
}
