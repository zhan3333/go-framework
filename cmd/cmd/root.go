package cmd

import (
	"github.com/spf13/cobra"
	"go-framework/core/boot"
)

var Root = &cobra.Command{
	Use:   "artisan",
	Short: "命令行",
	Long:  `运行指定的命令行程序`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	boot.SetInCommand()
	boot.Boot()
}
