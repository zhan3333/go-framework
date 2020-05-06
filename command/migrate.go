package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

var action = flag.String("action", "", "操作类型")
var name = flag.String("name", "", "文件名称")
var saveDir = "db/migrations"

func main() {
	flag.Parse()
	switch *action {
	case "create":
		create()
	}
}

func create() {
	var err error
	var out []byte
	fmt.Printf("create migrate: %s", *name)
	fmt.Println()
	cmd := exec.Command("/bin/bash", "-c", fmt.Sprintf("migrate create -ext sql -dir %s %s", saveDir, *name))
	if out, err = cmd.Output(); err != nil {
		fmt.Println(err)
		fmt.Println(string(out))
		os.Exit(1)
	}
	fmt.Println("create success")
}
