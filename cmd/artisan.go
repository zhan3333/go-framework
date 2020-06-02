package main

import (
	"flag"
	"fmt"
	"go-framework/bootstrap"
	"go-framework/cmd/migrate"
	"os"
)

func init() {
	bootstrap.SetInCommand()
	bootstrap.Bootstrap()
	flag.Usage = usage
}

func main() {
	flag.Parse()
	var cmd string
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}
	switch cmd {
	case "migrate":
		if len(os.Args) < 2 {
			fmt.Printf("Err: Migrate 命令需要 action 参数 \n")
			return
		}
		action := os.Args[2]
		if action == "migrate" {
			migrate.Migrate()
		}
		if action == "rollback" {
			migrate.Rollback()
		}
	default:
		flag.Usage()
	}
}

func usage() {
	_, _ = fmt.Fprintf(os.Stdout, `migrater verson: migrater/1.0.0

Commands:
artisan migrate migrate|rollback
`)
	flag.PrintDefaults()
}
