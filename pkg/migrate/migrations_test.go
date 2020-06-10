package migrate_test

import (
	"github.com/stretchr/testify/assert"
	"go-framework/bootstrap"
	db2 "go-framework/pkg/db"
	"go-framework/pkg/migrate"
	"testing"
)

func TestMain(m *testing.M) {
	bootstrap.SetInTest()
	bootstrap.Bootstrap()
	m.Run()
}

func TestGetAllMigrations(t *testing.T) {
	allMigrations := migrate.GetAllMigrations()
	assert.NotNil(t, allMigrations)
}

func TestGetNeedMigrateFiles(t *testing.T) {
	err := db2.Def().Exec("truncate migrations").Error
	assert.Nil(t, err)
	needMigrateFiles := migrate.GetNeedMigrateFiles(migrate.Files)
	assert.NotNil(t, needMigrateFiles)
	assert.Equal(t, len(migrate.Files), len(needMigrateFiles))
	m := migrate.Migration{
		Migration: migrate.Files[0].Key(),
		Batch:     1,
	}
	db2.Def().Save(&m)

	needMigrateFiles = migrate.GetNeedMigrateFiles(migrate.Files)
	assert.Equal(t, 0, len(needMigrateFiles))
}

func TestGetNeedRollbackKeys(t *testing.T) {
	err := db2.Def().Exec("truncate migrations").Error
	assert.Nil(t, err)
	var needRollbackMs []migrate.File
	needRollbackMs = migrate.GetNeedRollbackKeys(1)
	assert.Equal(t, 0, len(needRollbackMs))

	m := migrate.Migration{
		Migration: migrate.Files[0].Key(),
		Batch:     1,
	}
	db2.Def().Create(&m)

	needRollbackMs = migrate.GetNeedRollbackKeys(1)
	assert.Equal(t, 1, len(needRollbackMs))
	assert.Equal(t, migrate.Files[0].Key(), needRollbackMs[0].Key())
}

func TestGetNextBatchNo(t *testing.T) {
	var nextBatch uint
	err := db2.Def().Exec("truncate migrations").Error
	assert.Nil(t, err)
	nextBatch = migrate.GetNextBatchNo()
	assert.Equal(t, uint(1), nextBatch)
	m := migrate.Migration{
		Migration: migrate.Files[0].Key(),
		Batch:     nextBatch,
	}
	db2.Def().Create(&m)

	nextBatch = migrate.GetNextBatchNo()
	assert.Equal(t, uint(2), nextBatch)
}
