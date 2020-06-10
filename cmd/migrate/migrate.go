package migrate

import (
	"fmt"
	"go-framework/pkg/migrate"
)

func Migrate() {
	var err error
	mfs := migrate.GetNeedMigrateFiles(migrate.Files)
	nextBatch := migrate.GetNextBatchNo()
	fmt.Printf("Migrate file has %d \n", len(mfs))
	if len(mfs) == 0 {
		fmt.Printf("No Migrate file need migration \n")
		return
	}
	for _, mf := range mfs {
		fmt.Printf("[Migrating] %s ... \n", mf.Key())
		err = mf.Up()
		if err != nil {
			fmt.Printf("[Migrate failed] %s: %s \n", mf.Key(), err.Error())
			break
		}
		err = migrate.CreateMigrate(mf.Key(), nextBatch)
		if err != nil {
			fmt.Printf("[Migrate failed] %s: %s \n", mf.Key(), err.Error())
			break
		}
		fmt.Printf("[Migrated] %s successed \n", mf.Key())
	}
}

func Rollback() {
	var err error
	mfs := migrate.GetNeedRollbackKeys(1)
	fmt.Printf("Rollback file has %d \n", len(mfs))
	if len(mfs) == 0 {
		fmt.Printf("No Migrate file need Rollback \n")
		return
	}
	for _, mf := range mfs {
		fmt.Printf("[Rollbacking] %s ... \n", mf.Key())
		err = mf.Down()
		if err != nil {
			fmt.Printf("[Rollback failed] %s: %s", mf.Key(), err.Error())
			break
		}
		err = migrate.DeleteMigrate(mf.Key())
		if err != nil {
			fmt.Printf("[Migrate failed] %s: %s", mf.Key(), err.Error())
			break
		}
		fmt.Printf("[Rollbacked] %s successed \n", mf.Key())
	}
}
