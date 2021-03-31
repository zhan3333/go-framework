package main

import (
	"flag"
	"fmt"
	"github.com/zhan3333/go-migrate/v2"
	_ "go-framework/core/boot/console"
	"os"
	"strconv"
)

func init() {
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
		step := 0
		if len(os.Args) >= 4 {
			step, err = strconv.Atoi(os.Args[3])
			if err != nil {
				fmt.Printf("Err: step 参数必须为数字")
			}
		}
		if action == "migrate" {
			if step == 0 {
				step = 999
			}
			if err = migrate.Migrate(step); err != nil {
				fmt.Printf("migrate failed: %+v\n", err)
			}
		}
		if action == "rollback" {
			if step == 0 {
				step = 1
			}
			if err = migrate.Rollback(step); err != nil {
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
