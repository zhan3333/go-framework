package main

import (
	"flag"
	"fmt"
	"github.com/zhan3333/go-migrate"
	"go-framework/boot"
	"os"
)

func init() {
	boot.SetInCommand()
	boot.Boot()
	flag.Usage = usage
}

func main() {
	var err error
	flag.Parse()
	var cmd string
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}
	switch cmd {
	case "migrate":
		if len(os.Args) <= 2 {
			fmt.Printf("Err: Migrate 命令需要 action 参数 \n")
			return
		}
		action := os.Args[2]
		if action == "migrate" {
			if err = migrate.Migrate(1); err != nil {
				fmt.Printf("migrate failed: %+v\n", err)
			}
		}
		if action == "rollback" {
			if err = migrate.Rollback(1); err != nil {
				fmt.Printf("rollback failed: %+v\n", err)
			}
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
