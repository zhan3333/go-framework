package main

import (
	"flag"
	"fmt"
	"go-framework/bootstrap"
	"go-framework/db/migrations"
	"os"
)

func init() {
	bootstrap.SetInCommand()
	bootstrap.Bootstrap()
	flag.Usage = usage
}

func main() {
	flag.Parse()
	var action string
	if len(os.Args) > 1 {
		action = os.Args[1]
	}
	fmt.Printf("action %s \n", action)
	switch action {
	case "migrate":
		migrate()
	case "rollback":
		rollback()
	default:
		flag.Usage()
	}
}

func usage() {
	_, _ = fmt.Fprintf(os.Stdout, `migrater verson: migrater/1.0.0
Usage: migrater migrate/rollback  

Options:
`)
	flag.PrintDefaults()
}

func migrate() {
	var err error
	mfs := migrations.GetNeedMigrateFiles(migrations.MigrateFiles)
	nextBatch := migrations.GetNextBatchNo()
	fmt.Printf("Migrate file has %d \n", len(mfs))
	if len(mfs) == 0 {
		fmt.Printf("No migrate file need migration \n")
		return
	}
	for _, mf := range mfs {
		fmt.Printf("[Migrating] %s ... \n", mf.Key())
		err = mf.Up()
		if err != nil {
			fmt.Printf("[Migrate failed] %s: %s \n", mf.Key(), err.Error())
			break
		}
		err = migrations.CreateMigrate(mf.Key(), nextBatch)
		if err != nil {
			fmt.Printf("[Migrate failed] %s: %s \n", mf.Key(), err.Error())
			break
		}
		fmt.Printf("[Migrated] %s successed \n", mf.Key())
	}
}

func rollback() {
	var err error
	mfs := migrations.GetNeedRollbackKeys(1)
	fmt.Printf("Rollback file has %d \n", len(mfs))
	if len(mfs) == 0 {
		fmt.Printf("No migrate file need rollback \n")
		return
	}
	for _, mf := range mfs {
		fmt.Printf("[Rollbacking] %s ... \n", mf.Key())
		err = mf.Down()
		if err != nil {
			fmt.Printf("[Rollback failed] %s: %s", mf.Key(), err.Error())
			break
		}
		err = migrations.DeleteMigrate(mf.Key())
		if err != nil {
			fmt.Printf("[Migrate failed] %s: %s", mf.Key(), err.Error())
			break
		}
		fmt.Printf("[Rollbacked] %s successed \n", mf.Key())
	}
}
