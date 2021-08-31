package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-framework/cmd/commands"
	"go-framework/core/boot"
	"os"
)

func main() {
	if err := commands.Root.Execute(); err != nil {
		fmt.Printf("err: %+v\n", err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(func() {
		boot.SetInCommand()
		fmt.Println("boot start")
		if err := boot.Boot(); err != nil {
			fmt.Println("boot failed: %w", err)
		}
		fmt.Println("booted")
	})
	commands.LoadMigrate()
	commands.LoadVersion()
}
