package cmd

import (
	"fmt"
	"os"
)

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "lgo",
	Short: "go-framework",
}

var configFile string

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "configs/local.toml", "config file")
}
