package main

import (
	"fmt"
	"go-framework/cmd/cmd"
	"os"
)

func main() {
	if err := cmd.Root.Execute(); err != nil {
		fmt.Printf("err: %+v\n", err)
		os.Exit(1)
	}
}
