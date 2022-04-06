package cmd

import (
	"fmt"
	"os"

	"go-framework/pkg/boot"
)

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "lgo",
	Short: "go-framework",
}

var configFile string
var booted *boot.Booted

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(func() {
		var err error
		if booted, err = boot.Boot(boot.WithConfigFile(configFile)); err != nil {
			booted.Logger.Panicf("boot: %s", err.Error())
		}
	})
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "configs/local.toml", "config file")
}
