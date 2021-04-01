package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-framework/conf"
	"os"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show app version",

	Run: func(cmd *cobra.Command, args []string) {
		_, _ = fmt.Fprintf(os.Stdout, "[%s] %s version v1", conf.Name, conf.Env)
	},
}

func init() {
	Root.AddCommand(versionCmd)
}
